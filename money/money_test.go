package money_test

import (
	"testing"

	"github.com/fareez-ahamed/tdd/money"
)

func TestMultiplication(t *testing.T) {
	tt := []struct {
		dollar money.Dollar
		times  int
		want   money.Dollar
	}{
		{money.NewDollar(5), 2, money.NewDollar(10)},
		{money.NewDollar(5), 3, money.NewDollar(15)},
	}
	for _, ti := range tt {
		got := ti.dollar.Times(ti.times)
		if !ti.want.Equals(&got) {
			t.Errorf("Expected %v, Got %v", ti.want, got)
		}
	}
}

func TestEquality(t *testing.T) {
	tt := []struct {
		dollar1 money.Dollar
		dollar2 money.Dollar
		want    bool
	}{
		{money.NewDollar(5), money.NewDollar(5), true},
		{money.NewDollar(5), money.NewDollar(6), false},
	}
	for _, ti := range tt {
		got := ti.dollar1.Equals(&ti.dollar2)
		if got != ti.want {
			t.Errorf("Expected %t, Got %t", ti.want, got)
		}
	}
}
