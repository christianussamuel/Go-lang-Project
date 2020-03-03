package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana"}
	var bFruits = fruits[0:2]
	fmt.Println(cap(bFruits))
	fmt.Println(len(bFruits))
	fmt.Println(fruits)
	fmt.Println(bFruits)

	var cFruits = append(bFruits, "papaya")
	fmt.Println(fruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits)

	var fruits1 = []string{"apple"}
	var aFruits = []string{"watermelon", "pinnaple"}
	var copiedElemen = copy(fruits1, aFruits)
	fmt.Println(fruits1)
	fmt.Println(aFruits)
	fmt.Println(copiedElemen)

}
