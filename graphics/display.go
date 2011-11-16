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

type Drawer interface {
	Draw(*Canvas)
}

type drawFunc func(*Canvas)

func (me drawFunc) Draw(c *Canvas) {
	me(c)
}

//Holds a Drawer and its Canvas.
type canvasHolder struct {
	canvas Canvas
	drawer Drawer
}

var screen *sdl.Surface
var displayTitle string
var drawers []*canvasHolder
var displayDead chan interface{}
var running bool

func NewDisplay() {
	drawers = make([]*canvasHolder, 0)
	displayDead = make(chan interface{})
}

//Sets the title of the window.
func SetDisplayTitle(title string) {
	displayTitle = title
	if screen != nil {
		sdl.WM_SetCaption(displayTitle, "")
	}
}

//Returns the title of this window.
func GetDisplayTitle() string {
	return displayTitle
}

func AddDrawer(drawer Drawer) {
	ch := new(canvasHolder)
	ch.drawer = drawer
	ch.canvas = newCanvas(screen)
	drawers = append(drawers, ch)
}

func RemoveDrawer(drawer Drawer) {
	for n, a := range drawers {
		if a.drawer == drawer {
			end := len(drawers) - 1
			for i := n; i < end; i++ {
				drawers[i] = drawers[i+1]
			}
			drawers = drawers[0 : len(drawers)-1]
			break
		}
	}
}

func AddDrawFunc(drawer func(*Canvas)) {
	ch := new(canvasHolder)
	ch.drawer = drawFunc(drawer)
	ch.canvas = newCanvas(screen)
	drawers = append(drawers, ch)
}

func RemoveDrawFunc(drawer func(*Canvas)) {
	for n, a := range drawers {
		if a.drawer == drawFunc(drawer) {
			end := len(drawers) - 1
			for i := n; i < end; i++ {
				drawers[i] = drawers[i+1]
			}
			drawers = drawers[0 : len(drawers)-1]
			break
		}
	}
}

func run() {
	for running {
		for _, a := range drawers {
			a.canvas.pane = screen
			a.canvas.load()
			a.drawer.Draw(&a.canvas)
		}
		screen.Flip()
		time.Sleep(16000000)
	}
	displayDead <- nil
}

//Opens a window.
func OpenDisplay(width, height int) {
	if screen == nil {
		sdl.Init(sdl.INIT_VIDEO)
		ttf.Init()
		screen = sdl.SetVideoMode(width, height, 32, sdl.RESIZABLE|sdl.DOUBLEBUF)
		running = true
		sdl.WM_SetCaption(displayTitle, "")
		go run()
	}
}

//Closes the window.
func CloseDisplay() {
	if screen != nil {
		running = false
		<-displayDead
		screen.Free()
		screen = nil
		sdl.Quit()
	}
}
