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
package graphics

/*
#cgo LDFLAGS: -lSDL -lSDL_gfx -lSDL_image
#include "SDL/SDL.h"
#include "SDL/SDL_gfxPrimitives.h"
#include "SDL/SDL_rotozoom.h"
#include "SDL/SDL_image.h"

*/
import "C"
import (
	"../util"
)

//Used to draw and to hold data for the drawing context.
type Canvas struct {
	viewport    viewport
	pane        *C.SDL_Surface
	color       Color
	translation util.Point
	origin      util.Point
}

func newCanvas(surface *C.SDL_Surface) (p Canvas) {
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
	r := toSDL_Rect(me.viewport.Bounds)
	C.SDL_SetClipRect(me.pane, &r)
}

//Returns the bounds of this Canvas
func (me *Canvas) GetViewport() util.Bounds {
	return me.viewport.Bounds
}

//Pushs a viewport to limit the drawing space to the given bounds within the current drawing space.
func (me *Canvas) PushViewport(x, y, width, height int) {
	me.viewport.push(util.Bounds{util.Point{x, y}, util.Size{width, height}})
	r := toSDL_Rect(me.viewport.Bounds)
	C.SDL_SetClipRect(me.pane, &r)
	me.origin = me.translation.AddOf(me.viewport.Point)
}

//Exits the current viewport, unless there is no viewport.
func (me *Canvas) PopViewport() {
	if me.viewport.pt != 0 {
		me.viewport.pop()
		r := toSDL_Rect(me.viewport.Bounds)
		C.SDL_SetClipRect(me.pane, &r)
		me.origin = me.translation.AddOf(me.viewport.Point)
	}
}

//Sets the color that the Canvas will draw with.
func (me *Canvas) SetRGB(r, g, b byte) {
	me.color = Color{Red: r, Green: g, Blue: b}
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
	r := sdl_Rect(x+me.origin.X, y+me.origin.Y, width, height)
	C.roundedBoxRGBA(screen, C.Sint16(r.x), C.Sint16(r.y), C.Sint16(int(r.x)+int(r.w)), C.Sint16(int(r.y)+int(r.h)), C.Sint16(radius), C.Uint8(me.color.Red), C.Uint8(me.color.Green), C.Uint8(me.color.Blue), C.Uint8(me.color.Alpha))
}

//Fills a rectangle at the given coordinates and size on this Canvas.
func (me *Canvas) FillRect(x, y, width, height int) {
	r := sdl_Rect(x+me.origin.X, y+me.origin.Y, width, height)
	C.boxRGBA(screen, C.Sint16(r.x), C.Sint16(r.y), C.Sint16(int(r.x)+int(r.w)), C.Sint16(int(r.y)+int(r.h)), C.Uint8(me.color.Red), C.Uint8(me.color.Green), C.Uint8(me.color.Blue), C.Uint8(me.color.Alpha))
}

//Draws the text at the given coordinates.
func (me *Canvas) DrawText(text *Text, x, y int) {
	var dest C.SDL_Rect
	dest.x = C.Sint16(x + me.origin.X)
	dest.y = C.Sint16(y + me.origin.Y)
	C.SDL_BlitSurface(text.text, nil, me.pane, &dest)
}

//Draws the image at the given coordinates.
func (me *Canvas) DrawAnimation(animation *Animation, x, y int) {
	me.DrawImage(animation.GetImage(), x, y)
}

//Draws the image at the given coordinates.
func (me *Canvas) DrawImage(img *Image, x, y int) {
	C.SDL_SetAlpha(img.img, C.SDL_SRCALPHA, 255)
	var dest C.SDL_Rect
	dest.x = C.Sint16(x + me.origin.X)
	dest.y = C.Sint16(y + me.origin.Y)
	C.SDL_BlitSurface(img.img, nil, me.pane, &dest)
}
