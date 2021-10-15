package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 6.4
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))


	//根据反射的值，来获取对应的类型和数值
	v:= reflect.ValueOf(a)
	fmt.Println("kind is float64: ",v.Kind() == reflect.Float64)
	fmt.Println("type: ",v.Type())
	fmt.Println("value: ",v.Float())


}
