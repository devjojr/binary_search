package main

import "fmt"

func binarySearch(arr []int, data int) bool {
	first := 0
	last := len(arr) - 1
	found := false

	for first <= last && !found {
		mid := (first + last) / 2
		if arr[mid] == data {
			found = true
		} else {
			if data < arr[mid] {
				last = mid - 1
			} else {
				first = mid + 1
			}
		}
	}
	return found
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println(binarySearch(arr, 11))
	fmt.Println(binarySearch(arr, 25))
	fmt.Println(binarySearch(arr, 2))
}
