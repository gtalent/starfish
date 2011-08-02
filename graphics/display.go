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

import (
	"time"
	"sdl"
	"sdl/ttf"
)

//Holds a Drawer and its Canvas.
type canvasHolder struct {
	canvas Canvas
	drawer func(*Canvas)
}

type Display struct {
	surface *sdl.Surface
	title   string
	panes   []*canvasHolder
	dead    chan interface{}
	running bool
}

func NewDisplay() *Display {
	s := new(Display)
	s.panes = make([]*canvasHolder, 0)
	s.dead = make(chan interface{})
	sdl.WM_SetCaption(s.title, "")
	return s
}

func (me *Display) AddDrawer(drawer func(*Canvas)) {
	ch := new(canvasHolder)
	ch.drawer = drawer
	ch.canvas = newCanvas(me.surface)
	me.panes = append(me.panes, ch)
}

func (me *Display) RemoveDrawer(drawer func(*Canvas)) {
	for n, a := range me.panes {
		if a.drawer == drawer {
			end := len(me.panes) - 1
			for i := n; i < end; i++ {
				me.panes[i] = me.panes[i+1]
			}
			me.panes = me.panes[0 : len(me.panes)-1]
			break
		}
	}
}

func (me *Display) run() {
	for me.running {
		for _, a := range me.panes {
			a.canvas.pane = me.surface
			a.canvas.load()
			a.drawer(&a.canvas)
		}
		me.surface.Flip()
		time.Sleep(16000000)
	}
	me.dead <- nil
}

//Opens a window.
func (me *Display) Open(width, height int) {
	if me.surface == nil {
		sdl.Init(sdl.INIT_VIDEO)
		ttf.Init()
		me.surface = sdl.SetVideoMode(width, height, 32, sdl.RESIZABLE|sdl.DOUBLEBUF)
		me.running = true
		go me.run()
	}
}

//Closes the window.
func (me *Display) Close() {
	if me.surface != nil {
		me.running = false
		<-me.dead
		me.surface.Free()
		me.surface = nil
		sdl.Quit()
	}
}
