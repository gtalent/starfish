/*
   Copyright 2011-2012 gtalent2@gmail.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package graphics

/*
#cgo LDFLAGS: -lSDL -lSDL_ttf
#include "SDL/SDL.h"
#include "SDL/SDL_ttf.h"

SDL_Surface* openDisplay(int w, int h) {
	return SDL_SetVideoMode(w, h, 32, SDL_DOUBLEBUF | SDL_HWACCEL);
}

SDL_Surface* openDisplayFullscreen(int w, int h) {
	return SDL_SetVideoMode(w, h, 32, SDL_DOUBLEBUF | SDL_HWACCEL | SDL_FULLSCREEN);
}
*/
import "C"
import (
	"time"
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

var screen *C.SDL_Surface
var displayTitle string
var drawers []*canvasHolder
var displayDead chan interface{}
var running bool

//Sets the title of the window.
func SetDisplayTitle(title string) {
	displayTitle = title
	if screen != nil {
		C.SDL_WM_SetCaption(C.CString(displayTitle), C.CString(""))
	}
}

//Returns the title of this window.
func GetDisplayTitle() string {
	return displayTitle
}

func DisplayWidth() int {
	return int(screen.w)
}

func DisplayHeight() int {
	return int(screen.h)
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
		var d Drawer = drawFunc(drawer)
		if a.drawer == d {
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
		C.SDL_Flip(screen)
		time.Sleep(16000000)
	}
	displayDead <- nil
}

//Opens a window.
//Returns an indicator of success.
func OpenDisplay(width, height int, fullscreen bool) bool {
	if C.SDL_Init(C.SDL_INIT_VIDEO) != 0 {
		return false
	}
	C.TTF_Init()
	if fullscreen {
		screen = C.openDisplayFullscreen(C.int(width), C.int(height))
	} else {
		screen = C.openDisplay(C.int(width), C.int(height))
	}
	if screen == nil {
		return false
	}
	running = true
	C.SDL_WM_SetCaption(C.CString(displayTitle), C.CString(""))
	C.SDL_GL_SetAttribute(C.SDL_GL_SWAP_CONTROL, 1)
	go run()
	return true
}

//Closes the window.
func CloseDisplay() {
	if screen != nil {
		displayDead = make(chan interface{})
		running = false
		<-displayDead
		C.SDL_FreeSurface(screen)
		screen = nil
		C.SDL_Quit()
	}
}
