package main

import "fmt"

func main() {
	/**注释*/
	fmt.Println("dddddd")

	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// 声明一个变量并初始化
	var a1 = "RUNOOB"
	fmt.Println(a1)

	// 没有初始化就为零值
	var b1 int
	fmt.Println(b1)

	// bool 零值为 false
	var c1 bool
	fmt.Println(c1)
}
