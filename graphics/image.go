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

type imageLoader struct {}

func (me *imageLoader) load(path string) (interface{}, bool) {
	i := sdl.Load(path)
	return i, i != nil
}

var images = newResourceCatalog(new(imageLoader))

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

//Nils this image and lets the resource manager know this object is no longer using the image data.
func (me *Image) Free() {
	images.checkin <- me.img
	me.img = nil
	me.path = ""
}
