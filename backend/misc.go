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

import "github.com/gtalent/starfish/util"

var drawFunc = func() {}

func SetDrawFunc(f func()) {
	drawFunc = f
}

var QuitFunc = func() {}
var KeyUp = func(e KeyEvent) {}
var KeyDown = func(e KeyEvent) {}
var MouseWheelFunc = func(e MouseWheelEvent) {}
var MouseButtonUp = func(e MouseEvent) {}
var MouseButtonDown = func(e MouseEvent) {}

type KeyEvent struct {
	Key int
}

type MouseEvent struct {
	util.Point
	Button int
}

type MouseWheelEvent struct {
	util.Point
	Up bool
}
