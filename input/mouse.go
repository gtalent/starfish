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

//Click
type Click struct {
	//The time of the click.
	Time int64
	//Indicates whether or not this is a double click.
	Double bool
}

//clickMgr
type clicker struct {
	click          [255]Click
	input          chan byte
	clickListeners []func(Click)
}

func (me *clicker) run() {
	for {
		i := <-me.input
		t := time.Nanoseconds()
		if t-me.click[i].Time > clickTimeout { //then this is a new click
			me.click[i].Time = t
			me.click[i].Double = false
			go func() {
				time.Sleep(clickTimeout)
				for _, a := range me.clickListeners {
					go a(me.click[i])
				}
			}()
		} else {
			me.click[i].Double = true
		}
	}
}

//Makes and runs a clicker
func newClicker() clicker {
	var c clicker
	c.input = make(chan byte)
	c.clickListeners = make([]func(Click), 0)
	go c.run()
	return c
}

//mouseButton
type mouseButton struct {
	lastPress   int64 // the time of the last mousebutton press
	lastRelease int64 // the time of the last mousebutton release
	lastClick   int64 // the time of the last mouse click 
}

//mouseManager
type mouseManager struct {
	mice                     [256]mouseButton
	mouseButtonDownListeners []func(int)
	mouseButtonUpListeners   []func(int)
	input                    chan *sdl.MouseButtonEvent
	addMouseUpChan chan func(int)
	addMouseDownChan chan func(int)
	removeMouseUpChan chan func(int)
	removeMouseDownChan chan func(int)
	clicker                  clicker
}

//Creates a new mouseManager and runs it.
func newMouseManager() mouseManager {
	var m mouseManager
	m.mouseButtonDownListeners = make([]func(int), 0)
	m.mouseButtonUpListeners = make([]func(int), 0)
	m.input = make(chan *sdl.MouseButtonEvent)
	m.clicker = newClicker()
	go m.run()
	return m
}

func (me *mouseManager) addMouseDown(f func(int)) {
	me.mouseButtonDownListeners = append(me.mouseButtonDownListeners, f)
}

func (me *mouseManager) addMouseUp(f func(int)) {
	me.mouseButtonUpListeners = append(me.mouseButtonUpListeners, f)
}

func (me *mouseManager) removeMouseDown(f func(int)) {
	l := me.mouseButtonDownListeners
	var i int
	for i, _ = range l {
		if l[i] == f {
			break
		}
	}

	for i := i; i + 1 < len(l); i++ {
		l[i] = l[i+1]
	}
}

func (me *mouseManager) removeMouseUp(f func(int)) {
	l := me.mouseButtonUpListeners
	var i int
	for i, _ = range l {
		if l[i] == f {
			break
		}
	}

	for i := i; i + 1 < len(l); i++ {
		l[i] = l[i+1]
	}
}

func (me *mouseManager) run() {
	for {
		et := <-me.input
		switch et.Type {
		case sdl.MOUSEBUTTONUP:
			t := time.Nanoseconds()
			if t-me.mice[et.Button].lastRelease < clickTimeout {
				//if click
				me.clicker.input <- et.Button
			} else {
				//if non-click release
				for _, a := range me.mouseButtonUpListeners {
					go a(int(et.Button))
				}
			}
			me.mice[et.Button].lastRelease = time.Nanoseconds()
		case sdl.MOUSEBUTTONDOWN:
			for _, a := range me.mouseButtonDownListeners {
				go a(int(et.Button))
			}
			me.mice[et.Button].lastPress = time.Nanoseconds()
		}
	}

	for {
		select {
		case a := <-me.addMouseUpChan:
			me.addMouseUp(a)
		case b := <-me.addMouseDownChan:
			me.addMouseDown(b)
		case c := <-me.removeMouseUpChan:
			me.removeMouseUp(c)
		case d := <-me.removeMouseDownChan:
			me.removeMouseDown(d)
		default:
		}
	}
}
