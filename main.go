package main

import "fmt"

// 構造体の定義
type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

// 面積を計算する関数の定義
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

// 出力する関数の定義
func printTriangleArea(t triangle) {
	fmt.Println(t.getArea())
}

func printSquareArea(s square) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{height: 3, base: 6}
	s := square{sideLength: 5}

	printTriangleArea(t)
	printSquareArea(s)
}
