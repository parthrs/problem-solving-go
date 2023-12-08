package arrays

/*
You consider two pots at a time:
- Start from beginning; This way the following logic is strict (i.e. you can only jump by 2 or 3 from here on)
- 00 means pot at left (i.e. 10) and move by two
- 10 would mean just jump by two (since potted plant already on left)
- 01 means you jump by three (since you jump by one first, make it 10 - can't be 11 since no adjacent pots - and then jump by two like before)
- If you've reached the last pot, and its 0, pot plant. The very first point is what makes this happen, you arrive at the last pot
  either by jumping 2 or 3, so it will be either 100 (jump by 2) or 0100 (jump by 3) - either way last three digits will always either be 010 or 100.
*/

func CanPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i += 2 { // Since we are also decrementing n, check for that in addition to i
		if flowerbed[i] == 0 { // If current pot is empty two condition can arise -
			if i == len(flowerbed)-1 || flowerbed[i+1] == 0 { // If the next one is also 0, in which case, pot (i.e. n--) and jump by 2 NOTE: Its important to check if i+1 is going out of bounds first!!!
				n--
			} else { // If the next one is 1, then jump by 3 in total (1 here and two at the start of the loop)
				i++
			}
		}
	}
	return n == 0
}
