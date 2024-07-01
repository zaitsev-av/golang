package lesson3

import (
	"fmt"
	"strconv"
)

// ссылки на ячейки памяти
// b := &a говорим что b ссылается на переменную а
// *b = "Go" присваиваем новое значение для ячейки памяти на которую ссылается b

func Get() {
	lesson3()
}

func lesson3() {
	a := "Hello World"
	b := &a
	*b = "Go"
	//fmt.Println(a)
	//pointer()
	//arrs()
	//slices()
}

type Point struct {
	X int
	Y int
	S string
}

// создаем метод для структуры Point
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

// массивы
func arrs() {
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"
	fmt.Println(arr)
	//numbers := [...]int{1, 2, 3} // инициализируем массив с элементами (...) -> говорят сообщаю компилятору что он должен сам определить размерность массива
	//fmt.Println(numbers)
}

// слайсы
func slices() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println(slice)
	// make ключевое слово для инициализации пустого массива
	newSlice := make([]string, 3)
	fmt.Println(fmt.Sprintf("len: %d", len(newSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(newSlice)))
	for i := 0; i < len(newSlice); i++ {
		if i == 3 {
			break
		}
		newSlice[i] = "slice " + strconv.Itoa(i)
	}
	newSlice = append(newSlice, "4")
	//fmt.Println(newSlice, "newSlice")
	fmt.Println(fmt.Sprintf("len: %d", len(newSlice)))
	fmt.Println(fmt.Sprintf("cap: %d", cap(newSlice)))

}
