package main

import "fmt"

//1
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

//3
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	s := S{}
	f(&s) //4

	var i I //声明 i
	i = &s  //赋值 s 到 i
	fmt.Println(i.Get())
}
