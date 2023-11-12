package main

import (
	"ds/trees"
	"fmt"
)

func main() {

	root := &trees.BinaryTree{Value: 5}

	root.Insert(3)
	root.Insert(7)
	root.Insert(1)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	root.PreorderTraversal()
	fmt.Println()
	root.InvertBinaryTree()
	root.PreorderTraversal()

}
