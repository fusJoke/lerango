package main

import "fmt"

type LinkNode struct {
	Data int
	Next *LinkNode
}

func main() {
	node := new(LinkNode)
	node.Data = 2

	node1 := new(LinkNode)
	node1.Data =3

	node2 := new(LinkNode)
	node2.Data =4

	node.Next  = node1
	node1.Next = node2

	nowNode := node

	for {
		if nowNode != nil {
			fmt.Println(nowNode.Data)
			nowNode = nowNode.Next
			continue
		}
		break
	}


}
