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

type imageNode struct {
	uses int
	img  *sdl.Surface
}

type imageCatalog struct {
	images   map[string]*imageNode
	checkout, checkin chan interface{}
}

var images imageCatalog

func (me *imageCatalog) run() {
	for {
		select {
		case input := <-me.checkout:
			path := input.(string)
			i, ok := me.images[path]
			if ok {
				i.uses++
				me.checkout <- i.img
			} else {
				tmp := sdl.Load(path)
				if tmp != nil {
					i = new(imageNode)
					i.img = tmp
					i.uses++
					me.images[path] = i
					me.checkout <- i.img
				}
			}
		case input := <-me.checkin:
			path := input.(string)
			i, ok := me.images[path]
			if ok {
				i.uses--
				me.checkin <- true
			} else {
				me.checkin <- false
			}
		}
		me.checkout <- nil
	}
}

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
