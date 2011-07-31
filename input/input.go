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


type inputManager struct {
	quitListeners []func()
	mice          [100]mouse
}

func (me *inputManager) run() {
	for {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch et := e.(type) {
			case *sdl.QuitEvent:
				for _, a := range me.quitListeners {
					go a()
				}
			case *sdl.MouseButtonEvent:
			}
		}
		time.Sleep(6000000)
	}
}
