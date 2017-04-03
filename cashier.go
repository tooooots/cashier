package cashier

import (
	"fmt"
	"sort"
)

func Cashier(remaining int) (refund map[int]int) {

	// slice of available coins, sorted from large to small
	coins := []int{5, 2, 1, 10, 20, 50}
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	fmt.Printf("available coins: %v \n", coins)

	// construct a map of coin=quantity
	refund = make(map[int]int)

	for _, coin := range coins {

		for i := 1; remaining > 0; i++ {
			// substract by the largest coin unless < 0
			if remaining-coin >= 0 {
				remaining = remaining - coin
				// append to refund map with coin=quantity
				refund[coin] = i
				continue
			}
			break
		}
		// break the loop when we have the complete change
		if remaining == 0 {
			return
		}
	}
	return
}
