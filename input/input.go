/*
   Copyright 2011-2012 gtalent2@gmail.com

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

/*
#cgo LDFLAGS: -lSDL
#include "SDL/SDL.h"

int eventType(SDL_Event *e) {
	return e->type;
}

int eventKey(SDL_Event *e) {
	return e->key.keysym.sym;
}

int eventMouseButton(SDL_Event *e) {
	return e->button.button;
}

int eventMouseX(SDL_Event *e) {
	return e->button.x;
}

int eventMouseY(SDL_Event *e) {
	return e->button.y;
}
*/
import "C"

var running = true

//Initializes the input system and returns a bool indicating success.
func Init() {
	go run()
}

func run() {
	scrollFunc := func(b bool, x, y int) {
		var e MouseWheelEvent
		e.Up = b
		e.X = x
		e.Y = y
		mouseWheelListenersLock.Lock()
		for _, v := range mouseWheelListeners {
			go v.MouseWheelScroll(e)
		}
		mouseWheelListenersLock.Unlock()
	}
	for running {
		var e C.SDL_Event
		C.SDL_WaitEvent(&e)
		switch C.eventType(&e) {
		case C.SDL_QUIT:
			go func() {
				quitListenersLock.Lock()
				for _, v := range quitListeners {
					go v.Quit()
				}
				quitListenersLock.Unlock()
			}()
			running = false
		case C.SDL_KEYDOWN:
			go func() {
				var ke KeyEvent
				ke.Key = int(C.eventKey(&e))
				keyPressListenersLock.Lock()
				for _, v := range keyPressListeners {
					go v.KeyPress(ke)
				}
				keyPressListenersLock.Unlock()
			}()
		case C.SDL_KEYUP:
			go func() {
				var ke KeyEvent
				ke.Key = int(C.eventKey(&e))
				keyReleaseListenersLock.Lock()
				for _, v := range keyReleaseListeners {
					go v.KeyRelease(ke)
				}
				keyReleaseListenersLock.Unlock()
			}()
		case C.SDL_MOUSEBUTTONDOWN:
			x := int(C.eventMouseX(&e))
			y := int(C.eventMouseY(&e))
			switch C.eventMouseButton(&e) {
			case C.SDL_BUTTON_WHEELUP:
				go scrollFunc(false, x, y)
			case C.SDL_BUTTON_WHEELDOWN:
				go scrollFunc(true, x, y)
			default:
				go func() {
					var me MouseEvent
					me.X = x
					me.Y = y
					me.Button = int(C.eventMouseButton(&e))
					mousePressListenersLock.Lock()
					for _, v := range mousePressListeners {
						go v.MouseButtonPress(me)
					}
					mousePressListenersLock.Unlock()
				}()
			}
		case C.SDL_MOUSEBUTTONUP:
			x := int(C.eventMouseX(&e))
			y := int(C.eventMouseY(&e))
			switch C.eventMouseButton(&e) {
			case C.SDL_BUTTON_WHEELUP:
				go scrollFunc(false, x, y)
			case C.SDL_BUTTON_WHEELDOWN:
				go scrollFunc(true, x, y)
			default:
				go func() {
					var me MouseEvent
					me.Button = int(C.eventMouseButton(&e))
					me.X = int(C.eventMouseX(&e))
					me.Y = int(C.eventMouseY(&e))
					mouseReleaseListenersLock.Lock()
					for _, v := range mouseReleaseListeners {
						go v.MouseButtonRelease(me)
					}
					mouseReleaseListenersLock.Unlock()
				}()
			}
		}
	}
}
