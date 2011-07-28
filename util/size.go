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

/*
  Used to track a Cartesian size.
*/
type Size struct {
	Width  int
	Height int
}

/*
  Sets the width attribute for this Size.
  Takes:
  	width - the new width of this Size
*/
func (me *Size) SetWidth(width int) {
	me.Width = width
}

/*
  Sets the width attribute for this Size.
  Takes:
  	height - the new height of this Size
*/
func (me *Size) SetHeight(height int) {
	me.Height = height
}

/*
  Sets the dimensions of this Size.
  Takes:
  	width - the new width of this Size
  	height - the new height of this Size
*/
func (me *Size) SetSize(width int, height int) {
	me.Width = width
	me.Height = height
}

/*
  Sets the dimensions of this Size according to the given Size.
  Takes:
  	size - the Size object representing the values this Size should take on
*/
func (me *Size) Set(size *Size) {
	me.SetSize(size.Width, size.Height)
}

/*
  Gets the width attribute of this Size.
  Returns:
  	the width attribute of this Size
*/
func (me *Size) GetWidth() int {
	return me.Width
}

/*
  Gets the height attribute of this Size.
  Returns:
  	the height attribute of this Size
*/
func (me *Size) GetHeight() int {
	return me.Height
}

/*
  Gets the width and height of this Size.
  Returns:
 	width - the width attribute of this Size
 	height - the height attribute of this Size
*/
func (me *Size) Get() (int, int) {
	return me.Width, me.Height
}
