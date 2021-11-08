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

func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		// Sort of a reverse bubble sort
		// 	Also implied, only start the reverse loop if the last elem of the sublist
		//  has a greater element.
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			temp := a[j-1]
			a[j-1] = a[j]
			a[j] = temp
		}
	}
	return a
}

func main() {
	fmt.Println(InsertionSort([]int{8, 3, 2, 1, 5, 7, 7, 0}))
	fmt.Println(InsertionSort([]int{0, 1, 2, 3, 5, 7, 7, 8}))
	fmt.Println(InsertionSort([]int{7, 9, 100, 11, -5, -7, 0, 10}))
}
