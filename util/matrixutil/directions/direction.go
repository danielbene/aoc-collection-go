package directions

type Direction struct {
	X int
	Y int
}

var (
	UP         = Direction{X: 0, Y: -1}
	DOWN       = Direction{X: 0, Y: 1}
	LEFT       = Direction{X: -1, Y: 0}
	RIGHT      = Direction{X: 1, Y: 0}
	UP_LEFT    = Direction{X: -1, Y: -1}
	UP_RIGHT   = Direction{X: 1, Y: -1}
	DOWN_LEFT  = Direction{X: -1, Y: 1}
	DOWN_RIGHT = Direction{X: 1, Y: 1}
)

func GetDirections() []Direction {
	return []Direction{
		UP,
		DOWN,
		LEFT,
		RIGHT,
		UP_LEFT,
		UP_RIGHT,
		DOWN_LEFT,
		DOWN_RIGHT,
	}
}
