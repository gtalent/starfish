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

/*
#cgo LDFLAGS: -lSDL
#include "SDL/SDL.h"
*/
import "C"

//An RGB color representation.
type Color struct {
	Red, Green, Blue, Alpha byte
}

func (me *Color) toSDL_Color() C.SDL_Color  {
	return C.SDL_Color{C.Uint8(me.Red), C.Uint8(me.Green), C.Uint8(me.Blue), C.Uint8(me.Alpha)}
}

func (me *Color) toUint32() uint32 {
	return (uint32(me.Red) << 16) | (uint32(me.Green) << 8) | uint32(me.Blue)
}
