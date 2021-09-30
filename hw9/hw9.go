package main

type Point struct {
	x, y int
}

func (p Point) getX() {
	println(p.x)
}

func (p *Point) setX(n int) {
	p.x = n
}

func main() {
	var p Point = Point{1, 1}
	p.setX(3)
	p.getX()
}
