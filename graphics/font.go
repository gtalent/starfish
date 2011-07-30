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
	"strconv"
	"sdl/ttf"
)

type fontKey struct {
	path string
	size int
}

func (me *fontKey) String() string {
	return me.path + strconv.Itoa(me.size)
}

var fonts = newResourceCatalog(func(key resourceKey) (interface{}, bool) {
	k := key.(*fontKey)
	font := ttf.OpenFont(k.path, k.size)
	return font, font == nil
}, func(path resourceKey, val interface{}) {
	val.(*ttf.Font).Close()
})

type Font struct {
	path string
	size int
	font *ttf.Font
}

//Loads the TrueType Font at the given path, or nil if the font was not found.
func LoadFont(path string, size int) (font *Font) {
	var key fontKey
	key.path = path
	key.size = size
	i := fonts.checkout(&key).(*ttf.Font)
	font = new(Font)
	font.font = i
	font.path = path
	return
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
