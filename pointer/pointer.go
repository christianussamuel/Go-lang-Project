package main

import "fmt"

var (
// numA int  = 4
// numB *int = &numA
// numC int  = 44

)

func main() {

	//numB = &numC
	//value saja yg bisa ubah
	//address tidak bisa
	// numA = 19

	//&numA adalah untuk mengambil address reference memorinya (value disimpan)
	// fmt.Println("number A(value) ", numA)
	// fmt.Println("number A(address) ", &numA)
	//*numB untuk mengambl value yang ada di alamat memorinya (dereference)
	// fmt.Println("number B(value) ", *numB)
	// fmt.Println("number B(address) ", numB)

	// fmt.Println("number C(value) ", numC)
	// fmt.Println("number C(address) ", &numC)

	// fmt.Println("number B(value) ", *numB)
	// fmt.Println("number B(address) ", numB)

	var numPointer = 4
	fmt.Println("before :", numPointer)

	change(&numPointer, 10)
	fmt.Println("after :", numPointer)

	change2(numPointer, 101)
	fmt.Println("after :", numPointer)

}

func change(original *int, value int) {
	*original = value
}

//copy by value
func change2(original int, value int) {
	original = value
}
