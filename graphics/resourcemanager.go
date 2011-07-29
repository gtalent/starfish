/*
   Copyright 2011 gtalent2@gmail.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package graphics

type resourceNode struct {
	uses int
	rsrc interface{}
}

type resourceCatalog struct {
	images   map[string]*resourceNode
	checkout chan interface{}
	checkin  chan interface{}
	load     func(string) (interface{}, bool)
}

func newResourceCatalog(load func(string) (interface{}, bool)) (r resourceCatalog) {
	r.images = make(map[string]*resourceNode)
	r.checkout = make(chan interface{})
	r.checkin = make(chan interface{})
	r.load = load
	go r.run()
	return r
}

func (me *resourceCatalog) run() {
	for {
		select {
		case input := <-me.checkout:
			key := input.(string)
			i, ok := me.images[key]
			if ok {
				i.uses++
				me.checkout <- i.rsrc
			} else {
				tmp, ok := me.load(key)
				if ok {
					i = new(resourceNode)
					i.rsrc = tmp
					i.uses++
					me.images[key] = i
					me.checkout <- i.rsrc
				} else {
					me.checkout <- nil
				}
			}
		case input := <-me.checkin:
			key := input.(string)
			i, ok := me.images[key]
			if ok {
				i.uses--
				if i.uses == 0 {
					me.images[key] = nil, false
				}
				me.checkin <- true
			} else {
				me.checkin <- false
			}
		}
	}
}
