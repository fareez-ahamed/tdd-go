package money

type Dollar struct {
	Value int
}

func (d *Dollar) Times(n int) *Dollar {
	return &Dollar{d.Value * n}
}

func (d *Dollar) Equals(a *Dollar) bool {
	return d.Value == a.Value
}
