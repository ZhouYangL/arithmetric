package main

import (
	"testing"
	"reflect"
	"fmt"
)

type student struct {
	s string
	age int
}

func Test_reflect(c *testing.T)  {
	s := student{"zhou", 10}
	var v interface{}
	v = s
	k := reflect.ValueOf(v).MapKeys()
	fmt.Println(k)
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(reflect.ValueOf(v))
}
