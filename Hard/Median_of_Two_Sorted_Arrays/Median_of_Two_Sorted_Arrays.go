package Median_of_Two_Sorted_arrays

// Runtime: 12 ms, faster than 89.45% of Go online submissions for Median of Two Sorted Arrays.
// Memory Usage: 5.9 MB, less than 10.00% of Go online submissions for Median of Two Sorted Arrays.
// 02/14/2020
// MM/DD/YY
// Suprisingly, shell sort has the fastest runtime for this method out of other sorting algorithms, such as
// quicksort, merge sort, and merging (from merge sort)

func shell_sort(array []int) []int {
	gap := len(array) / 2
	for gap >= 1 {
		for i := gap; i < len(array); i++ { //iterate through array, starting from gap
			comparator := array[i] //make comparisons with this
			var output int         //for accessing x outside the array
			var index int          //for negative indexes

			for x := i; x > gap-2; x -= gap { //iterate throguh array with gap as the step
				output = x     //to access x outside the loop
				if x-gap < 0 { //in case of negative index
					index = len(array) - x - gap
				} else {
					index = x - gap
				}

				if array[index] <= comparator { //break when correct spot is found
					break
				} else { //otherwise, move elements forward to make space
					array[x] = array[index]
				}
			}
			array[output] = comparator //insert comparator in the correct spot
		}
		gap /= 2 //increment the gap
	}
	return array
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	array := append(nums1, nums2...)
	array = shell_sort(array)
	if len(array)%2 == 1 {
		return float64(array[len(array)/2])
	} else {
		return float64(array[len(array)/2]+array[len(array)/2-1]) / 2
	}
}
