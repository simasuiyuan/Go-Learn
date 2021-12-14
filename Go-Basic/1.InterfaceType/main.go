package main

import (
	"fmt"
	"reflect"
)

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

// create empty interface

type Interface interface { // declare a Interface type that is a interface
	Method()
}

type String struct { // declare a String type with struct
	Value string
}

func (s *String) Method() {} // implement Method

type Interger int

func (i Interger) Method() {}

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

	iface = Interger(100)
	fmt.Printf("Value: %v, Type: %T\n", iface, iface)
}
