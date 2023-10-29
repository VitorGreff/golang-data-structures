package searching

import "fmt"

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func Find(tree *BinaryTree, x int) (ret bool) {
	if tree == nil {
		return false
	} else if x < tree.Value {
		if tree.Left == nil {
			return false
		} else {
			return Find(tree.Left, x)
		}
	} else if x > tree.Value {
		if tree.Right == nil {
			return false
		} else {
			return Find(tree.Right, x)
		}
	} else {
		return true
	}
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

func PreorderTraversal(tree *BinaryTree) {
	fmt.Println(tree.Value)
	if tree.Left != nil {
		PreorderTraversal(tree.Left)
	}
	if tree.Right != nil {
		PreorderTraversal(tree.Right)
	}
}

func PostOrderTraversal(tree *BinaryTree) {
	if tree.Left != nil {
		PostOrderTraversal(tree.Left)
	}
	if tree.Right != nil {
		PostOrderTraversal(tree.Right)
	}
	fmt.Println(tree.Value)

}
