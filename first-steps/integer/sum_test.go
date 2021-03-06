package integer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should return a sum of slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("\ngot: %d\nwant: %d\ngiven: %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot: %v\nwant: %v\n", got, want)
	}
}

func TestSumAllRest(t *testing.T) {
	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	}

	t.Run("should sum something slices", func(t *testing.T) {
		got := SumAllRest([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSum(t, got, want)
	})

	t.Run("should sum empty slices", func(t *testing.T) {
		got := SumAllRest([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkSum(t, got, want)
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}
