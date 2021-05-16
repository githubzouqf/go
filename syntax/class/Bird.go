package main

/**定义类*/

type Bird struct {
}

func (b *Bird) Fly() {

}

type IFly interface {
	Fly()
}

func main() {
	var fly IFly = new(Bird)
	fly.Fly()
}
