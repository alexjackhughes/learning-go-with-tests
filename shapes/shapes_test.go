package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	assert := func(t *testing.T, got, want float64) {
		t.Helper()

		if got != want {
			t.Errorf("Expected %.2f but we received %.2f.", want, got)
		}
	}

	t.Run("Calcuate the perimeter", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		assert(t, got, want)
	})

	t.Run("Calculate the area of a square", func(t *testing.T) {
		got := Area(10.0, 10.0)
		want := 100.0

		assert(t, got, want)
	})

	t.Run("Calculate the area of a rectangle", func(t *testing.T) {
		got := Area(10.0, 30.0)
		want := 300.0

		assert(t, got, want)
	})
}
