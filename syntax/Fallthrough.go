package main

import "fmt"

func main() {
	var countResult = true
	var value = 10

	switch value {
	case 10:
		fmt.Println("1、case 条件语句为 10")
		fallthrough
	case 2:
		fmt.Println("2、case 条件语句为 2")
		fallthrough
	case 4:
		fmt.Println("3、case 条件语句为 4")
		fallthrough
	case 5:
		fmt.Println("4、case 条件语句为 5")
	case 6:
		fmt.Println("5、case 条件语句为 6")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}

	switch {
	case false == countResult:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true == countResult:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false == countResult:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true == countResult:
		fmt.Println("4、case 条件语句为 true")
	case false == countResult:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
