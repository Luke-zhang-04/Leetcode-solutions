package Two_Sum

// Runtime: 76 ms, faster than 8.22% of Go online submissions for Two Sum.
// Memory Usage: 2.9 MB, less than 100.00% of Go online submissions for Two Sum.
// 16/12/2019
// DD/MM/YY

func in(target int, list []int) bool {
	for _, i := range list {
		if target == i {
			return true
		}
	}
	return false
}

func count(target int, list []int) int {
	total := 0
	for _, i := range list {
		if i == target {
			total++
		}
	}
	return total
}

func indice(target int, list []int) []int {
	indices := make([]int, 0, 0)
	for i := 0; i < len(list); i++ {
		if list[i] == target {
			indices = append(indices, i)
		}
	}
	return indices
}

func indexOf(target int, list []int) int {
	for k, v := range list {
		if target == v {
			return k
		}
	}
	return -1
}

func twoSum(nums []int, target int) []int {
	for _, i := range nums {
		if in(target-i, nums) {
			if count(target-i, nums) >= 2 {
				indices := indice(target-i, nums)
				return []int{indices[0], indices[1]}
			} else if indexOf(i, nums) == indexOf(target-i, nums) {
				continue
			} else {
				return []int{indexOf(i, nums), indexOf(target-i, nums)}
			}
		}
	}
	return nil
}
