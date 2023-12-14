package hashmap

import "fmt"

/*
Crude solution but for now its okay
*/

func EqualPairs(grid [][]int) (count int) {
	rowMap := map[string]int{}
	colMap := map[string]int{}

	for _, a := range grid {
		rowMap[fmt.Sprintf("%v", a)] += 1
	}
	for i := 0; i < len(grid); i++ { // col ptr
		sl := []int{}
		for j := 0; j < len(grid); j++ { // row ptr
			sl = append(sl, grid[j][i])
		}
		colMap[fmt.Sprintf("%v", sl)] += 1
	}
	for k, v := range rowMap {
		if val, found := colMap[k]; found {
			count += (v * val)
		}
	}
	return
}
