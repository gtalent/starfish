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
	"testing"
)

type Color_Uint32Test struct {
	color  Color
	output uint32
}

func TestUint32(t *testing.T) {
	tests := make([]Color_Uint32Test, 0)
	tests = append(tests, Color_Uint32Test{Color{0, 0, 0}, 0})
	tests = append(tests, Color_Uint32Test{Color{255, 0, 0}, 255 << 16})
	tests = append(tests, Color_Uint32Test{Color{0, 255, 0}, 255 << 8})
	tests = append(tests, Color_Uint32Test{Color{0, 0, 255}, 255})
	tests = append(tests, Color_Uint32Test{Color{255, 255, 255}, 255 | (255 << 8) | (255 << 16)})

	for _, a := range tests {
		result := a.color.toUint32()
		if result != a.output {
			t.Errorf("Color (%d, %d, %d) gives %d.", a.color.Red, a.color.Green, a.color.Blue, result)
		}
	}
}
