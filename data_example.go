package main

import (
	. "bitbucket.org/cyberspouk/learning/datastructure/library"

)

func main() {
	list := NewListArray()
	list.Add(&Node{Value:1})
	list.Add(&Node{Value:2})
	list.Add(&Node{Value:3})
	list.Add(&Node{Value:4})
	list.Add(&Node{Value:5})
	list.Add(&Node{Value:6})
	list.Add(&Node{Value:7})
	list.ShowList()
}