/*
   Copyright 2011-2014 starfish authors

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
package starfish

import (
	p "github.com/gtalent/starfish/plumbing"
)

var running = false

//Blocks until CloseDisplay is called, regardless of whether or not OpenDisplay has been called.
func Main() {
	go func() {
		for running {
			p.Draw()
			//time.Sleep(time.Duration(drawInterval))
		}
	}()
	p.HandleEvents()
}