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

import "sync"

var quitListenersLock sync.Mutex
var quitListeners []QuitListener

type QuitListener interface {
	Quit()
}

type genericQuitListener func()

func (me genericQuitListener) Quit() {
	me()
}

//Adds a function to be called when the display is asked to close.
func AddQuitListenerFunc(listener func()) {
	AddQuitListener(genericQuitListener(listener))
}

//Removes the specified quit listener function.
func RemoveQuitListenerFunc(listener func()) {
	RemoveQuitListener(genericQuitListener(listener))
}

//Adds an interface to be called when the display is asked to close.
func AddQuitListener(listener QuitListener) {
	quitListenersLock.Lock()
	quitListeners = append(quitListeners, listener)
	quitListenersLock.Unlock()
}

//Removes the specified quit listener.
func RemoveQuitListener(listener QuitListener) {
	quitListenersLock.Lock()
	pt := 0
	for i, v := range quitListeners {
		if v == listener {
			pt = i
			break
		}
	}
	quitListeners[pt] = quitListeners[len(quitListeners)-1]
	quitListeners = quitListeners[:len(quitListeners)-1]
	quitListenersLock.Unlock()
}


