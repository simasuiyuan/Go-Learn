package main

import (
	"fmt"
	"reflect"
)

// 1. interface example
type Priority int8

func (p Priority) String() string {
	fmt.Println(reflect.TypeOf(p))
	switch p {
	case 0:
		return "low"
	case 1:
		return "high"
	}
	// switch int8(p) {
	// case 0:
	// 	return "low"
	// case 1:
	// 	return "high"
	// }
	return "unknown"
}

// 2. create empty interface

type Interface interface { // declare a Interface type that is a interface
	Method()
}

type String struct { // declare a String type with struct
	Value string
}

func (s *String) Method() {} // implement Method

// 3. interface: type assertion
type Interger int

func (i Interger) Method() {} // empty interface

func main() {
	var p Priority = 0

	fmt.Println(p.String())         // metods call
	fmt.Println(Priority.String(p)) // function call
	fmt.Println(p)
	/*
		 p is a interface type

		 type Stringer interface {
			 String() string
		 }

	*/

	var iface Interface

	iface = &String{"hello world"}
	fmt.Printf("Value: %v, Type: %T\n", iface, iface)

	// 3. interface: type assertion
	iface = Interger(100)
	fmt.Printf("Value: %v, Type: %T\n", iface, iface)

	t, ok := iface.(Interger) // type assertion
	fmt.Printf("OK? %t, Value %v, Type %T\n", ok, t, t)

	iface = &String{"hello"}
	t, ok = iface.(Interger) // type assertion
	fmt.Printf("OK? %t, Value %v, Type %T\n", ok, t, t)

	describe("hello")
	describe(Interger(100))
	describe(10)

}

//type assertion function example
func describe(i interface{}) {
	switch v := i.(type) {
	case Interger:
		fmt.Printf("int %d\n", v)
	case string:
		fmt.Printf("string %s\n", v)
	default:
		fmt.Printf("unknown %T - %v\n", i, i)
	}
}
