package main

import "fmt"

func main() {

	//var a string
	////pair<statictype:string, value:"aceld">
	//a = "aceld"
	//
	////pair<type:string, value:"aceld">
	//var allType interface{}
	//allType = a
	//
	//str, _ := allType.(string)
	//fmt.Println(str)

	//var a string
	//a = "hello"
	var alltype interface{} //指针类型
	//alltype = a
	if s, ok := alltype.(string); ok {
		fmt.Printf("type is %T\n", s)
	}

}
