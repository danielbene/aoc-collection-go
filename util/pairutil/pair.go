package pairutil

// looks good, but does not really useful :)
// keeping it for reference
type Pair[T comparable] struct {
	Left  T
	Right T
}

type Pairs[T comparable] struct {
	Pairs []Pair[T]
}

func (p *Pairs[T]) Put(left T, right T) {
	p.Pairs = append(p.Pairs, Pair[T]{Left: left, Right: right})
}

func (p Pairs[T]) GetLefts() []T {
	var lefts []T
	for _, pair := range p.Pairs {
		lefts = append(lefts, pair.Left)
	}

	return lefts
}

func (p Pairs[T]) GetFirstByLeft(left T) (T, bool) {
	for _, val := range p.Pairs {
		if val.Left == left {
			return val.Right, true
		}
	}

	var defValue T
	return defValue, false
}
