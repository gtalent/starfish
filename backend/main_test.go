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
package backend

import (
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	OpenDisplay(800, 600, false)
	SetDisplayTitle("Narf!")
	img := LoadImage("../example/box.png")
	p := 0
	SetDrawFunc(func() {
		p++
		SetClipRect(0, 0, DisplayWidth(), DisplayHeight())
		c := Color{0, 0, 100, 255}
		FillRect(p, p, p+100, p+100, c)
		FillRoundedRect(100, 100, 400, 400, 2, c)
		DrawImage(img, p, p)
	})

	for {
		Draw()
		time.Sleep(15000000)
	}
	CloseDisplay()
}
