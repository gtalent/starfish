/*
   Copyright 2011-2014 starfish authors

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

import (
	p "github.com/gtalent/starfish/plumbing"
	"time"
)

var displayTitle string
var drawers []*canvasHolder
var displayDead chan interface{}
var drawInterval = 0
var running = false

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
	p.SetDisplayTitle(title)
}

//Returns the title of this window.
func GetDisplayTitle() string {
	return displayTitle
}

//Returns the width of the display window.
func DisplayWidth() int {
	return p.DisplayWidth()
}

//Returns the height of the display window.
func DisplayHeight() int {
	return p.DisplayHeight()
}

//Adds a drawer object to run when the screen draws.
func AddDrawer(drawer Drawer) {
	ch := new(canvasHolder)
	ch.drawer = drawer
	ch.canvas = newCanvas()
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
	ch.canvas = newCanvas()
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

//Sets the time in milliseconds between draws when autodraw is on.
func SetDrawInterval(ms int) {
	drawInterval = ms * 1000000
}

//Used to manually draw the screen.
func Draw() {
	p.Draw()
}

//Opens a window.
//Returns an indicator of success.
func OpenDisplay(w, h int, fullscreen bool) bool {
	if !running {
		p.SetDrawFunc(func() {
			for _, a := range drawers {
				a.canvas.load()
				a.drawer.Draw(&a.canvas)
			}
		})
		p.OpenDisplay(w, h, fullscreen)
		running = true
		SetDrawInterval(16)
		startAnimTick()
	}
	return true
}

//Closes the window.
func CloseDisplay() {
	running = false
	p.CloseDisplay()
}

//Blocks until CloseDisplay is called, regardless of whether or not OpenDisplay has been called.
func Main() {
	go func() {
		for running {
			p.Draw()
			time.Sleep(time.Duration(drawInterval))
		}
	}()
	p.HandleEvents()
}
