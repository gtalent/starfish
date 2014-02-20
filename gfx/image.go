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
	starfish "github.com/gtalent/starfish"
	b "github.com/gtalent/starfish/plumbing"
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
	img     *b.Image
	key     imageKey
	size    starfish.Size
	srcBnds starfish.Bounds
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
	img.ResetClipRect()
	return
}

func (me *Image) SetSize(w, h int) {
	me.size = starfish.Size{w, h}
}

func (me *Image) ResetSize() {
	me.size = starfish.Size{me.DefaultWidth(), me.DefaultHeight()}
}

//Resets a default image source clip rect to the full image.
func (me *Image) ResetClipRect() {
	me.srcBnds = starfish.Bounds{starfish.Point{0, 0}, me.DefaultSize()}
}

//Sets a default image source clip rect.
func (me *Image) SetClipRect(x, y, w, h int) {
	me.srcBnds = starfish.Bounds{starfish.Point{x, y}, starfish.Size{w, h}}
}

func (me *Image) clipX() int {
	return me.srcBnds.X
}

func (me *Image) clipY() int {
	return me.srcBnds.Y
}

func (me *Image) clipW() int {
	return me.srcBnds.Width
}

func (me *Image) clipH() int {
	return me.srcBnds.Height
}

//Returns the width of the image.
func (me *Image) Width() int {
	return me.size.Width
}

//Returns the height of the image.
func (me *Image) Height() int {
	return me.size.Height
}

//Returns a starfish.Size object representing the size of this Image.
func (me *Image) Size() starfish.Size {
	return me.size
}

//Returns the width of the image.
func (me *Image) DefaultWidth() int {
	return int(me.img.W())
}

//Returns the height of the image.
func (me *Image) DefaultHeight() int {
	return int(me.img.H())
}

//Returns a starfish.Size object representing the size of this Image.
func (me *Image) DefaultSize() starfish.Size {
	var s starfish.Size
	s.Width = me.DefaultWidth()
	s.Height = me.DefaultHeight()
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
