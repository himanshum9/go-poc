//Challenge 6 · Rotated Sorted Array Search
//Goal: Given a rotated sorted array (e.g. [5,6,7,1,2,3,4]) and a target such as 6, return the index of the target in O(log n) time.

package main

import "fmt"

func search(sl []int, t int) {
	l := 0
	r := len(sl) - 1
	for l <= r {
		mid := l + (r-l)/2
		if sl[mid] == t {
			fmt.Println(mid)
			return
		}
		if sl[l] <= sl[mid] {
			if t >= sl[l] && t < sl[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {

			if t > sl[mid] && t <= sl[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}

		}

	}
	fmt.Println("-1")
}

func main() {
	sl := []int{5, 6, 7, 1, 2, 3, 4}
	target := 3
	search(sl, target)
}
