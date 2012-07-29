package backend

import "testing"

func TestMain(t *testing.T) {
	OpenDisplay()
	SetDisplayTitle("Narf!")
	Main()
}
