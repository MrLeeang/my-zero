package main

import (
	"fmt"
	"reflect"
)

type XXX struct {
	a string
	b int64
}

func (x *XXX) H(d string) {
	fmt.Println(d)
}

func main() {

	aa := &XXX{
		a: "a",
		b: 0,
	}

	// 获取类型和反射值
	value := reflect.ValueOf(aa)
	typeInfo := reflect.TypeOf(aa)

	// fmt.Println("Value:", value)                 // 输出: Value: 3.14
	// fmt.Println("Type:", typeInfo)               // 输出: Type: float64
	// fmt.Println("Kind:", typeInfo.Kind())        // 输出: Kind: float64
	// fmt.Println("Value as float64:", value.Float()) // 输出: Value as float64: 3.14
	fmt.Println(value)

	me0 := value.MethodByName("H")
	me, _ := typeInfo.MethodByName("H")

	me0.Call([]reflect.Value{reflect.ValueOf("jj")})

	me.Func.Call([]reflect.Value{reflect.ValueOf("jj")})

}
