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

/*
#cgo LDFLAGS: -lSDL -lSDL_image -lSDL_gfx
#include "SDL/SDL.h"
#include "SDL/SDL_rotozoom.h"
#include "SDL/SDL_image.h"

*/
import "C"

import (
	"strconv"
)

type imageKey struct {
	path   string
	width  int
	height int
}

func (me *imageKey) String() string {
	return me.path + strconv.Itoa(me.width) + strconv.Itoa(me.height)
}

//Resizes images from imageFiles and caches them.
var images = newFlyweight(
	func(path key) interface{} {
		key := path.(*imageKey)
		i := C.IMG_Load(C.CString(key.path))
		if (i != nil) && (int(i.w) != key.width || int(i.h) != key.height) {
			i = resize(i, key.width, key.height)
		}
		return i
	},
	func(path key, img interface{}) {
		i := img.(*C.SDL_Surface)
		C.SDL_FreeSurface(i)
	})

type Image struct {
	img  *C.SDL_Surface
	key  imageKey
}

//Returns a unique string that can be used to identify the values of this Image.
func (me *Image) String() string {
	return me.key.String()
}

//Loads the image at the given path, or nil if the image was not found.
func LoadImage(path string) (img *Image) {
	var key imageKey
	key.path = path
	i := images.checkout(&key).(*C.SDL_Surface)
	img = new(Image)
	img.img = i
	img.key = key
	return
}

//Loads the image at the given path, or nil if the image was not found.
func LoadImageSize(path string, width, height int) (img *Image) {
	var key imageKey
	key.path = path
	key.width = width
	key.height = height
	i := images.checkout(&key).(*C.SDL_Surface)
	img = new(Image)
	img.img = i
	img.key = key
	return
}

//Returns the width of the image.
func (me *Image) Width() int {
	return int(me.img.w)
}

//Returns the height of the image.
func (me *Image) Height() int {
	return int(me.img.h)
}

//Returns the path to the image on the disk.
func (me *Image) Path() string {
	return me.key.path
}

//Nils this image and lets the resource manager know this object is no longer using the image data.
func (me *Image) Free() {
	images.checkin(&imageKey{path: me.key.path, width: me.Width(), height: me.Height()})
	me.img = nil
	me.key.path = ""
}

func resize(img *C.SDL_Surface, width, height int) *C.SDL_Surface {
	if img.w == 0 || img.h == 0 {
		return nil
	}
	xstretch := C.double(float64(width) / float64(img.w))
	ystretch := C.double(float64(height) / float64(img.h))
	retval := C.zoomSurface(img, xstretch, ystretch, 1)
	C.SDL_FreeSurface(img)
	return retval
}
