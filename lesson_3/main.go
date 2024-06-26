package main

import (
	"fmt"
	"math"
)

// ссылки на ячейки памяти
// b := &a говорим что b ссылается на переменную а
// *b = "Go" присваиваем новое значение для ячейки памяти на которую ссылается b

func main() {
	a := "Hello World"
	b := &a
	*b = "Go"
	fmt.Println(a)
	pointer()
	//EquableTriangle(5, 12, 13)
}

// codewars
func EquableTriangle(a, b, c int) bool {
	var perimeter = float64(a + b + c)
	p := perimeter / 2
	square := math.Sqrt(p * (p - float64(a)*(p-float64(b)*(p-float64(c)))))
	//fmt.Println(square == perimeter)
	return square == perimeter
}

func combat(health, damage float64) float64 {
	result := health - damage
	if result < 0 {
		return 0.0
	}
	return result
}

type Point struct {
	X int
	Y int
	S string
}

func (p Point) method() {
	fmt.Println(p.X + p.Y)
}

func pointer() {
	p := Point{X: 1, Y: 2}
	//fmt.Println(p)
	//fmt.Println(p.Y)
	//p.X = 5
	//fmt.Println(p)
	//
	//g := &p
	//
	//fmt.Println(g.X, "ссылка на p")
	p.method()

	p2 := &p

	p2.method()
}
