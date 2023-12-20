package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p *Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func main() {
	//1 example
	p := Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Println(p)
	//2 example
	r := &Point{2, 3}
	r.ScaleBy(2)
	fmt.Println(*r)
	//3 example
	v := Point{4, 5}
	pptr := &v
	pptr.ScaleBy(2)
	fmt.Println(v)
}
