//https://www.sunapro.com/go-interface/
package main

import "fmt"

// shapeという型は必ずgetAreaメソッドを持っていることを保証している。
type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	t := triangle{height: 3, base: 6}
	s := square{sideLength: 5}

	printArea(t)
	printArea(s)
}
