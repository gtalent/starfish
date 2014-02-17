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
package input

import b "github.com/gtalent/starfish/plumbing"

//Initializes the input system and returns a bool indicating success.
func Init() {
	b.QuitFunc = func() {
		quitListenersLock.Lock()
		for _, v := range quitListeners {
			go v.Quit()
		}
		quitListenersLock.Unlock()
	}
	b.KeyDown = func(e b.KeyEvent) {
		var ke KeyEvent
		ke.Key = e.Key
		keyPressListenersLock.Lock()
		for _, v := range keyPressListeners {
			go v.KeyPress(ke)
		}
		keyPressListenersLock.Unlock()
	}
	b.KeyUp = func(e b.KeyEvent) {
		var ke KeyEvent
		ke.Key = e.Key
		keyReleaseListenersLock.Lock()
		for _, v := range keyReleaseListeners {
			go v.KeyRelease(ke)
		}
		keyReleaseListenersLock.Unlock()
	}
	b.MouseWheelFunc = func(i b.MouseWheelEvent) {
		var e MouseWheelEvent
		e.Up = i.Up
		e.X = i.X
		e.Y = i.Y
		mouseWheelListenersLock.Lock()
		for _, v := range mouseWheelListeners {
			go v.MouseWheelScroll(e)
		}
		mouseWheelListenersLock.Unlock()
	}
	b.MouseButtonDown = func(e b.MouseEvent) {
		var me MouseEvent
		me.X = e.X
		me.Y = e.Y
		me.Button = e.Button
		mousePressListenersLock.Lock()
		for _, v := range mousePressListeners {
			go v.MouseButtonPress(me)
		}
		mousePressListenersLock.Unlock()
	}
	b.MouseButtonUp = func(e b.MouseEvent) {
		var me MouseEvent
		me.Button = e.Button
		me.X = e.X
		me.Y = e.Y
		mouseReleaseListenersLock.Lock()
		for _, v := range mouseReleaseListeners {
			go v.MouseButtonRelease(me)
		}
		mouseReleaseListenersLock.Unlock()
	}
	go b.HandleInput()
}
