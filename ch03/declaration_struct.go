package main

import "fmt"

type Point struct {
	X, Y int
}

var (
	p = Point{1, 2}  // 初始化一个Point对象
	q = &Point{1, 2} // 初始化一个Point指针对象
	r = Point{X: 1}  // Y值默认为0
	s = Point{}      // X和Y的值都默认为0
)

func main() {
	fmt.Println(p, q, r, s)
}
