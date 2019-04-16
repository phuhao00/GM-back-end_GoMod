package main

import (
	"fmt"
	"reflect"
)

type IT struct {
	i interface{}
}
func main()  {
	t:=IT{}
	t.i=int64(6)
	if t.i==nil {
		fmt.Println(t.i.(int64))
	}else {
		t.i=t.i.(int64)+3
		fmt.Println(t.i.(int64))
		fmt.Println(reflect.TypeOf(t.i))
	}

}
