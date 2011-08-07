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
package util

import "strconv"

/*
  Used to track a point on a 2 dimensional Cartesian plane.
*/
type Point struct {
	/*
	 * The x coordinate.
	 */
	X int
	/*
	 * The y coordinate.
	 */
	Y int
}

//Compares the value of this Point to the value of the given Point.
func (me *Point) Equals(p Point) bool {
	return me.X == p.X && me.Y == p.Y
}

/*
 Gets the x coordinate of this Point.
 Returns:
	the x coordinate of this Point
*/
func (me *Point) GetX() int {
	return me.X
}

/*
 Gets the y coordinate of this Point.
 Returns:
	the y coordinate of this Point
*/
func (me *Point) GetY() int {
	return me.Y
}

/*
 Gets the x and y coordinates of this Point.
 Returns:
	the x coordinate of this Point
	the y coordinate of this Point
*/
func (me *Point) Get() (int, int) {
	return me.X, me.Y
}

/*
  Sets the x coordinate of this Point.
  x - the new x coordinate for this point
*/
func (me *Point) SetX(x int) {
	me.Y = x
}

/*
  Sets the y coordinate of this Point.
  y - the new y coordinate for this point
*/
func (me *Point) SetY(y int) {
	me.Y = y
}

/*
  Sets the x coordinate of this Point.
  x - the new x coordinate for this point
  y - the new y coordinate for this point
*/
func (me *Point) SetPoint(x int, y int) {
	me.X = x
	me.Y = y
}

/*
  Sets this Point's values to that of the given Point.
  point - the Point who's value that this Point is to take
*/
func (me *Point) Set(point *Point) {
	me.SetPoint(point.X, point.Y)
}

// Returns the value of this Point with the coordinates of the given added to its coordinates.
func (me *Point) AddOf(point Point) (p Point) {
	p.X = me.X + point.X
	p.Y = me.Y + point.Y
	return
}

// Adds the coordinates of the given Point to the coordinates of this Point.
func (me *Point) AddTo(point Point) {
	me.X += point.X
	me.Y += point.Y
}

// Returns the value of this Point with the coordinates of the given subtracted from its coordinates.
func (me *Point) SubtractOf(point Point) (p Point) {
	p.X = me.X - point.X
	p.Y = me.Y - point.Y
	return
}

// Subtracts the coordinates of the given Point from the coordinates of this Point.
func (me *Point) SubtractFrom(point Point) {
	me.X -= point.X
	me.Y -= point.Y
}

// Returns the value of this Point with the coordinates of the given divided by its coordinates.
func (me *Point) DivideOf(point Point) (p Point) {
	p.X = me.X / point.X
	p.Y = me.Y / point.Y
	return
}

// Divides the coordinates of the given Point by the coordinates of this Point.
func (me *Point) DivideBy(point Point) {
	me.X /= point.X
	me.Y /= point.Y
}

// Returns the value of this Point with the coordinates of the given multiplied by its coordinates.
func (me *Point) MultiplyOf(point Point) (p Point) {
	p.X = me.X * point.X
	p.Y = me.Y * point.Y
	return
}

// Multiplies the coordinates of the given Point by the coordinates of this Point.
func (me *Point) MultiplyBy(point Point) {
	me.X *= point.X
	me.Y *= point.Y
}

// Returns this Point as a Size.
func (me *Point) ToSize() Size {
	return Size{me.X, me.Y}
}

func (me *Point) String() string {
	return "(" + strconv.Itoa(me.X) + strconv.Itoa(me.Y) + ")"
}
