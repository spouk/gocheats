package main

import (
	. "github.com/spouk/gocheats/datastructure/library"
	"time"
	"fmt"
)

func main() {
	e := NewExampleSlice()
	e.Log.Printf("Result direct: %v\n", e.RandomElement(10, 100))
	result := &[]int{}
	//var result *[]int
	e.RandomElementRecursive(10, 100, result)
	e.Log.Printf("Result recursive: %v\n", result)
	e.TimeTracker(time.Now(), "sdfgfg")

	e.Stock = e.RandomElement(10, 100)
	e.Log.Print(e.Stock)

	var s = e.RandomElement(10, 100)
	e.Log.Printf("Result %v\n", s)
	for x:=0; x < 10; x ++ {
		cutslice(s)
	}
}
func cutslice(v []int) {
	if len(v) == 0 {
		return
	}
	element := v[0]
	v = append(v[:0], v[0+1:]...)
	fmt.Printf("Element: [%v] [%v]\n", element, v)
}
