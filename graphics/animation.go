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

//A type to automatically flip through a series of images.
type Animation struct {
	interval int
	images   []*Image

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
