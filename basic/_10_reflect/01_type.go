package _10_reflect

import (
	"fmt"
	"playground/basic/_02_struct/vehicle"
	"reflect"
)

/**
 * Kind() 方法返回 Kind 类型, 详情参照 reflect.Kind [ uint 的一种别名类型 ]
 * type.Kind() 的返回值是 Kind 类型, 那么实际参数是 Kind.String() 的返回值
 * fmt.Println(param), 实际参数是 param.String() 的返回值
 */

// 基础数据类型的反射
func basicType() {
	typeI := reflect.TypeOf(1)
	typeS := reflect.TypeOf("hello")
	fmt.Println(typeI.Kind()) // int
	fmt.Println(typeS.Kind()) // string
}

// 结构体 & 指针类型的反射
func structType() {
	car1 := vehicle.Car{}
	car2 := new(vehicle.Car)
	car1Type := reflect.TypeOf(car1)
	car2Type := reflect.TypeOf(car2)
	fmt.Println(car1Type.Kind()) // struct
	fmt.Println(car2Type.Kind()) // ptr
}

// 指针解析
func ptrType() {
	car := &vehicle.Car{}
	// 结构体没有 String() 方法的情况下, 就是所有包外可见的属性键值对字符串: &{foo: bar, ...}
	// 相反, 输出 String() 的返回值
	fmt.Println(car)
	carType := reflect.TypeOf(car)
	fmt.Println(carType.Kind()) // ptr
	carElem := carType.Elem()
	fmt.Println(carElem.Kind()) // 解析指针指向的原始数据类型, struct
}

// 结构体: 获取属性
func getField() {
	//
}

// 结构体: 获取方法
func getMethod() {
	car1 := reflect.TypeOf(&vehicle.Car{})
	// todo:: 存疑, NumMethod 只会对 interface 类型的私有方法计数, 其他类型的私有方法不会计数, 如何证明
	// Ps: NumMethod() 不会计数结构体的私有方法, 但是 interface 的私有方法会计数
	for i, methodCnt := 0, car1.NumMethod(); i < methodCnt; i++ {
		method := car1.Method(i)
		fmt.Printf("name: %s\ttype: %s\texported: %t\n", method.Name, method.Type, method.IsExported())
	}

	fmt.Println("=================")
}

// 函数
func funcType() {
	//
}
