package main

import "fmt"

func BubbleSort(a []int) []int {
	/*
		Best case (1 or less swaps): O(n)
				Only loop once over the array
	*/
	swap := true
	for swap { // O(n) in the worst case
		swap = false
		for i := 1; i < len(a); i++ { // O(n)
			if a[i] < a[i-1] {
				temp := a[i-1]
				a[i-1] = a[i]
				a[i] = temp
				swap = true
			}
		}
	}
	return a // Worst case O(n^2)
}

func SelectionSort(a []int) []int {
	/*
		Best case O(n^2)
	*/
	start := 0

	for start < len(a)-1 { // O(n)
		lowest := start
		swap := false
		for i := start + 1; i < len(a); i++ { // O(n)
			if a[i] < a[lowest] {
				lowest = i
				swap = true
			}
		}
		if swap { // Swap only once, with the lowest elem
			temp := a[start]
			a[start] = a[lowest]
			a[lowest] = temp
		}
		start += 1
	}
	return a // Worst case O(n^2)
}

func inplaceInsert(arr []int, ind int, elem int) {
	if ind >= len(arr) {
		return
	}

	prev := arr[ind]

	for i := ind + 1; i < len(arr); i++ { // O(n)
		temp := arr[i]
		arr[i] = prev
		prev = temp
	}

	arr[ind] = elem
}

func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ { // O(n)
		j := i - 1
		for ; j >= 0; j-- { // O(n)
			if a[j] < a[i] {
				inplaceInsert(a[:i+1], j+1, a[i]) // Go slices are passed by value, but the value is a header (which includes a ptr)
				break
			}
		}
		if j < 0 && a[i] < a[0] {
			inplaceInsert(a[:i+1], 0, a[i])
		}
	}
	return a
}

func main() {
	fmt.Println(SelectionSort([]int{8, 3, 2, 1, 5, 7, 7, 0}))
	fmt.Println(SelectionSort([]int{0, 1, 2, 3, 5, 7, 7, 8}))
	fmt.Println(SelectionSort([]int{7, 9, 100, 11, -5, -7, 0, 10}))
}
