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
package graphics

import (
	"dog/base/util"
	"sdl"
)

func toSDL_Rect(b util.Bounds) sdl.Rect {
	var r sdl.Rect
	r.X = int16(b.X)
	r.Y = int16(b.Y)
	r.W = uint16(b.Width)
	r.H = uint16(b.Height)
	return r
}

func sdl_Rect(x, y, width, height int) sdl.Rect {
	return sdl.Rect{int16(x), int16(y), uint16(width), uint16(height)}
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
