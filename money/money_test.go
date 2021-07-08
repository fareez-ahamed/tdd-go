package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := &Dollar{5}
	dollar := five.Times(2)
	if dollar.Value != 10 {
		t.Errorf("Expected %d, Got %d", 10, five.Value)
	}
	dollar = five.Times(3)
	if dollar.Value != 15 {
		t.Errorf("Expected %d, Got %d", 15, five.Value)
	}
}
