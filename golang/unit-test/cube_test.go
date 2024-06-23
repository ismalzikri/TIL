package main

import "testing"

var (
	cube             Cube    = Cube{4}
	volumeShouldbe   float64 = 64
	wideShouldbe     float64 = 96
	circularShouldbe float64 = 48
)

func TestCalculateVolume(t *testing.T) {
	t.Logf("Volume : %.2f", cube.Volume())

	if cube.Volume() != volumeShouldbe {
		t.Errorf("WRONG! should be %.2f", volumeShouldbe)
	}
}

func TestCalculateWide(t *testing.T) {
	t.Logf("Wide : %.2f", cube.Wide())

	if cube.Wide() != wideShouldbe {
		t.Errorf("WRONG! should be %.2f", wideShouldbe)
	}
}

func TestCalculateCircular(t *testing.T) {
	t.Logf("Circular : %.2f", cube.Circular())

	if cube.Circular() != circularShouldbe {
		t.Errorf("WRONG! should be %.2f", circularShouldbe)
	}
}

func BenchmarkCalculateWide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cube.Wide()
	}
}
