package gosnack

type Direction struct {
	X int
	Y int
}

func NewDirection(x, y int) *Direction {
	return &Direction{x, y}
}

func (d *Direction) IsDirection(ed *Direction) bool {
	return *d == *ed
}

var (
	LeftDirection  *Direction = NewDirection(-1, 0)
	RightDirection *Direction = NewDirection(1, 0)
	UpDirection    *Direction = NewDirection(0, -1)
	DownDirection  *Direction = NewDirection(0, 1)
)
