/*
   Copyright 2011-2012 gtalent2@gmail.com

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
package gfx

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

var screen *C.SDL_Surface
var autorun = true
var displayTitle string
var drawers []*canvasHolder
var displayDead chan interface{}
var kill = make(chan interface{})
var drawInterval = 0

//An interface used to for telling the display what to draw.
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

//Returns the width of the display window.
func DisplayWidth() int {
	return int(screen.w)
}

//Returns the height of the display window.
func DisplayHeight() int {
	return int(screen.h)
}

//Adds a drawer object to run when the screen draws.
func AddDrawer(drawer Drawer) {
	ch := new(canvasHolder)
	ch.drawer = drawer
	ch.canvas = newCanvas(screen)
	drawers = append(drawers, ch)
}

//Removes the given drawer object.
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

//Adds a draw function to call when the screen draws.
func AddDrawFunc(drawer func(*Canvas)) {
	ch := new(canvasHolder)
	ch.drawer = drawFunc(drawer)
	ch.canvas = newCanvas(screen)
	drawers = append(drawers, ch)
}

//Removes the given draw function.
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
	for autorun {
		select {
		case <-kill:
			break
		default:
			Draw()
		}
		time.Sleep(time.Duration(drawInterval))
	}
}

//Sets whether or not the draw functions will be called automatically. On by default.
func SetDraw(autodraw bool) {
	autorun = autodraw
}

//Sets the time in milliseconds between draws when autodraw is on.
func SetDrawInterval(ms int) {
	drawInterval = ms * 1000000
}

//Used to manually draw the screen.
func Draw() {
	for _, a := range drawers {
		a.canvas.pane = screen
		a.canvas.load()
		a.drawer.Draw(&a.canvas)
	}
	C.SDL_Flip(screen)
}

//Opens a window.
//Returns an indicator of success.
func OpenDisplay(width, height int, fullscreen bool) bool {
	if C.SDL_Init(C.SDL_INIT_VIDEO) != 0 {
		return false
	}
	C.TTF_Init()
	var flags C.Uint32 = C.SDL_DOUBLEBUF
	flags |= C.SDL_SWSURFACE
	flags |= C.SDL_HWACCEL

	if fullscreen {
		screen = C.openDisplayFullscreen(C.int(width), C.int(height))
	} else {
		screen = C.openDisplay(C.int(width), C.int(height))
	}
	if screen == nil {
		return false
	}
	C.SDL_WM_SetCaption(C.CString(displayTitle), C.CString(""))
	C.SDL_GL_SetAttribute(C.SDL_GL_SWAP_CONTROL, 1)
	SetDrawInterval(16)
	go run()
	return true
}

//Closes the window.
func CloseDisplay() {
	if screen != nil {
		close(kill)
		C.SDL_FreeSurface(screen)
		screen = nil
		C.SDL_Quit()
	}
}
