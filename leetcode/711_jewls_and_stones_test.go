package goleetcode

import "testing"

func TestNoJewelInStones(t *testing.T) {
	actual := numJewelsInStones("z", "ZZZ")
	if actual != 0 {
		t.Errorf("%v jewels expected, %v jewels found", 0, actual)
	}
}

func TestSomeJewelInStones(t *testing.T) {
	actual := numJewelsInStones("az", "aaaZzZ")
	if actual != 4 {
		t.Errorf("%v jewels expected, %v jewels found", 4, actual)
	}
}

func TestNilStones(t *testing.T) {
	actual := numJewelsInStones("az", "")
	if actual != 0 {
		t.Errorf("%v jewels expected, %v jewels found", 0, actual)
	}
}
