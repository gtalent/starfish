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
	"strconv"
	"sdl"
	"sdl/ttf"
)

type fontKey struct {
	path string
	size int
}

func (me *fontKey) String() string {
	return me.path + strconv.Itoa(me.size)
}

var fonts = newResourceCatalog(
	func(key resourceKey) (interface{}, bool) {
		k := key.(*fontKey)
		font := ttf.OpenFont(k.path, k.size)
		fmt.Println(sdl.GetError())
		return font, font != nil
	},
	func(path resourceKey, val interface{}) {
		val.(*ttf.Font).Close()
	})

//A drawable representation of a string.
type Text struct {
	color Color
	text  *sdl.Surface
}

//Returns a Color object representing the color of the text.
func (me *Text) Color() Color {
	return me.color
}

//Returns the width of this text.
func (me *Text) Width() int {
	return int(me.text.W)
}

//Returns the height of this text.
func (me *Text) Height() int {
	return int(me.text.H)
}

//A font type that represents a TTF file loaded from storage, used to create Text objects for drawing.
type Font struct {
	path  string
	size  int
	font  *ttf.Font
	color Color
}

//Loads the TrueType Font at the given path, or nil if the font was not found.
func LoadFont(path string, size int) (font *Font) {
	var key fontKey
	key.path = path
	key.size = size

	if f := fonts.checkout(&key); f != nil {
		font = new(Font)
		font.font = f.(*ttf.Font)
		font.path = path
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
	t.text = ttf.RenderText_Blended(me.font, text, me.color.toSDL_Color())
	return t.text != nil
}

//Returns the size of this font.
func (me *Font) Size() int {
	return int(me.size)
}

//Returns the path to the font on the disk.
func (me *Font) Path() string {
	return me.path
}

//Nils this font and lets the resource manager know this object is no longer using the font data.
func (me *Font) Free() {
	images.checkin(&fontKey{path: me.path, size: me.size})
	me.font = nil
	me.size = 0
	me.path = ""
}
