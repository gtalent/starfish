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
)

type viewport struct {
	util.Bounds
	list [500]util.Bounds
	pt   uint
}

func newViewport() (v viewport) {
	v.pt = 0
	r := &v.list[0]
	r.X = 0
	r.Y = 0
	r.Width = 65000
	r.Height = 65000
	v.list[0].X = 0
	v.list[0].Y = 0
	v.list[0].Width = 65000
	v.list[0].Height = 65000
	return
}

//Sets the root drawing bounds of this viewport.
func (me *viewport) setBounds(x, y, width, height int) {
	me.list[0].X = 0
	me.list[0].Y = 0
	me.list[0].Width = 65000
	me.list[0].Height = 65000
}

func (me *viewport) push(rect util.Bounds) {
	me.list[me.pt] = rect
	me.pt++
	me.calcBounds()
}

func (me *viewport) pop() {
	if me.pt < 1 {
		return
	}
	me.pt--
	me.calcBounds()
}

func (me *viewport) calcBounds() {
	me.X = 0
	me.Y = 0
	me.Width = 65000
	me.Height = 65000
	for i := uint(0); i < me.pt; i++ {
		r := &me.list[i]
		nx1 := me.X + (r.X)
		ny1 := me.Y + (r.Y)
		wc := (r.Width)
		hc := (r.Height)
		ox2 := me.X + me.Width
		oy2 := me.Y + me.Height

		if nx1+wc > ox2 {
			me.Width = ox2 - nx1
		} else {
			me.Width = wc
		}
		me.X = nx1

		if ny1+hc > oy2 {
			me.Height = oy2 - ny1
		} else {
			me.Height = hc
		}
		me.Y = ny1
	}
}
