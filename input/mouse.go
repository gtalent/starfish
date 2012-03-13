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

import "sync"

var mousePressListenersLock sync.Mutex
var mousePressListeners []MouseButtonPressListener

var mouseReleaseListenersLock sync.Mutex
var mouseReleaseListeners []MouseButtonReleaseListener

func AddMousePressFunc(listener func(MouseEvent)) {
	AddMousePressListener(genericMouseListener(listener))
}

func RemoveMousePressFunc(listener func(MouseEvent)) {
	RemoveMousePressListener(genericMouseListener(listener))
}

func AddMouseReleaseFunc(listener func(MouseEvent)) {
	AddMouseReleaseListener(genericMouseListener(listener))
}

func RemoveMouseReleaseFunc(listener func(MouseEvent)) {
	RemoveMouseReleaseListener(genericMouseListener(listener))
}

func AddMouseButtonListener(listener MouseButtonListener) {
	AddMousePressListener(listener)
	AddMouseReleaseListener(listener)
}

func RemoveMouseButtonListener(listener MouseButtonListener) {
	RemoveMousePressListener(listener)
	RemoveMouseReleaseListener(listener)
}

func AddMousePressListener(listener MouseButtonPressListener) {
	mousePressListenersLock.Lock()
	mousePressListeners = append(mousePressListeners, listener)
	mousePressListenersLock.Unlock()
}

func RemoveMousePressListener(listener MouseButtonPressListener) {
	mousePressListenersLock.Lock()
	pt := 0
	for i, v := range mousePressListeners {
		if v == listener {
			pt = i
			mousePressListeners[pt] = mousePressListeners[len(mousePressListeners)-1]
			mousePressListeners = mousePressListeners[:len(mousePressListeners)-1]
			mousePressListenersLock.Unlock()
			break
		}
	}
}

func AddMouseReleaseListener(listener MouseButtonReleaseListener) {
	mouseReleaseListenersLock.Lock()
	mouseReleaseListeners = append(mouseReleaseListeners, listener)
	mouseReleaseListenersLock.Unlock()
}

func RemoveMouseReleaseListener(listener MouseButtonReleaseListener) {
	mouseReleaseListenersLock.Lock()
	pt := 0
	for i, v := range mouseReleaseListeners {
		if v == listener {
			pt = i
			mouseReleaseListeners[pt] = mouseReleaseListeners[len(mouseReleaseListeners)-1]
			mouseReleaseListeners = mouseReleaseListeners[:len(mouseReleaseListeners)-1]
			mouseReleaseListenersLock.Unlock()
			break
		}
	}
}
