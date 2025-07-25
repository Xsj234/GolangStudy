package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("value:", reflect.ValueOf(arg))
}

func main() {
	var a int64 = 2222
	reflectNum(a)
}
