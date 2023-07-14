package _10_reflect

import (
	"fmt"
	"playground/basic/_02_struct/vehicle"
	"reflect"
)

/**
 * 1. value.Kind() == type.Kind()  返回值都是 reflect.Kind
 * 2. value.Addr() & value.Element() 是一对互逆操作
 */

// 基础数据值的反射
func basicValue() {
	valueI := reflect.ValueOf(1)
	valueS := reflect.ValueOf("hello")
	// Kind.String() => reflect.kindNames[Kind]
	fmt.Println(valueI.Kind()) // int
	fmt.Println(valueS.Kind()) // string
}

// 结构体 & 指针值的反射
func structValue() {
	car1 := vehicle.Car{}
	car2 := new(vehicle.Car)
	car1Value := reflect.ValueOf(car1)
	car2Value := reflect.ValueOf(car2)
	// Kind.String() => reflect.kindNames[Kind]
	fmt.Println(car1Value.Kind()) // struct
	fmt.Println(car2Value.Kind()) // ptr
	// 结构体值 & 指针值 相互转换
	car2PtrValue := car2Value.Elem() // ptr -> struct
	fmt.Println(car2PtrValue.Kind())
	car2StructValue := car2PtrValue.Addr() // struct -> ptr
	fmt.Println(car2StructValue.Kind())
	// ps. struct -> ptr 不能用 car1Value.Addr() 来进行转换
	// 会报错: reflect.Value.Addr of unaddressable value [recovered]
}

// 是否可寻址
func addressable() {
	//
}
