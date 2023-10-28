package searching

import "fmt"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func Insert(tree *BinaryTree, x int) {
	if x < tree.Value {
		if tree.Left == nil {
			tree.Left = &BinaryTree{x, nil, nil}
		} else {
			Insert(tree.Left, x)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BinaryTree{x, nil, nil}
		} else {
			Insert(tree.Right, x)
		}
	}
}

func InorderTraversal(tree *BinaryTree) {
	if tree.Left != nil {
		InorderTraversal(tree.Left)
	}
	fmt.Println(tree.Value)
	if tree.Right != nil {
		InorderTraversal(tree.Right)
	}
}
