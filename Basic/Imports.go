package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"math/cmplx"
	"math/rand"

)

func tambah(x, y float32) float32 {
	return (x * y) + (x / y)
}

func tuker(x, y string) (string, string) {
	return y, x
}

func split(a float32) (x, y float32) {
	x = a * 5 / 8
	y = a - x
	return
}

var (
	c, python, java bool
	Bole            bool       = false
	maxint          uint64     = 1<<64 - 1
	z               complex128 = cmplx.Sqrt(-5 + 12i)

	x1, y1 int = 3, 4
	f1         = math.Sqrt(float64(x1*x1 + y1*y1))
	z1         = (f1)

	e = "\n"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x
}

func needFloat(x float32) float32 {
	return x
}

func main() {
	fmt.Println("Now you have problems. \n",
		math.Sqrt(7), "\n")
	fmt.Println("random number \n",
		rand.Intn(100), "\n")
	/*Seed random int*/
	h := md5.New()
	io.WriteString(h, "And Leon's getting larger ! \n")
	var seed uint64 = binary.BigEndian.Uint64(h.Sum(nil))
	fmt.Println(seed)
	rand.Seed(int64(seed))
	fmt.Println(rand.Int(), "\n")

	fmt.Println(math.Pi, "\n")

	fmt.Println(tambah(4, 6), "\n")

	a, b := tuker("ini kiri", "ini kanan")
	fmt.Println(a, b, "\n")

	fmt.Println(split(3))

	var i int
	i = 4
	fmt.Println("\n", i, c, python, java, "\n")

	var i1, j1 int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i1, j1, c, python, java, "\n")

	k := 7
	go1 := "lang"
	fmt.Println(i, k, go1, "\n")

	fmt.Printf("Type :	 Value :	%v\n", Bole)
	fmt.Printf("Type :	 Value :	%v\n", maxint)
	fmt.Printf("Type :	 Value :	%v\n", z, z)

	fmt.Println("\n", x1, y1, z1, e)

	i2 := 99
	f2 := 3.149
	g3 := 0.34 + 0.35i
	fmt.Println("i2 is of type", i2, e, "f2 is type", f2, e, "g3 is type", g3, e)

	const Truth = true
	fmt.Println("rules", Truth)

	//numeric constants
	fmt.Println(e, needInt(Small))
	fmt.Println(e, needFloat(Big))

}
