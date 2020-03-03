package main

import (
	"fmt"

)

var (
	point = 5
)

func main() {

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("Try harder")
			}
		}
	}

}
