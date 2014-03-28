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
