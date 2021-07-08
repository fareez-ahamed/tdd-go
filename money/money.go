package money

type Dollar struct {
	value int
}

func (d *Dollar) Times(n int) {
	d.value *= n
}
