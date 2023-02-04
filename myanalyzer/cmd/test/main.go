package main

import "fmt"

type N struct {
	number int
}

func (n *N) PrintN() {
	fmt.Println(n.number)
}

type X N

func (x *X) PrintX() {
	fmt.Println(x.number)
}

type Y X

func (y *Y) PrintY() {
	fmt.Println(y.number)
}

func main() {
	n := N{number: 10}
	x := X{number: 100}
	y := Y{number: 1000}

	n.PrintN()
	x.PrintX()
	y.PrintY()

	// x.PrintN() はできない
	// -> Underlying type
}
