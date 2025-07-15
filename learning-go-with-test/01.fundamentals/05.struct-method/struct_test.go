package structmethod

import (
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {
// 	t.Run("rectangle should return area value", func(t *testing.T) {
// 		rectangle := Rectangle{2.0, 2.0}
// 		got := Area(rectangle)
// 		want := 4.0

// 		assertCorrectMessage(t, got, want)
// 	})

// 	t.Run("circle should return area value", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := Area(circle)
// 		want := 314.1592653589793

// 		assertCorrectMessage(t, got, want)
// 	})
// }

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		checkArea(t, rectangle, 72.0)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})

// }

func TestArea(t *testing.T) {
	areaTestCases := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, testCase := range areaTestCases {
		got := testCase.shape.Area()
		want := testCase.want

		assertCorrectMessage(t, got, want)
	}
}

// func assertCorrectMessage(t testing.TB, got, want any) {
// 	// t.Helper() is needed to tell the test suite that this method is a helper. By doing this, when it fails, the line number reported will be in our function call rather than inside our test helper. This will help other developers track down problems more easily.
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %g want %g", got, want)
	}
}
