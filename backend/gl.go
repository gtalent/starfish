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
#cgo LDFLAGS: -lGL -lglut
#include "gl_freebsd_linux_darwin.h"
*/
import "C"

func OpenDisplay(w, h int, full bool) {
	var c C.int = 0
	C.glutInit(&c, nil)
	C.glutInitDisplayMode(C.GLUT_DEPTH | C.GLUT_DOUBLE | C.GLUT_RGBA)
	C.glutInitWindowSize(C.int(w), C.int(h))
	C.glutCreateWindow(C.CString(""))
}

func CloseDisplay() {
	C.glutExit()
}

//Sets the title of the window.
func SetDisplayTitle(title string) {
	C.glutSetWindowTitle(C.CString(title))
}

//Returns the width of the display window.
func DisplayWidth() int {
	return int(C.glutGet(C.GLUT_WINDOW_WIDTH))
}

//Returns the height of the display window.
func DisplayHeight() int {
	return int(C.glutGet(C.GLUT_WINDOW_HEIGHT))
}

//Used to manually draw the screen.
func Draw() {
	drawFunc()
}

//MISC
//An RGB color representation.
type Color struct {
	Red, Green, Blue, Alpha byte
}

func (me *Color) toUint32() uint32 {
	return (uint32(me.Red) << 16) | (uint32(me.Green) << 8) | uint32(me.Blue)
}

//GFX HANDLING

//Pushes a viewport to limit the drawing space to the given bounds within the current drawing space.
func SetClipRect(x, y, w, h int) {
	y += h
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(w), C.GLsizei(h))
	C.glLoadIdentity()
	C.glOrtho(0.0, C.GLdouble(w), C.GLdouble(h), 0.0, -1.0, 1.0)
	C.glTranslated(C.GLdouble(-x), C.GLdouble(-y), 0)
}

func FillRoundedRect(x, y, w, h, radius int, c Color) {
	C.glColor4b(C.GLbyte(c.Red), C.GLbyte(c.Green), C.GLbyte(c.Blue), C.GLbyte(c.Alpha))
}

func FillRect(x, y, w, h int, c Color) {
	C.glColor4b(C.GLbyte(c.Red), C.GLbyte(c.Green), C.GLbyte(c.Blue), C.GLbyte(c.Alpha))
	C.glRecti(C.GLint(x), C.GLint(y), C.GLint(x+w), C.GLint(y+h))
}

//Draws the image at the given coordinates.
func DrawImage(img *Image, x, y int) {
}

//IMAGE HANDLING

type Image struct {
}

func (me *Image) W() int {
	return 0
}

func (me *Image) H() int {
	return 0
}

func LoadImage(path string) *Image {
	return nil
}

func FreeImage(img *Image) {
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
