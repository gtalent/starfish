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
	//"fmt"
	"dog/base/util"
)

func TestViewportPushPop(t *testing.T) {
	viewport := newViewport()
	viewport.calcBounds()
	initial := viewport.GetBounds()
	tests := make([]util.Bounds, 0)
	tests = append(tests, util.Bounds{util.Point{42, 42}, util.Size{100, 100}})

	for _, test := range tests {
		viewport.push(test)
		if !viewport.Equals(test) {
			t.Errorf("viewport.push is broken")
		}
		viewport.pop()
		if !viewport.Equals(initial) {
			t.Error("viewport.pop is broken\n\tviewport is:\t\t", viewport.Bounds, "\n\tviewport should be:\t", initial)
		}
	}
}
