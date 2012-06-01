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
	"github.com/gtalent/starfish/util"
)

type viewport struct {
	list         [500]util.Bounds
	translations [500]util.Point
	pt           uint
}

func newViewport() (v viewport) {
	v.pt = 0
	v.list[0].X = 0
	v.list[0].Y = 0
	v.list[0].Width = 65000
	v.list[0].Height = 65000
	return
}

func (me *viewport) translate() util.Point {
	return me.translations[me.pt]
}

func (me *viewport) bounds() util.Bounds {
	return me.list[me.pt]
}

func (me *viewport) push(rect util.Bounds) {
	me.pt++
	me.list[me.pt] = rect
	me.calcBounds()
}

func (me *viewport) pop() {
	if me.pt < 1 {
		return
	}
	me.pt--
}

func (me *viewport) calcBounds() {
	if me.pt == 0 {
		return
	}
	p := &me.list[me.pt-1]
	n := &me.list[me.pt]
	t := &me.translations[me.pt]
	*t = me.translations[me.pt-1]
	n.Point.AddTo(p.Point)
	//make sure the point of origin is not negative
	if n.X < p.X {
		t.X = n.X - p.X
		n.Width -= p.X - n.X
		n.X = p.X
		if n.Width < 0 {
			n.Width = 0
		}
	}
	if n.Y < p.Y {
		t.Y = n.Y - p.Y
		n.Height -= p.Y - n.Y
		n.Y = p.Y
		if n.Height < 0 {
			n.Height = 0
		}
	}
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
}
