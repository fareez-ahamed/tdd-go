package money

import "testing"

func TestMultiplication(t *testing.T) {
	tt := []struct {
		dollar *Dollar
		times  int
		want   int
	}{
		{&Dollar{5}, 2, 10},
		{&Dollar{5}, 3, 15},
	}
	for _, ti := range tt {
		got := ti.dollar.Times(ti.times)
		if ti.want != got.Value {
			t.Errorf("Expected %d, Got %d", ti.want, got)
		}
	}
}

func TestEquality(t *testing.T) {
	tt := []struct {
		dollar1 *Dollar
		dollar2 *Dollar
		want    bool
	}{
		{&Dollar{5}, &Dollar{5}, true},
		{&Dollar{5}, &Dollar{6}, false},
	}
	for _, ti := range tt {
		got := ti.dollar1.Equals(ti.dollar2)
		if got != ti.want {
			t.Errorf("Expected %t, Got %t", ti.want, got)
		}
	}
}
