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
package backend

/*
#cgo LDFLAGS: -lm -lGL -lGLU -lglfw -lIL

#include "capi.h"
#include "gl.h"

*/
import "C"
import "runtime"

var draw = make(chan interface{})
var drawRet = make(chan interface{})
var loadImg = make(chan string)
var loadImgRet = make(chan *Image)
var freeImg = make(chan *Image)
var freeImgRet = make(chan interface{})

func OpenDisplay(w, h int, full bool) {
	ret := make(chan interface{})
	go func() {
		runtime.LockOSThread()
		f := C.int(0)
		if full {
			f = 1
		}
		C.openDisplay(C.int(w), C.int(h), f)
		ret <- nil
		run()
		runtime.UnlockOSThread()
	}()
	<-ret
}

func CloseDisplay() {
	C.closeDisplay()
}

//Sets the title of the window.
func SetDisplayTitle(title string) {
	C.setDisplayTitle(C.CString(title))
}

//Returns the width of the display window.
func DisplayWidth() int {
	return int(C.displayWidth())
}

//Returns the height of the display window.
func DisplayHeight() int {
	return int(C.displayHeight())
}

//export Draw
func Draw() {
	draw <- nil
	<-drawRet
}

func run() {
	for {
		select {
		case <-draw:
			C.clear()
			drawFunc()
			C.flip()
			drawRet <- nil
		case path := <-loadImg:
			C.loadImage(C.CString(path))
			i := Image(C.loadImage(C.CString(path)))
			loadImgRet <- &i
		case img := <-freeImg:
			i := C.Image(*img)
			C.freeImage(&i)
			freeImgRet <- nil
		}
	}
}

//MISC
//An RGB color representation.
type Color C.Color

func (me *Color) toUint32() uint32 {
	return (uint32(me.Red) << 16) | (uint32(me.Green) << 8) | uint32(me.Blue)
}

//GFX HANDLING

//Pushes a viewport to limit the drawing space to the given bounds within the current drawing space.
func SetClipRect(x, y, w, h int) {
	C.setClipRect(C.int(x), C.int(y), C.int(w), C.int(h))
}

func FillRoundedRect(x, y, w, h, radius int, c Color) {
}

func FillRect(x, y, w, h int, c Color) {
	C.fillRect(C.int(x), C.int(y), C.int(w), C.int(h), C.Color(c))
}

//Draws the image at the given coordinates.
func DrawImage(img *Image, x, y int) {
	i := C.Image(*img)
	C.drawImage(&i, C.int(x), C.int(y))
}

//IMAGE HANDLING

type Image C.Image

func (me *Image) W() int {
	return int(me.w)
}

func (me *Image) H() int {
	return int(me.h)
}

func LoadImage(path string) *Image {
	loadImg <- path
	return <-loadImgRet
}

func FreeImage(img *Image) {
	freeImg <- img
	<-freeImgRet
}

func ResizeAngleOf(image *Image, angle float64, width, height int) *Image {
	return nil
}

//TEXT HANDLING

type Font struct {
}

func LoadFont(path string, size int) *Font {
	return nil
}

func FreeFont(val *Font) {
}

func (me *Font) WriteTo(text string, t *Image, c Color) bool {
	return false
}

//INPUT HANDLING

func HandleInput() {
}
