package sorting

func Selectionsort(arr []int) []int {
	ret := []int{}
	for len(arr) > 0 {
		smallestIndex := findSmallest(arr)
		ret = append(ret, arr[smallestIndex])
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return ret
}

func findSmallest(arr []int) int {
	var smallest int = arr[0]
	var index int = 0
	for i, value := range arr {
		if arr[i] < smallest {
			smallest = value
			index = i
		}
	}
	return index
}
