package main

import "fmt"

var (
	//explicit variable declaration, initialized with value
	firstname string = "Samuel"
	//explicit variable declaration, deffered assigned value
	lastname             string
	address                     = "Jl. Tubagus Ismail No.25"
	age                         = 20
	height                      = 170.31
	hobby, favgame       string = "playing games", "Dragon Nest"
	igaccount, fbaccount string

	//unused variable biasa nya untuk return fungsi yang tdak
	//perlu digunakan (dari 3 value hanya perlu return 2, gunakan _)
	_ = "will gone"

	//uint=unsigned int tidak ada negativ, 0 masuk

	vuint uint
	// vuint8  uint8
	// vuint16 uint16
	// vuint32 uint32
	// vuint64 uint64

	vint int
	// vint8  int8
	// vint16 int16
	// vint32 int32
	// vint64 int64

	bolean bool
	//lanjutkan sampai hal 87
)

func main() {

	lastname = "christianus"
	igaccount = "@xxx"
	fbaccount = "@xxxx"
	vuint = 255
	vint = 128

	name := new(string)

	fmt.Println(firstname, lastname)
	fmt.Println("Address 	:", address)
	fmt.Println("Age		:", age)
	fmt.Println("Height		:", height)
	fmt.Println("Hobby		:", hobby, "+", favgame)
	fmt.Println(name)

	fmt.Println(vuint, vint)
	fmt.Println(bolean)

}
