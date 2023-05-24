package gosnack

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p *Point) IsEqual(ep *Point) bool {
	return *p == *ep
}
