package main

import "fmt"

// An inversion in the array is when the array index of two locations in the array are opposite of the values.
// [0, 2, 1]  --> Inversion is (2,1)
// [2, 0, 1]  --> Inversion is again (2,1); i.e. they may need be subsequent

func DiscoverInversions(a []int) ([]int, int) {

	// Base case
	if len(a) == 1 {
		return a, 0
	}

	lh := a[:len(a)/2]
	rh := a[len(a)/2:]

	lh, x := DiscoverInversions(lh)
	rh, y := DiscoverInversions(rh)

	// Magic begins
	output := make([]int, len(a), len(a))
	i := 0
	j := 0
	z := 0 // Count of inversions

	for k := range a {
		if i < len(lh) && j < len(rh) {
			if lh[i] < rh[j] {
				output[k] = lh[i]
				i++
			} else {
				output[k] = rh[j]
				z += (len(lh) - i)
				j++
			}
		} else if j < len(rh) {
			output[k] = rh[j]
			j++
		} else if i < len(lh) {
			output[k] = lh[i]
			i++
		}
	}
	return output, (x + y + z)
}

func main() {
	fmt.Println(DiscoverInversions([]int{6, 5, 4, 3, 2, 1}))
}
