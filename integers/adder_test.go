package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assert := func(t *testing.T, got, want int) {
		t.Helper()

		if got != want {
			t.Errorf("Expected %d but we received %d.", want, got)
		}
	}

	t.Run("2 + 2 = 4", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assert(t, got, want)
	})

	t.Run("3 + 7 = 4", func(t *testing.T) {
		got := Add(3, 7)
		want := 10
		assert(t, got, want)
	})

	t.Run("1 + -4 = -3", func(t *testing.T) {
		got := Add(1, -4)
		want := -3
		assert(t, got, want)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
