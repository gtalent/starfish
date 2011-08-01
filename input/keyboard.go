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

//clicker
type keyboard struct {
	keys             [256]key
	input            chan sdl.Event
	typeListeners   []func(byte)
	pressListeners   []func(byte)
	releaseListeners []func(byte)
	addUpChan        chan func(byte)
	addDownChan      chan func(byte)
	addTypeChan     chan func(byte)
	removeUpChan     chan func(byte)
	removeDownChan   chan func(byte)
	removeTypeChan  chan func(byte)
}

//Makes and runs a keyboard
func newKeyboard() keyboard {
	var c keyboard
	c.typeListeners = make([]func(byte), 0)
	c.pressListeners = make([]func(byte), 0)
	c.releaseListeners = make([]func(byte), 0)
	c.input = make(chan sdl.Event)
	c.addTypeChan = make(chan func(byte))
	c.addUpChan = make(chan func(byte))
	c.addDownChan = make(chan func(byte))
	c.removeTypeChan = make(chan func(byte))
	c.removeUpChan = make(chan func(byte))
	c.removeDownChan = make(chan func(byte))
	go c.run()
	return c
}

func (me *keyboard) addType(f func(byte)) {
	me.typeListeners = append(me.typeListeners, f)
}

func (me *keyboard) addDown(f func(byte)) {
	me.pressListeners = append(me.pressListeners, f)
}

func (me *keyboard) addUp(f func(byte)) {
	me.releaseListeners = append(me.releaseListeners, f)
}

func (me *keyboard) removeType(f func(byte)) {
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
}

func (me *keyboard) removeDown(f func(byte)) {
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
}

func (me *keyboard) removeUp(f func(byte)) {
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
}

func (me *keyboard) run() {
	f := func(i byte) {
		time.Sleep(clickTimeout)
		if me.keys[i].lastRelease > me.keys[i].lastPress {
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
		case et := <-me.input:
			switch et.(*sdl.MouseButtonEvent).Type {
			case sdl.MOUSEBUTTONUP: //release
				i := et.(*sdl.MouseButtonEvent).Button
				if clickTimeout < time.Nanoseconds()-me.keys[i].lastPress {
					for _, a := range me.releaseListeners {
						go a(i)
					}
				}
				me.keys[i].lastRelease = time.Nanoseconds()
			case sdl.MOUSEBUTTONDOWN: //press
				i := et.(*sdl.MouseButtonEvent).Button
				me.keys[i].lastPress = time.Nanoseconds()
				go f(i)
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
	lastType   int64 // the time of the last mouse click 
}
