package main

import "fmt"

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	vname1 int    = 0
	vname2 string = "zqf"
)

func main() {
	/**go可以自动推断声明的数据类型*/
	var str1 string = "str1"
	var str2 = "str2"
	/**默认值*/
	var age int
	/**声明变量，并设置默认值*/
	var declareVariableA, declareVariableB int = 1, 2
	/*使用:=声明，变量必须是未定义过得，不仅是类型或者是var*/
	noDeclareVariable := 0

	fmt.Println("打印不同变量的定义")
	fmt.Println(noDeclareVariable, declareVariableA, declareVariableB)
	fmt.Println(age)
	fmt.Println(str1, str2)

	/**格式化输出， %d 表示整型数字，%s 表示字符串*/
	var stockCode = 123
	var endDate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var targetUrl = fmt.Sprintf(url, stockCode, endDate)
	fmt.Println(targetUrl)

	var a0, b0 = 0, 0
	b0 = 1
	fmt.Println(a0, b0)

	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)

	//var a *int
	//var a []int
	//var a map[string] int
	//var a chan int
	//var a func(string) int
	//var a error // error 是接口

	var i1 int
	var f1 float64
	var b1 bool
	var s1 string
	fmt.Printf("%v %v %v %q\n", i1, f1, b1, s1)

	var aPoint *int
	var aArr []int
	var aMap map[string]int
	var aChan chan int
	var aFun func(string) int
	var aError error // error 是接口

	fmt.Printf("%v %v %v %v %v %v \n", aPoint, aArr, aMap, aChan, aFun, aError)

	println(vname1, vname2)

}
