package main

import (
	"ds/searching"
)

func main() {
	// rand.NewSource(2)
	// size := 10
	// min := 0
	// max := 100

	// arr := make([]int, size)
	// for i := 0; i < size; i++ {
	// 	arr[i] = rand.Intn(max-min+1) + min
	// }

	// arr := []int{2, 3, 1, 4, 5, 6}
	// fmt.Println(arr)
	// fmt.Println(searching.BinarySearch(arr, 0, len(arr)-1, 2))

	arvore := searching.BinaryTree{Value: 12, Left: nil, Right: nil}
	searching.Insert(&arvore, 12)
	searching.Insert(&arvore, 1)
	searching.Insert(&arvore, 21)
	searching.Insert(&arvore, 5)
	searching.Insert(&arvore, 15)
	searching.Insert(&arvore, 3)
	searching.InorderTraversal(&arvore)
}
