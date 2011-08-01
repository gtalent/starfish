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
type click struct {
	//The time of the click.
	time int64
}

//clickMgr
type clicker struct {
	click            [256]click
	mice             [256]mouseButton
	release          chan byte
	press            chan byte
	clickListeners   []func(byte)
	pressListeners   []func(byte)
	releaseListeners []func(byte)
}

func (me *clicker) addDown(f func(byte)) {
	me.pressListeners = append(me.pressListeners, f)
}

func (me *clicker) addUp(f func(byte)) {
	me.releaseListeners = append(me.releaseListeners, f)
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
		case i := <-me.release:
			me.mice[i].lastRelease = time.Nanoseconds()
			if me.mice[i].lastRelease < me.mice[i].lastPress {
				for _, a := range me.releaseListeners {
					go a(i)
				}
			}
			me.mice[i].lastRelease = time.Nanoseconds()
		case p := <-me.press:
			me.mice[p].lastPress = time.Nanoseconds()
			go f(p)
		}
	}
}

//Makes and runs a clicker
func newClicker() clicker {
	var c clicker
	c.release = make(chan byte)
	c.clickListeners = make([]func(byte), 0)
	c.pressListeners = make([]func(byte), 0)
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
	mice                [256]mouseButton
	input               chan *sdl.MouseButtonEvent
	addUpChan      chan func(byte)
	addDownChan    chan func(byte)
	removeUpChan   chan func(byte)
	removeDownChan chan func(byte)
	clicker             clicker
}

//Creates a new mouseManager and runs it.
func newMouseManager() mouseManager {
	var m mouseManager
	m.input = make(chan *sdl.MouseButtonEvent)
	m.addUpChan = make(chan func(byte))
	m.addDownChan = make(chan func(byte))
	m.removeUpChan = make(chan func(byte))
	m.removeDownChan = make(chan func(byte))
	m.clicker = newClicker()
	go m.run()
	return m
}

func (me *mouseManager) run() {
	for {
		et := <-me.input
		switch et.Type {
		case sdl.MOUSEBUTTONUP:
			me.clicker.release <- et.Button
		case sdl.MOUSEBUTTONDOWN:
			me.clicker.press <- et.Button
		}
	}

	for loop := true; loop; {
		select {
		case a := <-me.addUpChan:
			me.clicker.addUp(a)
		case b := <-me.addDownChan:
			me.clicker.addDown(b)
		case c := <-me.removeUpChan:
			me.clicker.removeUp(c)
		case d := <-me.removeDownChan:
			me.clicker.removeDown(d)
		default:
			loop = false
		}
	}
}
