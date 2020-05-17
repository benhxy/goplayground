package goleetcode

import "testing"

func TestDestCitySaoPaulo(t *testing.T) {
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	actualDest := destCity(paths)

	if actualDest != "Sao Paulo" {
		t.Errorf("Expected %v, actual %v", "Sao Paulo", actualDest)
	}
}
func TestDestCityA(t *testing.T) {
	paths := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	actualDest := destCity(paths)

	if actualDest != "A" {
		t.Errorf("Expected %v, actual %v", "A", actualDest)
	}
}
