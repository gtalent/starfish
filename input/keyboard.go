/*
   Copyright 2011-2014 starfish authors

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

var keyPressListenersLock sync.Mutex
var keyPressListeners []KeyPressListener

var keyReleaseListenersLock sync.Mutex
var keyReleaseListeners []KeyReleaseListener

func AddKeyPressFunc(listener func(key KeyEvent)) {
	AddKeyPressListener(genericKeyListener(listener))
}

func RemoveKeyPressFunc(listener func(key KeyEvent)) {
	RemoveKeyPressListener(genericKeyListener(listener))
}

func AddKeyReleaseFunc(listener func(key KeyEvent)) {
	AddKeyReleaseListener(genericKeyListener(listener))
}

func RemoveKeyReleaseFunc(listener func(key KeyEvent)) {
	RemoveKeyReleaseListener(genericKeyListener(listener))
}

func AddKeyPressListener(listener KeyPressListener) {
	keyPressListenersLock.Lock()
	keyPressListeners = append(keyPressListeners, listener)
	keyPressListenersLock.Unlock()
}

func RemoveKeyPressListener(listener KeyPressListener) {
	keyPressListenersLock.Lock()
	pt := 0
	for i, v := range keyPressListeners {
		if v == listener {
			pt = i
			keyPressListeners[pt] = keyPressListeners[len(keyPressListeners)-1]
			keyPressListeners = keyPressListeners[:len(keyPressListeners)-1]
			keyPressListenersLock.Unlock()
			break
		}
	}
}

func AddKeyReleaseListener(listener KeyReleaseListener) {
	keyReleaseListenersLock.Lock()
	keyReleaseListeners = append(keyReleaseListeners, listener)
	keyReleaseListenersLock.Unlock()
}

func RemoveKeyReleaseListener(listener KeyReleaseListener) {
	keyReleaseListenersLock.Lock()
	pt := 0
	for i, v := range keyReleaseListeners {
		if v == listener {
			pt = i
			keyReleaseListeners[pt] = keyReleaseListeners[len(keyReleaseListeners)-1]
			keyReleaseListeners = keyReleaseListeners[:len(keyReleaseListeners)-1]
			keyReleaseListenersLock.Unlock()
			break
		}
	}
}
