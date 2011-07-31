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
package input

import (
	"time"
	"sdl"
)

type mouse struct {
	lastPress   int64 // the time of the last mousebutton press
	lastRelease int64 // the time of the last mousebutton release
	lastClick   int64 // the time of the last mouse click 
}

type mouseManager struct {
	mice                     [50]mouse
	mouseButtonDownListeners []func()
	mouseButtonUpListeners   []func()
	input chan *sdl.MouseButtonEvent
}

func (me *mouseManager) run() {
	et := <-me.input
	me.mice[et.Button].lastRelease = time.Nanoseconds()
	if time.Nanoseconds()-me.mice[et.Button].lastRelease < 1000000000 {
		//if click
		me.mice[et.Button].lastRelease = time.Nanoseconds()
	} else { //if non-click release
		for _, a := range me.mouseButtonUpListeners {
			go a()
		}
	}

}
