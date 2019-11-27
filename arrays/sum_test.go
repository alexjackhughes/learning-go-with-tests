package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assert := func(t *testing.T, got, want int) {
		t.Helper()

		if got != want {
			t.Errorf("Expected %d but we received %d.", want, got)
		}
	}

	t.Run("Collection of 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assert(t, got, want)
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		assert(t, got, want)
	})

	t.Run("Collection of none", func(t *testing.T) {
		numbers := []int{}

		got := Sum(numbers)
		want := 0

		assert(t, got, want)
	})

	t.Run("Collection of one", func(t *testing.T) {
		numbers := []int{1}

		got := Sum(numbers)
		want := 1

		assert(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	assert := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Sum multiple arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assert(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	assert := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Sum multiple array tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assert(t, got, want)
	})

	t.Run("Sum multiple array tails that are longer than one element", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 2}, []int{0, 9, 9})
		want := []int{4, 18}

		assert(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assert(t, got, want)
	})
}
