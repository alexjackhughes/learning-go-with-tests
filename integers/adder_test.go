package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assert := func(t *testing.T, got, want int) {
		t.Helper()

		if got != want {
			t.Errorf("expected %q but got %q", want, got)
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

func TestMinus(t *testing.T) {
	assert := func(t *testing.T, got, want int) {
		t.Helper()

		if got != want {
			t.Errorf("Expected %d but we received %d.", want, got)
		}
	}

	t.Run("4 - 1 = 3", func(t *testing.T) {
		got := Subtract(4, 1)
		want := 3
		assert(t, got, want)
	})

	t.Run("1 - 5 = -4", func(t *testing.T) {
		got := Subtract(1, 5)
		want := -4
		assert(t, got, want)
	})

	t.Run("60 - 13 = 47", func(t *testing.T) {
		got := Subtract(60, 13)
		want := 47
		assert(t, got, want)
	})
}

func ExampleSubtract() {
	sum := Subtract(10, 5)
	fmt.Println(sum)
	// Output: 5
}
