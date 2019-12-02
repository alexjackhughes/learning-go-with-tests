package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	assert := func(t *testing.T, got, want float64) {
		t.Helper()

		if got != want {
			t.Errorf("Expected %g but we received %g.", want, got)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, areaTest := range areaTests {
		got := areaTest.shape.Area()
		want := areaTest.want

		assert(t, got, want)
	}

	t.Run("Calcuate the perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		assert(t, got, want)
	})

	t.Run("Calculate the area of a square", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		assert(t, got, want)
	})

	t.Run("Calculate the area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 30.0}
		got := rectangle.Area()
		want := 300.0

		assert(t, got, want)
	})

	t.Run("Calculate the area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		assert(t, got, want)
	})
}
