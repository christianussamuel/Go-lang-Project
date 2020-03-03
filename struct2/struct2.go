package main

import (
	"fmt"
	"strings"

)

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)

}

//s student dibawah adalah receiver type
func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

var (
	s1   = student{"john wick", 20}
	name = s1.getNameAt(2)
)

func main() {

	s1.sayHello()
	fmt.Println("nama panggilan :", name)

}
