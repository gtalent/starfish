package backend

import "testing"

func TestMain(t *testing.T) {
	OpenDisplay(800, 600, false)
	SetDisplayTitle("Narf!")
}
