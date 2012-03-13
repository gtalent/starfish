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
#cgo LDFLAGS: -lSDL -lSDL_ttf
#include "SDL/SDL.h"
#include "SDL/SDL_ttf.h"
*/
import "C"
import (
	"strconv"
)

type fontKey struct {
	path string
	size int
}

func (me *fontKey) String() string {
	return me.path + strconv.Itoa(me.size)
}

var fonts = newFlyweight(
	func(key key) interface{} {
		k := key.(*fontKey)
		font := C.TTF_OpenFont(C.CString(k.path), C.int(k.size))
		return font
	},
	func(path key, val interface{}) {
		C.TTF_CloseFont(val.(*C.TTF_Font))
	})

//A drawable representation of a string.
type Text struct {
	color Color
	text  *C.SDL_Surface
}

func (me *Text) Free() {
	C.SDL_FreeSurface(me.text)
}

//Returns a Color object representing the color of the text.
func (me *Text) Color() Color {
	return me.color
}

//Returns the width of this text.
func (me *Text) Width() int {
	return int(me.text.w)
}

//Returns the height of this text.
func (me *Text) Height() int {
	return int(me.text.h)
}

//A font type that represents a TTF file loaded from storage, used to create Text objects for drawing.
type Font struct {
	key   fontKey
	size  int
	font  *C.TTF_Font
	color Color
}

//Loads the TrueType Font at the given path, or nil if the font was not found.
func LoadFont(path string, size int) (font *Font) {
	var key fontKey
	key.path = path
	key.size = size

	if f := fonts.checkout(&key); f != nil {
		font = new(Font)
		font.font = f.(*C.TTF_Font)
		font.key = key
	}
	return
}

//Sets the color that this Font will draw with.
func (me *Font) SetColor(color Color) {
	me.color = color
}

//Sets the color that this Font will draw with.
func (me *Font) SetRGB(red, green, blue byte) {
	me.color.Red = red
	me.color.Green = green
	me.color.Blue = blue
}

//Loads text into the Text object passed in.
//Returns true if successful, false otherwise.
func (me *Font) Write(text string, t *Text) bool {
	t.color = me.color
	t.text = C.TTF_RenderText_Blended(me.font, C.CString(text), me.color.toSDL_Color())
	return t.text != nil
}

//Returns the size of this font.
func (me *Font) Size() int {
	return int(me.size)
}

//Returns the path to the font on the disk.
func (me *Font) Path() string {
	return me.key.path
}

//Nils this font and lets the resource manager know this object is no longer using the font data.
func (me *Font) Free() {
	fonts.checkin(&me.key)
	me.font = nil
	me.size = 0
	me.key.path = ""
}
