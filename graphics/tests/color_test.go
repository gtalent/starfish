package graphics

import (
	"testing"
)

type Color_Uint32Test struct {
	color  Color
	output uint32
}

func TestUint32(t *testing.T) {
	tests := make([]Color_Uint32Test, 0)
	tests = append(tests, Color_Uint32Test{Color{0, 0, 0}, 0})
	tests = append(tests, Color_Uint32Test{Color{255, 0, 0}, 255})
	tests = append(tests, Color_Uint32Test{Color{0, 255, 0}, 255 << 8})
	tests = append(tests, Color_Uint32Test{Color{0, 0, 255}, 255 << 16})
	tests = append(tests, Color_Uint32Test{Color{255, 255, 255}, 255 | (255 << 8) | (255 << 16)})

	for _, a := range tests {
		result := a.color.toUint32()
		if result != a.output {
			t.Errorf("Color (%d, %d, %d) gives %d.", a.color.Red, a.color.Green, a.color.Blue, result)
		}
	}
}
