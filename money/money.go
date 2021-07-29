package money

type Dollar struct {
	value int
}

func (d *Dollar) Times(n int) Dollar {
	return Dollar{d.value * n}
}

func (d *Dollar) Equals(a *Dollar) bool {
	return d.value == a.value
}

func NewDollar(n int) Dollar {
	return Dollar{n}
}
