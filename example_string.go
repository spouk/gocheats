package main

import (
	"fmt"
	"github.com/spouk/gocheats/utils"

)

func main() {

	r := utils.NewRandomize()

	for x := 0; x < 10; x++ {
		fmt.Printf("RandomString: %v\n", r.RandomString(10))
	}
	for x := 0; x < 10; x++ {
		fmt.Printf("RandomString: %v\n", r.RandomStringSlice(10, 20))
	}
	for x := 0; x < 10; x++ {
		fmt.Printf("RandomSlice: %v\n", r.RandomStringChoice(10,utils.Lhexdigits))
	}
	for x := 0; x < 10; x++ {
		fmt.Printf("RandomSlice: %v\n", r.RandomStringChoice(10,utils.Lhexdigits))
	}
	fmt.Println()
	for x := 0; x < 10; x++ {
		fmt.Printf("RandomSlice: %v\n", r.RandomStringSliceChoice(10,10,utils.Ldigits))
	}
	fmt.Println()
	for x := 0; x < 10; x++ {
		fmt.Printf("RandomSlice: %v\n", r.RandomStringSliceChoice(10,10,utils.Lhexdigits))
	}
	fmt.Println()
	for x := 0; x < 10; x++ {
		fmt.Printf("RandomSlice: %v\n", r.RandomStringSliceChoice(10,10,utils.LasciiUppercase))
	}
}
