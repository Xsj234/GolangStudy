package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex" json:"xxx" bson:"111"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		tagjson := t.Field(i).Tag.Get("json")
		tagbson := t.Field(i).Tag.Get("bson")
		fmt.Println("info: ", taginfo, " doc: ", tagdoc, " json: ", tagjson, " bson: ", tagbson)
	}
}

func main() {
	var re resume

	findTag(&re)

}
