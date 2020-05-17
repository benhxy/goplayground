package goleetcode

import (
	"fmt"
	"testing"
)

func TestZeroLengthInput(t *testing.T) {
	actual := SmallerNumbersThanCurrent(make([]int, 0))
	if len(actual) != 0 {
		t.Errorf("Expected empty array, actual %v", fmt.Sprint(actual))
	}
}

func TestRegularInput(t *testing.T) {
	input := []int{8, 1, 2, 2, 3}
	expected := []int{4, 0, 1, 1, 3}

	actual := SmallerNumbersThanCurrent(input)

	if !Equal(expected, actual) {
		t.Errorf("Expected %v, actual %v", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}

func TestSameNumbersInput(t *testing.T) {
	input := []int{1, 1, 1, 1, 1}
	expected := []int{0, 0, 0, 0, 0}

	actual := SmallerNumbersThanCurrent(input)

	if !Equal(expected, actual) {
		t.Errorf("Expected %v, actual %v", fmt.Sprint(expected), fmt.Sprint(actual))
	}
}
