package pairutil

type Pairs struct {
	Pairs []Pair
}

type Pair struct {
	Left  any
	Right any
}

func (p Pairs) GetFirstByLeft(left any) (any, bool) {
	for _, val := range p.Pairs {
		if val.Left == left {
			return val.Right, true
		}
	}

	return nil, false
}
