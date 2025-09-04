package main
import (
	"fmt"
)
type Point struct {
	X float64
	Y float64
}

func NewPoint(x float64, y float64) *Point{
	return &Point{X: x, Y: y}
}

func (p *Point) Distance(otherPoint *Point) float64 {
	x := (p.X - otherPoint.X) * (p.X - otherPoint.X)
	y := (p.Y - otherPoint.Y) * (p.Y - otherPoint.Y)
	return (x + y) * 0.5
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)
	fmt.Println(p1.Distance(p2))
}
