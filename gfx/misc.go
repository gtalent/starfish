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
package gfx

import (
	l "log"
	"os"
)

var log = l.New(os.Stdout, "  LOG: starfish backend: ", l.Ldate|l.Ltime)
var errlog = l.New(os.Stderr, "ERROR: starfish backend: error: ", l.Ldate|l.Ltime)

func LogOn(on bool) {
	if on {
		log = l.New(os.Stdout, "  LOG: starfish backend: ", l.Ldate|l.Ltime)
	} else {
		log = l.New(nil, "  LOG: starfish backend: ", l.Ldate|l.Ltime)
	}
}

func ErrLogOn(on bool) {
	if on {
		errlog = l.New(os.Stdout, "ERROR: starfish backend: ", l.Ldate|l.Ltime)
	} else {
		errlog = l.New(nil, "ERROR: starfish backend: ", l.Ldate|l.Ltime)
	}
}
