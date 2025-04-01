package bs

import (
	"testing"
)

func TestBSPresentLowerHalf(t *testing.T) {
	data := []int{14, 26, 36, 49, 56, 66, 80, 86, 87, 89}
	find := 49

	loc, found := BinarySearch(data, find)

	if loc != 3 || found != true {
		t.Error("Expected index = 3, and found = true but got index=", loc, "and found =", found)
	}

}

func TestBSPresentUpperHalf(t *testing.T) {
	data := []int{14, 26, 36, 49, 56, 66, 80, 86, 87, 89}
	find := 87

	loc, found := BinarySearch(data, find)

	if loc != 8 || found != true {
		t.Error("Expected index = 3, and found = true but got index=", loc, "and found =", found)
	}

}

func TestBSPresentIsMid(t *testing.T) {
	data := []int{14, 26, 36, 49, 56, 66, 80, 86, 87, 89}
	find := 56

	loc, found := BinarySearch(data, find)

	if loc != 4 || found != true {
		t.Error("Expected index = 3, and found = true but got index=", loc, "and found =", found)
	}

}

func TestBSNotFoundLower(t *testing.T) {
	data := []int{14, 26, 36, 49, 56, 66, 80, 86, 87, 89}
	find := 10

	loc, found := BinarySearch(data, find)

	if loc != 1 || found != false {
		t.Error("Expected index = 3, and found = true but got index=", loc, "and found =", found)
	}

}

func TestBSNotFoundAbove(t *testing.T) {
	data := []int{14, 26, 36, 49, 56, 66, 80, 86, 87, 89}
	find := 90

	loc, found := BinarySearch(data, find)

	if loc != 9 || found != false {
		t.Error("Expected index = 3, and found = true but got index=", loc, "and found =", found)
	}

}
