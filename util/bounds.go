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
package util

/*
  Represents location and size attributes.
*/
type Bounds struct {
	Point
	Size
}

//Compares the value of this Bounds to the value of the given Bounds.
func (me *Bounds) Equals(p Bounds) bool {
	return me.Point.Equals(p.Point) && me.Size.Equals(p.Size)
}

/*
  Sets this Bounds to the given coordinates and dimensions.
*/
func (me *Bounds) Set(x, y, width, height int) {
	me.X = x
	me.Y = y
	me.Width = width
	me.Height = height
}

/*
  Returns the x coordinate + the width.
*/
func (me *Bounds) X2() int {
	return me.X + me.Width
}

/*
  Returns the y coordinate + the height.
*/
func (me *Bounds) Y2() int {
	return me.Y + me.Height
}

//Returns X2 and Y2 in a Point.
func (me *Bounds) Point2() Point {
	return Point{me.X2(), me.Y2()}
}

//Returns true if the given Point is in this Bounds, false otherwise.
func (me *Bounds) ContainsPoint(p Point) bool {
	return p.X >= me.X && p.Y >= me.Y && p.X <= me.X2() && p.Y <= me.Y2()
}

//Returns true if the given Point is in this Bounds, false otherwise.
func (me *Bounds) Contains(x, y int) bool {
	return x >= me.X && y >= me.Y && x <= me.X2() && y <= me.Y2()
}

//Returns true if this Bounds intersects with the given Bounds.
func (me *Bounds) Intersects(b Bounds) bool {
	return me.ContainsPoint(me.Point) || me.Contains(b.X, b.Y2()) ||
		me.Contains(b.X2(), b.Y) || me.Contains(b.X2(), b.Y2())
}

func (me *Bounds) String() string {
	return "(" + me.Point.String() + ", " + me.Size.String() + ")"
}
