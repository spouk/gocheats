package library

import "fmt"

//списки (односвязный)
type BaseListOne struct {
}

//список двусвязный
type BaseListTwo struct {
}

//базовый элемент списка
type NodeList struct {
	Count int   //общее количество
	Start *Node //указатель на следующий элемент списка, если нет то NULL
}
type Node struct {
	Value interface{}
	Next  *Node
}

func NewListArray() *NodeList {
	return &NodeList{
		Count: 1,
		Start: nil,
	}
}
func (n *NodeList) Add(nn *Node) {
	if n.Start == nil {
		n.Start = nn
	} else {
		next := n.Start
		for next.Next !=nil {
			next = next.Next
		}
		next.Next = nn
	}
}
func (n *NodeList) ShowList() {
	if n.Start == nil {
		fmt.Printf("List is empty\n")
	} else {
		curent := n.Start
		for curent.Next != nil {
			fmt.Printf("Element: %v\n", curent)
			curent = curent.Next
		}
		fmt.Printf("LastElement: %v\n", curent)

	}
}
