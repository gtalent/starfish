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

//Passed into a resourceCatalog to load the resource.
type loader interface {
	load(key string) interface{}
}

type resourceNode struct {
	uses int
	rsrc interface{}
}

type resourceCatalog struct {
	images   map[string]*resourceNode
	checkout, checkin chan interface{}
	loader loader
}

func newResourceCatalog(loader loader) (r resourceCatalog) {
	r.images = make(map[string]*resourceNode)
	r.checkout = make(chan interface{})
	r.checkin = make(chan interface{})
	r.loader = loader
	r.run()
	return r
}

func (me *resourceCatalog) run() {
	for {
		select {
		case input := <-me.checkout:
			path := input.(string)
			i, ok := me.images[path]
			if ok {
				i.uses++
				me.checkout <- i.rsrc
			} else {
				tmp := me.loader.load(path)
				if tmp != nil {
					i = new(resourceNode)
					i.rsrc = tmp
					i.uses++
					me.images[path] = i
					me.checkout <- i.rsrc
				}
			}
		case input := <-me.checkin:
			path := input.(string)
			i, ok := me.images[path]
			if ok {
				i.uses--
				me.checkin <- true
			} else {
				me.checkin <- false
			}
			me.checkout <- nil
		}
	}
}


