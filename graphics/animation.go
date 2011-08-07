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
	"time"
)

//A type to automatically flip through a series of images.
type Animation struct {
	interval   int64
	lastUpdate int64
	slide      int
	images     []*Image
}

//Returns a string that can be used to identify the values of this Animation.
func (me *Animation) String() string {
	retval := strconv.Itoa64(me.interval)
	retval += "\n" + strconv.Itoa64(me.lastUpdate)
	retval += "\n" + strconv.Itoa(me.slide)
	for _, i := range me.images {
		retval += "\n" + i.String()
	}
	return retval
}

//Sets the number of milliseconds per image.
func (me *Animation) SetInterval(ms int) {
	me.interval = int64(ms) * 1000000
}

//Gets the current image.
func (me *Animation) GetImage() *Image {
	if me.images == nil {
		return nil
	}
	if time.Nanoseconds()-me.lastUpdate >= me.interval {
		slides := len(me.images)
		me.slide += int((time.Nanoseconds() - me.lastUpdate) / me.interval)
		me.slide -= (me.slide / slides) * slides
		me.lastUpdate = time.Nanoseconds()
	}
	return me.images[me.slide]
}

//Returns the image at the given index.
func (me *Animation) At(i int) *Image {
	return me.images[i]
}

//Returns the number of images in this Animation.
func (me *Animation) Size() int {
	return len(me.images)
}

func (me *Animation) LoadImage(path string) {
	if i := LoadImage(path); i != nil {
		me.images = append(me.images, i)
	}
}

func (me *Animation) LoadImageSize(path string, width, height int) {
	if i := LoadImageSize(path, width, height); i != nil {
		me.images = append(me.images, i)
	}
}

//Frees this Animations images, rendering it useless.
func (me *Animation) Free() {
	for _, a := range me.images {
		a.Free()
	}
}
