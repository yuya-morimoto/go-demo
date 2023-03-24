package sample

import "fmt"

const secret = "abc"

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

var (
	i int
	s string
	b bool
)

func Type() {
	//var i int
	//var i = 2
	i := 2
	ui := uint(i)
	fmt.Println(i)
	fmt.Printf("i: %v %T\n", i, i)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	f := 1.2345
	s := "hello"
	b := true
	fmt.Printf("i: %v %T\n", f, f)
	fmt.Printf("i: %v %T\n", s, s)
	fmt.Printf("i: %v %T\n", b, b)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v\n", pi, title)

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Println(z)

	fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux)

}
