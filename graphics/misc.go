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
package graphics

/*
#cgo LDFLAGS: -lSDL
#include "SDL/SDL.h"
*/
import "C"
import (
	"github.com/gtalent/starfish/util"
)

func toSDL_Rect(b util.Bounds) C.SDL_Rect {
	var r C.SDL_Rect
	r.x = C.Sint16(b.X)
	r.y = C.Sint16(b.Y)
	r.w = C.Uint16(b.Width)
	r.h = C.Uint16(b.Height)
	return r
}

func sdl_Rect(x, y, width, height int) C.SDL_Rect {
	var r C.SDL_Rect
	r.x = C.Sint16(x)
	r.y = C.Sint16(y)
	r.w = C.Uint16(width)
	r.h = C.Uint16(height)
	return r
}

//Returns the difference between the two integers given.
func diff(i, ii int) int {
	if i < 0 {
		i = -i
	}
	if ii < 0 {
		ii = -ii
	}
	if i > ii {
		return i - ii
	}
	return ii - i
}
