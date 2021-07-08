package money

type Dollar struct {
	Value int
}

func (d *Dollar) Times(n int) *Dollar {
	return &Dollar{d.Value * n}
}
