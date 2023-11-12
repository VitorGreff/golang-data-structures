package trees

import "fmt"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) Insert(x int) {
	if x < tree.Value {
		if tree.Left == nil {
			tree.Left = &BinaryTree{x, nil, nil}
		} else {
			tree.Left.Insert(x)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BinaryTree{x, nil, nil}
		} else {
			tree.Right.Insert(x)
		}
	}
}

func (tree *BinaryTree) InorderTraversal() {
	if tree.Left != nil {
		tree.Left.InorderTraversal()
	}
	fmt.Println(tree.Value)
	if tree.Right != nil {
		tree.Right.InorderTraversal()
	}
}

func (tree *BinaryTree) PreorderTraversal() {
	fmt.Println(tree.Value)
	if tree.Left != nil {
		tree.Left.PreorderTraversal()
	}
	if tree.Right != nil {
		tree.Right.PreorderTraversal()
	}
}

func (tree *BinaryTree) PostOrderTraversal() {
	if tree.Left != nil {
		tree.Left.PostOrderTraversal()
	}
	if tree.Right != nil {
		tree.Right.PostOrderTraversal()
	}
	fmt.Println(tree.Value)

}

func (tree *BinaryTree) InvertBinaryTree() {
	if tree == nil {
		return
	}
	tree.Left, tree.Right = tree.Right, tree.Left
	tree.Left.InvertBinaryTree()
	tree.Right.InvertBinaryTree()
}
