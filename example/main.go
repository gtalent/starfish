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
	"time"
	gfx "../graphics"
	"../input"
)

type Drawer struct {
	boxman *gfx.Image
	text   gfx.Text
}

func (me *Drawer) init() {
	me.boxman = gfx.LoadImageSize("dirt.png", 70, 70)
	font := gfx.LoadFont("LiberationSans-Bold.ttf", 32)
	if font != nil {
		font.SetRGB(0, 0, 255)
		font.Write("Narf!", &me.text)
		font.Free()
	} else {
		return
	}
	if me.boxman == nil {
		fmt.Println("Could not load boxman.")
		return
	}
}

func (me *Drawer) Draw(c *gfx.Canvas) {
	//clear screen
	c.SetColor(gfx.Color{Red: 0, Green: 0, Blue: 0})
	c.FillRect(0, 0, gfx.DisplayWidth(), gfx.DisplayHeight())

	c.SetColor(gfx.Color{Red: 0, Green: 0, Blue: 255, Alpha: 255})
	c.FillRect(42, 42, 100, 100)

	//draw boxman if he's not nil
	if me.boxman != nil {
		c.DrawImage(me.boxman, 200, 200)
		c.SetColor(gfx.Color{Red: 0, Green: 0, Blue: 0, Alpha: 100})
		c.FillRect(200, 200, 100, 100)
	}
	c.DrawText(&me.text, 400, 400)

	//draw a green rect in a viewport
	c.PushViewport(42, 42, 500, 500)
	{
		c.SetColor(gfx.Color{Red: 0, Green: 255, Blue: 0, Alpha: 127})
		c.FillRect(42, 42, 100, 100)
	}
	c.PopViewport()
}

func main() {
	//For a fullscreen at your screens native resolution, simply use this line instead:
	//if !gfx.OpenDisplay(0, 0, true) {
	if !gfx.OpenDisplay(800, 600, false) {
		return
	}

	input.Init()

	var pane Drawer
	pane.init()
	gfx.AddDrawFunc(func(c *gfx.Canvas) {
		pane.Draw(c)
	})
	running := true
	input.AddQuitFunc(func() {
		running = false
		pane.boxman.Free()
		gfx.CloseDisplay()
	})
	input.AddMouseWheelFunc(func(up input.MouseWheelEvent) {
		fmt.Println("Mouse wheel scrolling:", up)
	})
	input.AddKeyPressFunc(func(e input.KeyEvent) {
		if e.Key == input.Key_Escape {
			running = false
			pane.boxman.Free()
			gfx.CloseDisplay()
		}
	})
	input.AddMousePressFunc(func(e input.MouseEvent) {
		fmt.Println("Mouse Press!")
	})
	input.AddMouseReleaseFunc(func(e input.MouseEvent) {
		fmt.Println("Mouse Release!")
	})

	input.AddKeyPressFunc(func(i input.KeyEvent) {
		fmt.Println("Key Press!")
	})
	input.AddKeyReleaseFunc(func(i input.KeyEvent) {
		fmt.Println("Key Release!")
	})

	//read input
	for running {
		time.Sleep(6000000)
	}
}
