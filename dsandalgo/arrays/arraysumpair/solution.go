package main

func InListInt(inList []int32, k int32) bool {
	/*
		Loop over an integer array and return bool value depending on the presense of an integer.
	*/
	for _, num := range inList {
		if num == k {
			return true
		}
	}
	return false
}

func sortPair(k int32, l int32) [2]int32 {
	/*
		Sort two integers and return [min, max]
	*/
	if k > l {
		return [2]int32{l, k}
	}
	return [2]int32{k, l}
}

func ArraySumPair(inList []int32, k int32) [][2]int32 {
	/*
		Given an integer array, output all the unique pairs that sum up to a specific value k.
	*/

	// The inner map needs to be an array (fixed length) not slice (unknown length)
	// so it can be used as a key in a map.
	//https://blog.golang.org/maps#TOC_5.
	var pairList [][2]int32

	// Make an empty map
	pairMap := make(map[[2]int32]int32)

	// Logic: To look for pairs that make up a sum, subtract the current number from the total (k)
	for _, num := range inList {
		lookUp := k - num
		if ok := InListInt(inList, lookUp); ok { // Can be optimized to not search through the entire input list
			p := sortPair(lookUp, num)
			if _, ok := pairMap[p]; !ok { // The ok var is local to this scope (within the if)
				pairMap[p] = 1
			}
		}
	}
	for key, _ := range pairMap {
		pairList = append(pairList, key)
	}
	return pairList
}
