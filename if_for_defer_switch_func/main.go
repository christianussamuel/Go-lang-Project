package main

import (
	"fmt"
	"math"
	"runtime"

)

var (
	e = "\n"
)

func sqrt(x float64) string {

	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))

}
  
func pow(x, n, y float64) float64 {
	if v := math.Pow(x, n); v < 10 {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, y)
	}
	return y
}

func main() {
	sum := 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
	}
	fmt.Println("\n", sum)

	fmt.Println(e, sqrt(2), sqrt(-4), e)

	fmt.Println(pow(3, 2, 20))
	fmt.Println(pow(3, 3, 20))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)

	}
}
