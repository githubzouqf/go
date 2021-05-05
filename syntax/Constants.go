package main

import (
	"fmt"
	"unsafe"
)

/**常量枚举*/
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	x = "abc"
	y = len(x)
	z = unsafe.Sizeof(x)
)

func main() {
	/**显式常量定义*/
	const a1 string = "abc"
	/**隐式类型定义*/
	const b2 = "abc"

	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d \n", area)
	println()
	println(a, b, c)
	println()
	println(Unknown, Female, Male)
	println()
	println(x, y, z)
}
