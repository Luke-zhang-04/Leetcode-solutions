/**
 * Runtime: 164 ms, faster than 18.73% of C++ online submissions for Two Sum.
 * Memory Usage: 9.3 MB, less than 81.47% of C++ online submissions for Two Sum.
 * 02/21/2020
 * MM/DD/YYYY
 */

#include <vector>

bool contains(int target, std::vector<int>& list) {
    for (const int &i : list) {
        if (target == i) {
            return true;
        }
    }

    return false;
}

int count(int target, std::vector<int>& list) {
    int total = 0;

    for (const int &i : list) {
        if (i == target) {
			total++;
		}
	}
    
	return total;
}

int indexOf(int target, std::vector<int>& list) {
    for (unsigned int i = 0; i < list.size(); i++) {
        if (target == list[i]) {
            return i;
        }
    }

    return -1;
}

std::vector<int> indice(int target, std::vector<int>& list) {
    std::vector<int> indices;
    
    for (unsigned int i = 0; i < list.size(); i++) {
        if (list[i] == target) {
            indices.push_back(i);
        }
    }

    return indices;
}

class Solution {
    public:
    std::vector<int> twoSum(std::vector<int>& nums, int target) {
        for (const int &i : nums) {
            if (contains(target-i, nums)) {
                if (count(target-i, nums) >= 2) {
                    std::vector<int> indices = indice(target-i, nums);
                    std::vector<int> res = {indices[0], indices[1]};

                    return res;
                } else if (indexOf(i, nums) == indexOf(target-i, nums)) {
                    continue;
                } else {
                    std::vector<int> res = {indexOf(i, nums), indexOf(target-i, nums)};

                    return res;
                }
            }
        }
        std::vector<int> res = {};

        return res;
    }
};
