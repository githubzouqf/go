package main

import "fmt"

func GetYangHuiTriangleNextLine(inArr []int) []int {
	var out []int
	var i int
	arrLen := len(inArr)
	out = append(out, 1)
	if 0 == arrLen {
		return out
	}
	for i = 0; i < arrLen-1; i++ {
		out = append(out, inArr[i]+inArr[i+1])
	}
	out = append(out, 1)
	return out
}

func triangle(n int) {
	var item []int
	for i := 1; i <= n; i++ {
		item_len := len(item)
		if item_len == 0 {
			item = append(item, 1)
		} else {
			temp_s := []int{1}
			for j := 0; j < item_len-1; j++ {
				temp_s = append(temp_s, item[j]+item[j+1])
			}
			temp_s = append(temp_s, 1)
			item = temp_s
		}
		fmt.Println(item)
	}
}

func yanghuisanjiao(rows int) {
	for i := 0; i < rows; i++ {
		number := 1
		for k := 0; k < rows-i; k++ {
			fmt.Print("  ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%5d", number)
			number = number * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}

func main() {
	nums := []int{}
	var i int
	for i = 0; i < 10; i++ {
		nums = GetYangHuiTriangleNextLine(nums)
		fmt.Println(nums)
	}
	fmt.Println()
	fmt.Println()
	triangle(10)
	fmt.Println()
	fmt.Println()

	yanghuisanjiao(10)
}
