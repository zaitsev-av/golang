package main

import "fmt"

func main() {

	//const a string = "Hello World"
	//fmt.Println(a)
	//
	//fmt.Println(test())
	//fmt.Println(test1())
	//s, s2 := test()
	//s3 := test1()
	//fmt.Println(s, s2, s3)
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	sum2 := 0
	for sum2 < 100 { //запись аналогична while
		sum2 += 5
	}
	//fmt.Println(sum)
	//fmt.Println(sum2)

	//fmt.Println(test2())
	//fmt.Println(test3(2))
}

func test() (a, b string) {
	a = "Hello"
	b = "World"
	return a, b
}

func test1() string {
	a := "Hello"
	return a
}

func test2() string {
	a := 0
	var name string
	for a < 150 {
		a += 20
		if a == 100 {
			fmt.Println("a=", a)
			name = "Hello Go!"
		} else {
			fmt.Printf("a is not 100 a=%d", a)
		}

	}
	return name
}

func test3(a int) (bool, int) {
	if a == 2 {
		return true, a
	}
	return false, a
}

func test4() {
	//defer означает отложенное действие
	defer fmt.Println("world")
	// работает по принципу LIFO
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("hello")

}
