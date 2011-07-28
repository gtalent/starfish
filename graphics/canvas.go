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
	"fmt"
	"dog/base/util"
	"sdl"
)

//Used to draw and to hold data for the drawing context.
type Canvas struct {
	viewport    viewport
	pane        *sdl.Surface
	color       uint32
	translation util.Point
	origin      util.Point
}

func newCanvas(surface *sdl.Surface) (p Canvas) {
	p.pane = surface
	p.viewport.X = 0
	p.viewport.Y = 0
	p.viewport.Width = 65000
	p.viewport.Height = 65000
	return
}

//Loads the settings for this Pane onto the SDL Surface.
func (me *Canvas) load() {
	me.viewport.calcBounds()
	b := me.viewport.Bounds
	r := toSDL_Rect(b)
	me.pane.SetClipRect(&r)
}

//Sets the drawing bounds of this Canvas on screen.
func (me *Canvas) SetBounds(x, y, width, height int) {
	me.viewport.X = 0
	me.viewport.Y = 0
	me.viewport.Width = 65000
	me.viewport.Height = 65000
}

//Pushs a viewport to limit the drawing space to the given bounds within the current drawing space.
func (me *Canvas) PushViewport(x, y, width, height int) {
	me.viewport.push(util.Bounds{util.Point{x, y}, util.Size{width, height}})
	me.viewport.calcBounds()
	b := me.viewport.Bounds
	r := toSDL_Rect(b)
	me.pane.SetClipRect(&r)
	me.origin = me.translation.AddOf(me.viewport.Point)
}

//Exits the current viewport, unless there is no viewport.
func (me *Canvas) PopViewport() {
	if me.viewport.pt != 0 {
		me.viewport.pop()
		me.viewport.calcBounds()
		b := me.viewport.Bounds
		r := toSDL_Rect(b)
		me.pane.SetClipRect(&r)
		me.origin = me.translation.AddOf(me.viewport.Point)
	}
}

//Sets the color that the Canvas will draw with.
func (me *Canvas) SetColor(color Color) {
	me.color = color.toUint32()
}

//Fills a rectangle at the given coordinates and size on this Canvas.
func (me *Canvas) FillRect(x, y, width, height int) {
	me.pane.FillRect(&sdl.Rect{int16(x + me.origin.X), int16(y + me.origin.Y), uint16(width), uint16(height)}, me.color)
}

//Draws the image at the given coordinates with the given dimensions.
func (me *Canvas) DrawImage(img *Image, x, y, width, height int) {
	var dest sdl.Rect
	dest.X = int16(x + me.origin.X)
	dest.Y = int16(y + me.origin.Y)
	fmt.Println(x, ", ", y)
	fmt.Println(dest.X, ", ", dest.Y)
	/*src := sdl_Rect(0, 0, width, height)
	 */
	me.pane.Blit(&dest, img.img, nil)
}
