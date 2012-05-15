/*
   Copyright 2011-2012 gtalent2@gmail.com

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

import (
	"sync"
	"time"
)

type flynode struct {
	sync.Mutex
	key     key
	val     interface{}
	clients int
}

type key interface {
	String() string
}

type stringKey string

func (me stringKey) String() string {
	return string(me)
}

type flyweight struct {
	lock     sync.Mutex
	items    map[string]*flynode
	loader   func(*flyweight, key) interface{}
	unloader func(*flyweight, key, interface{})
}

func newFlyweight(loader func(*flyweight, key) interface{}, unloader func(*flyweight, key, interface{})) *flyweight {
	r := new(flyweight)
	r.loader = loader
	r.unloader = unloader
	r.items = make(map[string]*flynode)
	return r
}

func (me *flyweight) checkout(key key) interface{} {
	for {
		node, ok := me.items[key.String()]
		if !ok {
			node = new(flynode)
			node.key = key
			node.val = me.loader(me, key)
			node.clients = 1

			me.lock.Lock()
			if n, ok := me.items[key.String()]; ok {
				me.lock.Unlock()
				n.Lock()
				if n.clients == 0 {
					n.Unlock()
					time.Sleep(0)
					continue
				}
				n.clients++
				n.Unlock()
				return n.val
			}
			me.items[key.String()] = node
			me.lock.Unlock()
			return node.val
		} else {
			node.Lock()
			if node.clients == 0 {
				node.Unlock()
				time.Sleep(0)
				continue
			}
			node.clients++
			node.Unlock()
			return node.val
		}
	}
	return nil
}

func (me *flyweight) checkin(key key) {
	n := me.items[key.String()]
	if n == nil {
		me.lock.Unlock()
		return
	}
	n.Lock()
	n.clients--
	if n.clients == 0 {
		delete(me.items, key.String())
		me.unloader(me, key, n.val)
	}
	n.Unlock()
}
