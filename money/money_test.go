package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := &Dollar{5}
	five.Times(2)
	if five.value != 10 {
		t.Errorf("Expected %d, Got %d", 10, five.value)
	}
}
