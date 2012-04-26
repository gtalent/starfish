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
package main

import (
	"fmt"
	"../graphics"
	"../input"
)

type Drawer struct {
	boxman *graphics.Image
	text   graphics.Text
}

func (me *Drawer) init() bool {
	me.boxman = graphics.LoadImageSize("dirt.png", 70, 70)
	font := graphics.LoadFont("LiberationSans-Bold.ttf", 32)
	if font != nil {
		font.SetRGB(0, 0, 255)
		font.Write("Narf!", &me.text)
		font.Free()
	} else {
		return false
	}
	if me.boxman == nil {
		fmt.Println("Could not load boxman.")
		return false
	}
	return true
}

func (me *Drawer) Draw(c *graphics.Canvas) {
	//clear screen
	c.SetColor(graphics.Color{Red: 0, Green: 0, Blue: 0})
	c.FillRect(0, 0, graphics.DisplayWidth(), graphics.DisplayHeight())

	c.SetColor(graphics.Color{Red: 0, Green: 0, Blue: 255, Alpha: 255})
	c.FillRect(42, 42, 100, 100)

	//draw boxman if he's not nil
	if me.boxman != nil {
		c.DrawImage(me.boxman, 200, 200)
		c.SetColor(graphics.Color{Red: 0, Green: 0, Blue: 0, Alpha: 100})
		c.FillRect(200, 200, 100, 100)
	}
	c.DrawText(&me.text, 400, 400)

	//push a viewport at (42, 42)
	//Note: viewports may be nested
	c.PushViewport(42, 42, 500, 500)
	{
		//draw a green rect in a viewport
		c.SetColor(graphics.Color{Red: 0, Green: 255, Blue: 0, Alpha: 127})
		c.FillRect(42, 42, 100, 100)
	}
	c.PopViewport()
}

func main() {
	//For a fullscreen at your screens native resolution, simply use this line instead:
	//if !graphics.OpenDisplay(0, 0, true) {
	if !graphics.OpenDisplay(800, 600, false) {
		return
	}
	graphics.SetDisplayTitle("starfish example")

	input.Init()

	var pane Drawer
	if !pane.init() {
		return
	}
	graphics.AddDrawer(&pane)
	running := make(chan interface{})
	quit := func() {
		graphics.CloseDisplay()
		pane.boxman.Free()
		pane.text.Free()
		running <-nil
	}
	input.AddQuitFunc(quit)
	input.AddMouseWheelFunc(func(e input.MouseWheelEvent) {
		if e.Up {
			fmt.Println("Mouse wheel scrolling up.")
		} else {
			fmt.Println("Mouse wheel scrolling down.")
		}
	})
	input.AddMousePressFunc(func(e input.MouseEvent) {
		fmt.Println("Mouse Press!")
	})
	input.AddMouseReleaseFunc(func(e input.MouseEvent) {
		fmt.Println("Mouse Release!")
	})

	input.AddKeyPressFunc(func(e input.KeyEvent) {
		fmt.Println("Key Press!")
		if e.Key == input.Key_Escape {
			quit()
		}
	})
	input.AddKeyReleaseFunc(func(i input.KeyEvent) {
		fmt.Println("Key Release!")
	})

	<-running
}
