package main

import "fmt"

type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

//fakeString实现了String()方法，因此也实现了Stringer接口
func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

func main() {
	s := &fakeString{"Ceci n'est pas un string"}
	printString(s)
	printString("Hello, Gophers")

}
