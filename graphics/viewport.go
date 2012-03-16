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

import (
	"github.com/gtalent/WombatCore/util"
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
	if me.pt == 0 {
		me.Bounds = me.list[0]
		return
	}
	p := &me.list[me.pt-1]
	n := &me.list[me.pt]
	n.Point.AddTo(p.Point)
	//make sure the point of origin is not beyond the edge
	if n.X > p.X2() {
		n.X = p.X2()
	}
	if n.Y > p.Y2() {
		n.Y = p.Y2()
	}
	//make sure the new edge is not beyond the old edge
	if n.X2() > p.X2() {
		n.Width = p.X2() - n.X
	}
	if n.Y2() > p.Y2() {
		n.Height = p.Y2() - n.Y
	}
	me.Bounds = *n
}
