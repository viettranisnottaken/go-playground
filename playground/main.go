package main

import (
	"fmt"
	"reflect"
)

type Comic struct {
	Universe string
}

func main() {
	c := Comic{"123"}
	cPtr := &c

	fmt.Println(cPtr.Universe)

	cPtr.Universe = "345"

	fmt.Println(cPtr.Universe)

	reflection := reflect.ValueOf(cPtr).Elem()

	fmt.Println(reflection.FieldByName("Universe"))

	reflection.FieldByName("Universe").SetString("Uni")

	fmt.Println(reflection.FieldByName("Universe"))
}
