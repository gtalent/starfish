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
	"sdl"
)

type imageLoader struct{}

var images = newResourceCatalog(func(path string) (interface{}, bool) {
	i := sdl.Load(path)
	return i, i != nil
})

type Image struct {
	img  *sdl.Surface
	path string
}

//Loads the image at the given path, or nil if the image was not found.
func LoadImage(path string) (img *Image) {
	images.checkout <- path
	i := (<-images.checkout).(*sdl.Surface)
	if i != nil {
		img = new(Image)
		img.img = i
		img.path = path
	}
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
	images.checkin <- me.path
	me.img = nil
	me.path = ""
}

func (me *Image) Resize(width, height int) {
	if me.img.W == 0 || me.img.H == 0 {
		return
	}
	bpp := me.img.Format.BitsPerPixel
	flags := me.img.Flags
	rmask := me.img.Format.Rmask
	gmask := me.img.Format.Gmask
	bmask := me.img.Format.Bmask
	amask := me.img.Format.Amask
	r := sdl.CreateRGBSurface(flags, width, height, int(bpp), rmask, gmask, bmask, amask)
	xstretch := float64(width) / float64(me.img.W)
	ystretch := float64(height) / float64(me.img.H)
	e1 := int(me.img.H)
	e2 := int(me.img.W)
	e3 := int(ystretch)
	e4 := int(xstretch)
	for oy := 0; oy < e1; oy++ {
		for ox := 0; ox < e2; ox++ {
			for ny := 0; ny < e3; ny++ {
				for nx := 0; nx < e4; nx++ {
					xp := int(xstretch * float64(ox)) + nx
					yp := int(ystretch * float64(oy)) + ny
					color := me.img.At(ox, oy)
					r.Set(xp, yp, color)
				}
			}
		}
	}
	me.img = r
}
