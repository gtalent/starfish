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
package gfx

import (
	"testing"
)

type rsrcKey string

func (me rsrcKey) String() string {
	return string(me)
}

func TestResourceManager(t *testing.T) {
	outKey := ""
	inKey := ""
	inVal := 0
	rsrcs := newFlyweight(func(me *flyweight, key key) interface{} {
		outKey = key.String()
		return 42
	}, func(me *flyweight, key key, val interface{}) {
		inKey = key.String()
		inVal = val.(int)
	})
	v := rsrcs.checkout(stringKey("Narf!"))
	rsrcs.checkin(stringKey("Narf!"))
	if v != 42 {
		t.Error("flyweight.checkout does not return the right value.")
	}
	if outKey != "Narf!" {
		t.Error("flyweight does not recieve the right key to load.")
	}
	if inKey != "Narf!" {
		t.Error("flyweight does not recieve the right key to delete.")
	}
	if inVal != 42 {
		t.Error("flyweight does not recieve the right value to delete.")
	}
}
