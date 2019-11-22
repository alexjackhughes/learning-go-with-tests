package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	}

	t.Run("repeat character eight times", func(t *testing.T) {
		got := Repeat("e", 8)
		want := "eeeeeeee"
		assert(t, got, want)
	})

	t.Run("repeat character 1 time", func(t *testing.T) {
		got := Repeat("a", 1)
		want := "a"
		assert(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	repeat := Repeat("b", 10)
	fmt.Println(repeat)
	// Output: bbbbbbbbbb
}
