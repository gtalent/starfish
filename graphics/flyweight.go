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

import "sync"

type flynode struct {
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
	lock     *sync.Mutex
	items    map[string]*flynode
	loader   func(key) interface{}
	unloader func(key, interface{})
}

func newFlyweight(loader func(key) interface{}, unloader func(key, interface{})) *flyweight {
	r := new(flyweight)
	r.lock = new(sync.Mutex)
	r.loader = loader
	r.unloader = unloader
	r.items = make(map[string]*flynode)
	return r
}

func (me *flyweight) checkout(key key) interface{} {
	me.lock.Lock()
	node, ok := me.items[key.String()]
	if !ok {
		node = new(flynode)
		node.key = key
		node.val = me.loader(key)
		node.clients = 0
		me.items[key.String()] = node
	}
	node.clients++
	me.lock.Unlock()
	return node.val
}

func (me *flyweight) checkin(key key) {
	me.lock.Lock()
	n := me.items[key.String()]
	if n == nil {
		me.lock.Unlock()
		return
	}
	n.clients--
	if n.clients == 0 {
		delete(me.items, key.String())
		me.unloader(key, n.val)
	}
	me.lock.Unlock()
}
