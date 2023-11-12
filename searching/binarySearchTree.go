package searching

import "ds/trees"

func Find(tree *trees.BinaryTree, x int) (ret bool) {
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
