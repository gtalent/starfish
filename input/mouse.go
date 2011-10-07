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
	"wombat/core/util"
)

const (
	clickTimeout int64 = 200000000
)

//clicker
type clicker struct {
	mice             [256]mouseButton
	input            chan sdl.Event
	clickListeners   []func(byte, util.Point)
	pressListeners   []func(byte, util.Point)
	releaseListeners []func(byte, util.Point)
	addUpChan        chan func(byte, util.Point)
	addDownChan      chan func(byte, util.Point)
	addClickChan     chan func(byte, util.Point)
	removeUpChan     chan func(byte, util.Point)
	removeDownChan   chan func(byte, util.Point)
	removeClickChan  chan func(byte, util.Point)
}

//Makes and runs a clicker
func newClicker() clicker {
	var c clicker
	c.clickListeners = make([]func(byte, util.Point), 0)
	c.pressListeners = make([]func(byte, util.Point), 0)
	c.releaseListeners = make([]func(byte, util.Point), 0)
	c.input = make(chan sdl.Event)
	c.addClickChan = make(chan func(byte, util.Point))
	c.addUpChan = make(chan func(byte, util.Point))
	c.addDownChan = make(chan func(byte, util.Point))
	c.removeClickChan = make(chan func(byte, util.Point))
	c.removeUpChan = make(chan func(byte, util.Point))
	c.removeDownChan = make(chan func(byte, util.Point))
	go c.run()
	return c
}

func (me *clicker) addClick(f func(byte, util.Point)) {
	me.clickListeners = append(me.clickListeners, f)
}

func (me *clicker) addDown(f func(byte, util.Point)) {
	me.pressListeners = append(me.pressListeners, f)
}

func (me *clicker) addUp(f func(byte, util.Point)) {
	me.releaseListeners = append(me.releaseListeners, f)
}

func (me *clicker) removeClick(f func(byte, util.Point)) {
	l := me.clickListeners
	for i, _ := range l {
		if l[i] == f {
			l[i] = l[len(l)-1]
			l = l[0 : len(l)-1]
			break
		}
	}
}

func (me *clicker) removeDown(f func(byte, util.Point)) {
	l := me.pressListeners
	for i, _ := range l {
		if l[i] == f {
			l[i] = l[len(l)-1]
			l = l[0 : len(l)-1]
			break
		}
	}
}

func (me *clicker) removeUp(f func(byte, util.Point)) {
	l := me.releaseListeners
	for i, _ := range l {
		if l[i] == f {
			l[i] = l[len(l)-1]
			l = l[0 : len(l)-1]
			break
		}
	}
}

func (me *clicker) run() {
	f := func(i byte, p util.Point) {
		time.Sleep(clickTimeout)
		if me.mice[i].lastRelease > me.mice[i].lastPress {
			//call release listeners
			for _, a := range me.clickListeners {
				go a(i, p)
			}
		} else {
			//call press listeners
			for _, a := range me.pressListeners {
				go a(i, p)
			}
		}
	}
	for {
		select {
		case et := <-me.input:
			switch et.(*sdl.MouseButtonEvent).Type {
			case sdl.MOUSEBUTTONUP: //release
				i := et.(*sdl.MouseButtonEvent).Button
				if clickTimeout < time.Nanoseconds()-me.mice[i].lastPress {
					var p util.Point
					p.X = int(et.(*sdl.MouseButtonEvent).X)
					p.Y = int(et.(*sdl.MouseButtonEvent).Y)
					for _, a := range me.releaseListeners {
						go a(i, p)
					}
				}
				me.mice[i].lastRelease = time.Nanoseconds()
			case sdl.MOUSEBUTTONDOWN: //press
				i := et.(*sdl.MouseButtonEvent).Button
				me.mice[i].lastPress = time.Nanoseconds()
				var p util.Point
				p.X = int(et.(*sdl.MouseButtonEvent).X)
				p.Y = int(et.(*sdl.MouseButtonEvent).Y)
				go f(i, p)
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
