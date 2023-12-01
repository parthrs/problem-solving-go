package arrays

/*
You have the daily prices of a stock ticker given to you, buy a single share and then sell it when profit is maximum.

- You can only buy the share once
- The ith indice represents the day of the price
- You can only sell the stock after you have bought it

[1,2,3]
[2,1,3]
[3,2,1]
*/

/*
This is not really a two pointer problem per say.
Logic: The idea is to compare pairs of elems through the array,
       find the min amongst them, and hold on to that min. Thats
			 the price we buy at, until we find a new min.
			 Compare the profit (current price - min) with each new elem,
			 and keep track of the max of profit.
			 So given two elems in a pair, if min price is today, you are
			 better off buying today and selling today.
*/

func MaximizeProfit(prices []int) (maxProfit int) {
	if len(prices) < 1 {
		return 0
	}
	minPrice := prices[0]
	for price := 1; price < len(prices); price++ {
		minPrice = min(minPrice, prices[price])
		maxProfit = max(maxProfit, prices[price]-minPrice)
	}
	return
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
