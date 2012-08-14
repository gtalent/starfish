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
	"encoding/json"
	b "github.com/gtalent/starfish/backend"
	"github.com/gtalent/starfish/util"
)

type imageLabel struct {
	Str      string
	FilePath bool
}

type imageKey struct {
	Label  imageLabel
	Angle  float64
	Width  int
	Height int
}

func (me *imageKey) String() string {
	str, _ := json.Marshal(me)
	return string(str)
}

var images = newFlyweight(
	func(me *flyweight, path key) interface{} {
		key := path.(*imageKey)
		var i *b.Image
		var k imageKey
		if key.Label.FilePath {
			i = b.LoadImage(key.Label.Str)
		} else {
			json.Unmarshal([]byte(key.Label.Str), &k)
			i = me.checkout(&k).(*b.Image)
		}
		var w, h int
		if key.Width == -1 {
			w = int(i.W())
		} else {
			w = key.Width
		}
		if key.Height == -1 {
			h = int(i.H())
		} else {
			h = key.Height
		}
		if (i != nil) && (w != int(i.W()) || h != int(i.H()) || key.Angle != 0) {
			i = b.ResizeAngleOf(i, key.Angle, w, h)
		}
		return i
	},
	func(me *flyweight, path key, img interface{}) {
		i := img.(*b.Image)
		b.FreeImage(i)
	})

type Image struct {
	img *b.Image
	key imageKey
}

//Loads the image at the given path, or nil if the image was not found.
func LoadImage(path string) *Image {
	return LoadImageSize(path, -1, -1)
}

//Loads the image at the given path at the given angle, or nil if the image was not found.
func LoadImageAngle(path string, angle float64) *Image {
	return LoadImageSizeAngle(path, -1, -1, angle)
}

//Loads the image at the given path at the given size, or nil if the image was not found.
func LoadImageSize(path string, width, height int) *Image {
	return LoadImageSizeAngle(path, width, height, 0)
}

//Loads the image at the given path at the given angle and at the given size, or nil if the image was not found.
func LoadImageSizeAngle(path string, w, h int, angle float64) (img *Image) {
	var key imageKey
	key.Label.FilePath = true
	key.Label.Str = path
	key.Angle = angle
	key.Width = w
	key.Height = h
	i := images.checkout(&key).(*b.Image)
	img = new(Image)
	img.img = i
	img.key = key
	return
}

//Returns the width of the image.
func (me *Image) Width() int {
	return int(me.img.W())
}

//Returns the height of the image.
func (me *Image) Height() int {
	return int(me.img.H())
}

//Returns a util.Size object representing the size of this Image.
func (me *Image) Size() util.Size {
	var s util.Size
	s.Width = me.Width()
	s.Height = me.Height()
	return s
}

//Returns a unique string that can be used to identify the values of this Image.
func (me *Image) String() string {
	return me.key.String()
}

//Returns the path to the image on the disk.
func (me *Image) Path() string {
	return me.key.Label.Str
}

//Returns a version of this Image at the given angle and given size.
func (me *Image) ReSizeAngleOf(w, h int, angle float64) *Image {
	var key imageKey
	key.Label.FilePath = false
	key.Label.Str = me.key.String()
	key.Angle = angle
	key.Width = w
	key.Height = h
	i := images.checkout(&key).(*b.Image)
	img := new(Image)
	img.img = i
	img.key = key
	return img
}

//Returns a version of this Image at the given angle.
func (me *Image) ReangleOf(angle float64) *Image {
	return me.ReSizeAngleOf(me.key.Width, me.key.Height, angle)
}

//Returns a version of this Image at the given size.
func (me *Image) ResizeOf(w, h int) *Image {
	return me.ReSizeAngleOf(w, h, me.key.Angle)
}

//Nils this image and lets the resource manager know this object is no longer using the image data.
func (me *Image) Free() {
	images.checkin(&me.key)
	me.img = nil
	me.key.Label.Str = ""
}
