/*
   Copyright 2011-2012 starfish authors

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
	starfish "../"
	b "../plumbing"
)

//Used to draw and to hold data for the drawing context.
type Canvas struct {
	viewport    viewport
	color       Color
	translation starfish.Point
	origin      starfish.Point
}

func newCanvas() (p Canvas) {
	p.viewport = newViewport()
	return
}

//Loads the settings for this Pane onto the SDL Surface.
func (me *Canvas) load() {
	me.viewport.calcBounds()
	r := me.viewport.bounds()
	b.SetClipRect(r.X, r.Y, r.Width, r.Height)
}

//Returns the bounds of this Canvas
func (me *Canvas) GetViewport() starfish.Bounds {
	return me.viewport.bounds()
}

//Pushes a viewport to limit the drawing space to the given bounds within the current drawing space.
func (me *Canvas) PushViewport(x, y, width, height int) {
	me.origin.SubtractFrom(me.viewport.translate())
	me.viewport.push(starfish.Bounds{starfish.Point{X: int(x), Y: int(y)}, starfish.Size{Width: int(width), Height: int(height)}})
	r := me.viewport.bounds()
	b.SetClipRect(r.X, r.Y, r.Width, r.Height)
	me.origin = me.translation.AddOf(me.viewport.bounds().Point)
	me.origin.AddTo(me.viewport.translate())
}

//Exits the current viewport, unless there is no viewport.
func (me *Canvas) PopViewport() {
	if me.viewport.pt != 0 {
		me.origin.SubtractFrom(me.viewport.translate())
		me.viewport.pop()
		r := me.viewport.bounds()

		if r.Width == -1 {
			r.Width = b.DisplayWidth()
		}
		if r.Height == -1 {
			r.Height = b.DisplayHeight()
		}

		b.SetClipRect(r.X, r.Y, r.Width, r.Height)
		me.origin = me.translation.AddOf(me.viewport.bounds().Point)
		me.origin.AddTo(me.viewport.translate())
	}
}

//Sets the color that the Canvas will draw with.
func (me *Canvas) SetRGB(r, g, b byte) {
	me.color = Color{Red: r, Green: g, Blue: b, Alpha: 255}
}

//Sets the color that the Canvas will draw with.
func (me *Canvas) SetRGBA(r, g, b, a byte) {
	me.color = Color{Red: r, Green: g, Blue: b, Alpha: a}
}

//Sets the color that the Canvas will draw with.
func (me *Canvas) SetColor(color Color) {
	me.color = color
}

//Fills a rounded rectangle at the given coordinates and size on this Canvas.
func (me *Canvas) FillRoundedRect(x, y, width, height, radius int) {
	x += me.origin.X
	y += me.origin.Y
	b.FillRoundedRect(x, y, width, height, radius, me.color.bColor())
}

//Fills a rectangle at the given coordinates and size on this Canvas.
func (me *Canvas) FillRect(x, y, width, height int) {
	x += me.origin.X
	y += me.origin.Y
	b.FillRect(x, y, width, height, me.color.bColor())
}

//Draws the text at the given coordinates.
func (me *Canvas) DrawText(text *Text, x, y int) {
	b.DrawImage(text.text, x, y)
}

//Draws the image at the given coordinates.
func (me *Canvas) DrawAnimation(animation *Animation, x, y int) {
	me.DrawImage(animation.GetImage(), x, y)
}

//Draws the image at the given coordinates.
func (me *Canvas) DrawImage(img *Image, x, y int) {
	b.DrawImage(img.img, x, y)
}
