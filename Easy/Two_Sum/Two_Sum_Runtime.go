// Solution for Runtime and decent memory

// Runtime: 4 ms, faster than 95.13% of Go online submissions for Two Sum.
// Memory Usage: 3.1 MB, less than 59.62% of Go online submissions for Two Sum.
// 02/24/2020
// MM/DD/YY

func shell_sort(array []int) []int {
	gap := len(array)/2
	for gap >= 1 {
		for i := gap; i < len(array); i++ { //iterate through array, starting from gap
			comparator := array[i] //make comparisons with this
			var output int //for accessing x outside the array
			var index int //for negative indexes

			for x := i; x > gap-2; x -= gap { //iterate throguh array with gap as the step
				output = x //to access x outside the loop
				if x-gap < 0 { //in case of negative index
					index = len(array)-x-gap
				} else {
					index = x-gap
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
    array := shell_sort(arr)
    
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