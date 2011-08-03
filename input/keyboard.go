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
package input

import (
	"time"
	"sdl"
)

const (
	typeTimeout int64 = 150000000
)

type KeyEvent struct {
	Key      int
	Shift    bool
	Ctrl     bool
	Meta     bool
	Alt      bool
	CapsLock bool
}

func (me *KeyEvent) setMods(mod uint32) {
	me.Shift = 0 != mod|sdl.KMOD_LSHIFT|sdl.KMOD_RSHIFT
	me.Ctrl = 0 != mod|sdl.KMOD_LCTRL|sdl.KMOD_RCTRL
	me.Meta = 0 != mod|sdl.KMOD_LMETA|sdl.KMOD_RMETA
	me.Alt = 0 != mod|sdl.KMOD_LALT|sdl.KMOD_RALT
	me.CapsLock = 0 != mod|sdl.KMOD_CAPS|sdl.KMOD_CAPS
}

//clicker
type keyboard struct {
	keys             [1000]key
	input            chan sdl.Event
	typeListeners    []func(KeyEvent)
	pressListeners   []func(KeyEvent)
	releaseListeners []func(KeyEvent)
	addUpChan        chan func(KeyEvent)
	addDownChan      chan func(KeyEvent)
	addTypeChan      chan func(KeyEvent)
	removeUpChan     chan func(KeyEvent)
	removeDownChan   chan func(KeyEvent)
	removeTypeChan   chan func(KeyEvent)
}

//Makes and runs a keyboard
func newKeyboard() keyboard {
	var c keyboard
	c.typeListeners = make([]func(KeyEvent), 0)
	c.pressListeners = make([]func(KeyEvent), 0)
	c.releaseListeners = make([]func(KeyEvent), 0)
	c.input = make(chan sdl.Event)
	c.addTypeChan = make(chan func(KeyEvent))
	c.addUpChan = make(chan func(KeyEvent))
	c.addDownChan = make(chan func(KeyEvent))
	c.removeTypeChan = make(chan func(KeyEvent))
	c.removeUpChan = make(chan func(KeyEvent))
	c.removeDownChan = make(chan func(KeyEvent))
	go c.run()
	return c
}

func (me *keyboard) addType(f func(KeyEvent)) {
	me.typeListeners = append(me.typeListeners, f)
}

func (me *keyboard) addDown(f func(KeyEvent)) {
	me.pressListeners = append(me.pressListeners, f)
}

func (me *keyboard) addUp(f func(KeyEvent)) {
	me.releaseListeners = append(me.releaseListeners, f)
}

func (me *keyboard) removeType(f func(KeyEvent)) {
	l := me.typeListeners
	var i int
	for i, _ = range l {
		if l[i] == f {
			break
		}
	}

	for i := i; i+1 < len(l); i++ {
		l[i] = l[i+1]
	}
	l = l[0:len(l)-1]
}

func (me *keyboard) removeDown(f func(KeyEvent)) {
	l := me.pressListeners
	var i int
	for i, _ = range l {
		if l[i] == f {
			break
		}
	}

	for i := i; i+1 < len(l); i++ {
		l[i] = l[i+1]
	}
	l = l[0:len(l)-1]
}

func (me *keyboard) removeUp(f func(KeyEvent)) {
	l := me.releaseListeners
	var i int
	for i, _ = range l {
		if l[i] == f {
			break
		}
	}

	for i := i; i+1 < len(l); i++ {
		l[i] = l[i+1]
	}
	l = l[0:len(l)-1]
}

func (me *keyboard) run() {
	f := func(i KeyEvent) {
		time.Sleep(clickTimeout)
		if me.keys[i.Key].lastRelease > me.keys[i.Key].lastPress {
			//call type listeners
			for _, a := range me.typeListeners {
				go a(i)
			}
		} else {
			//call press listeners
			for _, a := range me.pressListeners {
				go a(i)
			}
		}
	}
	for {
		select {
		case in := <-me.input:
			et := in.(*sdl.KeyboardEvent)
			i := et.Keysym.Sym
			var event KeyEvent
			event.Key = int(i)
			event.setMods(et.Keysym.Mod)
			switch et.Type {
			case sdl.KEYUP: //release
				if clickTimeout < time.Nanoseconds()-me.keys[i].lastPress {
					for _, a := range me.releaseListeners {
						go a(event)
					}
				}
				me.keys[i].lastRelease = time.Nanoseconds()
			case sdl.KEYDOWN: //press
				me.keys[i].lastPress = time.Nanoseconds()
				go f(event)
			}
		//manage listeners
		case a := <-me.addUpChan:
			me.addUp(a)
		case b := <-me.addDownChan:
			me.addDown(b)
		case c := <-me.removeUpChan:
			me.removeUp(c)
		case d := <-me.removeDownChan:
			me.removeDown(d)
		case e := <-me.addTypeChan:
			me.addType(e)
		case f := <-me.removeTypeChan:
			me.removeType(f)
		}
	}
}


//mouseButton
type key struct {
	lastPress   int64 // the time of the last mousebutton press
	lastRelease int64 // the time of the last mousebutton release
	lastType    int64 // the time of the last mouse click 
}
