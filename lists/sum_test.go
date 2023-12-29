package lists

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("sum of any slices", func(t *testing.T) {
		list1 := []int{1, 2, 3, 4, 5}
		list2 := []int{6, 3, 2, 1, 5}

		got := SumAllTails(list1, list2)
		want := []int{14, 11}

		checkingSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		// Arrange
		list1 := []int{4, 2, 5, 1}
		var emptyList []int

		// Act
		got := SumAllTails(list1, emptyList)
		want := []int{8, 0}

		// Assert
		checkingSums(t, got, want)
	})
}

func checkingSums(t testing.TB, got, want []int) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
