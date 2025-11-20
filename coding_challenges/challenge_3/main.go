// Prompt: For n warehouses containing piles[i] packages, and a single truck with h hours before a deadline, find the minimum capacity k (packages/hour) that ships everything in time. The truck visits one warehouse per hour and, if a warehouse has fewer than k packages, ships all of them then idles for the rest of the hour.
// Output: Return the smallest integer k that satisfies the requirement.
// Examples:
// piles = [3, 6, 7, 11], h = 8 → k = 4
// piles = [30, 11, 23, 4, 20], h = 5 → k = 30
// Note: A second version of the prompt describes “items” instead of “packages”; the algorithm is identical.

// You can edit this code!
// Click here and start typing.
// Check leetcode 875 for optimized approach.
package main

import "fmt"

func main() {
	piles := []int{1, 1, 1, 1}
	h := 10
	try := 1
	if len(piles) > h {
		fmt.Println(0)
		return
	}
	for {
		count := 0
		for _, k := range piles {
			count += (k + try - 1) / try
		}
		if count <= h {
			fmt.Println(try)
			break
		}
		try++
	}
}
