package main

/**
Runtime: 12 ms, faster than 89.45% of Go online submissions for Median of Two Sorted Arrays.
Memory Usage: 5.9 MB, less than 10.00% of Go online submissions for Median of Two Sorted Arrays.
02/14/2020
MM/DD/YY
*/

/*
 * Main shellsort function
 * Suprisingly, shell sort has the fastest runtime for this method out of other sorting algorithms, such as
 * quicksort, merge sort, and merging (from merge sort)
 * Either that, or my other sort algorithms were badly programmed
 */
func shellSort(array []int) []int {
	gap := len(array) / 2

	for gap >= 1 {
		for i := gap; i < len(array); i++ { // Iterate through array, starting from gap
			comparator := array[i] // Make comparisons with this
			var output int         // For accessing x outside the array
			var index int          // For negative indexes

			for x := i; x > gap-2; x -= gap { // Iterate throguh array with gap as the step
				output = x     // To access x outside the loop
				if x-gap < 0 { // In case of negative index
					index = len(array) - x - gap
				} else {
					index = x - gap
				}

				if array[index] <= comparator { // Break when correct spot is found
					break
				} else { // Otherwise, move elements forward to make space
					array[x] = array[index]
				}
			}
			array[output] = comparator // Insert comparator in the correct spot
		}
		gap /= 2 // Increment the gap
	}

	return array
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	array := append(nums1, nums2...)
	array = shellSort(array)

	if len(array)%2 == 1 {
		return float64(array[len(array)/2])
	}

	return float64(array[len(array)/2]+array[len(array)/2-1]) / 2
}
