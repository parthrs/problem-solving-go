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
	// Best case O(n) - since the inner loop will not run if the array is sorted
	for i := 1; i < len(a); i++ { // O(n)
		// Sort of a reverse bubble sort
		// 	Also implied, only start the reverse loop if the last elem of the sublist
		//  has a greater element.
		for j := i; j > 0 && a[j-1] > a[j]; j-- { // Conditional O(n)
			temp := a[j-1]
			a[j-1] = a[j]
			a[j] = temp
		}
	}
	return a // Worst case O(n^2)
}

func MergeSort(a []int) []int {
	/*
		Worst case: O(nlogn) -
			log n levels in all times
			amount of work on each level = nlogn (log to the base 2)
	*/
	if len(a) <= 1 {
		return a
	}

	var left []int
	var right []int

	left = append(left, a[:len(a)/2]...)
	right = append(right, a[len(a)/2:]...)

	left = MergeSort(left)
	right = MergeSort(right)

	// Merge
	var result []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i += 1
		} else {
			result = append(result, right[j])
			j += 1
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i += 1
	}

	for j < len(right) {
		result = append(result, right[j])
		j += 1
	}

	return result
}

func main() {
	fmt.Println(MergeSort([]int{8, 3, 2, 1, 5, 7, 7, 0}))
	fmt.Println(MergeSort([]int{0, 1, 2, 3, 5, 7, 7, 8}))
	fmt.Println(MergeSort([]int{7, 9, 100, 11, -5, -7, 0, 10}))
}
