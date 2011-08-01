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
	clickTimeout int64 = 1000000000
)

//clicker
type clicker struct {
	mice             [256]mouseButton
	input            chan sdl.Event
	clickListeners   []func(byte)
	pressListeners   []func(byte)
	releaseListeners []func(byte)
	addUpChan        chan func(byte)
	addDownChan      chan func(byte)
	addClickChan      chan func(byte)
	removeUpChan     chan func(byte)
	removeDownChan   chan func(byte)
	removeClickChan   chan func(byte)
}

//Makes and runs a clicker
func newClicker() clicker {
	var c clicker
	c.clickListeners = make([]func(byte), 0)
	c.pressListeners = make([]func(byte), 0)
	c.input = make(chan sdl.Event)
	c.addClickChan = make(chan func(byte))
	c.addUpChan = make(chan func(byte))
	c.addDownChan = make(chan func(byte))
	c.removeClickChan = make(chan func(byte))
	c.removeUpChan = make(chan func(byte))
	c.removeDownChan = make(chan func(byte))
	go c.run()
	return c
}

func (me *clicker) addClick(f func(byte)) {
	me.clickListeners = append(me.clickListeners, f)
}

func (me *clicker) addDown(f func(byte)) {
	me.pressListeners = append(me.pressListeners, f)
}

func (me *clicker) addUp(f func(byte)) {
	me.releaseListeners = append(me.releaseListeners, f)
}

func (me *clicker) removeClick(f func(byte)) {
	l := me.clickListeners
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

func (me *clicker) removeDown(f func(byte)) {
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

func (me *clicker) removeUp(f func(byte)) {
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

func (me *clicker) run() {
	f := func(i byte) {
		time.Sleep(clickTimeout)
		if me.mice[i].lastRelease > me.mice[i].lastPress {
			//call release listeners
			for _, a := range me.clickListeners {
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
				me.mice[i].lastRelease = time.Nanoseconds()
				if me.mice[i].lastRelease < me.mice[i].lastPress {
					for _, a := range me.releaseListeners {
						go a(i)
					}
				}
				me.mice[i].lastRelease = time.Nanoseconds()
			case sdl.MOUSEBUTTONDOWN: //press
				i := et.(*sdl.MouseButtonEvent).Button
				me.mice[i].lastPress = time.Nanoseconds()
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
		case e := <-me.addClickChan:
			me.addClick(e)
		case f := <-me.removeClickChan:
			me.removeClick(f)
		}
	}
}


//mouseButton
type mouseButton struct {
	lastPress   int64 // the time of the last mousebutton press
	lastRelease int64 // the time of the last mousebutton release
	lastClick   int64 // the time of the last mouse click 
}
