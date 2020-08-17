package main

// Solution for Runtime and decent memory

/*
Runtime: 4 ms, faster than 94.56% of Go online submissions for Two Sum.
Memory Usage: 3.1 MB, less than 71.03% of Go online submissions for Two Sum.
16/08/2020
DD/MM/YY
*/

func shellSort(array []int) []int {
	gap := len(array)/2
	for gap >= 1 {
		for i := gap; i < len(array); i++ { // Iterate through array, starting from gap
			comparator := array[i] // Make comparisons with this
			var output int // For accessing x outside the array
			var index int // For negative indexes

			for x := i; x > gap-2; x -= gap { // Iterate throguh array with gap as the step
				output = x // To access x outside the loop
				if x-gap < 0 { // In case of negative index
					index = len(array)-x-gap
				} else {
					index = x-gap
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

func indexOf(array []int, target int) int {
    for index, elem := range array {
        if elem == target {
            return index
        }
    }
    return -1
}

func twoSum(nums []int, target int) []int {
    arr := make([]int, len(nums))
    copy(arr, nums)
    array := shellSort(arr)
    
    for ind, elem := range array {
        for i := len(array)-1; i > ind; i-- {
            if elem + array[i] == target {
                if elem != array[i] {
                    return []int {indexOf(nums, array[i]), indexOf(nums, elem)}
                } else {
                    foo := indexOf(nums, elem)  
                    return []int {foo, indexOf(nums[foo+1:], (array[i]))+foo+1}
                }
            } else if target-elem > array[i] {
                break
            }
        }
	}
	
    return []int {};
}