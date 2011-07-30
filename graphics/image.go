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
	"sdl"
)

type imageKey struct {
	path   string
	width  int
	height int
}

func (me *imageKey) String() string {
	return me.path + strconv.Itoa(me.width) + strconv.Itoa(me.height)
}

//Reads images from the disk and caches them.
var imageFiles = newResourceCatalog(func(path resourceKey) (interface{}, bool) {
	key := path.(*imageKey)
	i := sdl.Load(key.path)
	return i, i != nil
})

//Resizes images from imageFiles and caches them.
var images = newResourceCatalog(func(path resourceKey) (interface{}, bool) {
	key := path.(*imageKey)
	i := imageFiles.checkout(path).(*sdl.Surface)
	if i != nil {
		i = resize(i, key.width, key.height)
	}
	return i, i != nil
})

type Image struct {
	img  *sdl.Surface
	path string
}

//Loads the image at the given path, or nil if the image was not found.
func LoadImage(path string) (img *Image) {
	var key imageKey
	key.path = path
	i := imageFiles.checkout(&key).(*sdl.Surface)
	img = new(Image)
	img.img = i
	img.path = path
	return
}

//Loads the image at the given path, or nil if the image was not found.
func LoadImageSize(path string, width, height int) (img *Image) {
	var key imageKey
	key.path = path
	key.width = width
	key.height = height
	i := images.checkout(&key).(*sdl.Surface)
	img = new(Image)
	img.img = i
	img.path = path
	return
}

//Returns the width of the image.
func (me *Image) Width() int {
	return int(me.img.W)
}

//Returns the height of the image.
func (me *Image) Height() int {
	return int(me.img.H)
}

//Returns the path to the image on the disk.
func (me *Image) Path() string {
	return me.path
}

//Nils this image and lets the resource manager know this object is no longer using the image data.
func (me *Image) Free() {
	images.checkin(&imageKey{path: me.path, width: me.Width(), height: me.Height()})
	me.img = nil
	me.path = ""
}

func resize(img *sdl.Surface, width, height int) *sdl.Surface {
	if img.W == 0 || img.H == 0 {
		return nil
	}
	bpp := img.Format.BitsPerPixel
	flags := img.Flags
	rmask := img.Format.Rmask
	gmask := img.Format.Gmask
	bmask := img.Format.Bmask
	amask := img.Format.Amask
	r := sdl.CreateRGBSurface(flags, width, height, int(bpp), rmask, gmask, bmask, amask)
	xstretch := float64(width) / float64(img.W)
	ystretch := float64(height) / float64(img.H)
	e1 := float64(img.H)
	e2 := float64(img.W)
	e3 := float64(ystretch)
	e4 := float64(xstretch)
	for oy := float64(0); oy < e1; oy++ {
		for ox := float64(0); ox < e2; ox++ {
			for ny := float64(0); ny < e3; ny++ {
				for nx := float64(0); nx < e4; nx++ {
					xp := int(xstretch*float64(ox) + nx)
					yp := int(ystretch*float64(oy) + ny)
					color := img.At(int(ox), int(oy))
					r.Set(xp, yp, color)
				}
			}
		}
	}
	return r
}
