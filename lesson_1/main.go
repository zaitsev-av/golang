package main

import (
	"fmt"
)

const variable string = "asdfasdf"

//var number float64 = 222.76
//var one int8      // -128 -> 127 1 byte
//var two int16     // -32768 -> 32767 2 byte
//var four int32    // 2 милиарда 4byte
//var eigth int64   // пинтиллионы 8byte
//var a uint8       // 0 -> 255 1byte
//var b uint16      // 0 -> 65535 2byte
//var c uint32      // 0 -> 4bil 4byte
//var d uint64      // 0 -> 18pentillion 8byte
//var e byte        // uint8 1byte
//var f rune        // uint16 2byte
//var g int         // целое число, в зависимости от платформы может занимать либо 4 либо 8 байт int32 || int64
//var h uint        // целое беззнаковое число, в зависимости от платформы может занимать либо 4 либо 8 байт uint32 || uint64
//var i uint        // uint8 1byte
//var flo32 float32 // 1.4*10^-45 -> 3.4*10^38 число с плавающей точкой 4байта
//var flo64 float64 //  число с плавающей точкой которое больше числа выше 8байт
//// комплексные числа
//var comp64 complex64
//var comp128 complex128

func main() {
	//message := "Hello world"
	//var message2 string
	//message2 = "Golang"
	//var name string
	var (
		name = "Alexander"
		age  = 28
	)
	fmt.Println(name)
	fmt.Println(name, age)
	var c = fmt.Sprintf("My name %s. And i'm %d years old", name, age)
	fmt.Println(c)

}
