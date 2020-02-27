#include <vector>

// Runtime: 8 ms, faster than 93.28% of C++ online submissions for Two Sum.
// Memory Usage: 9.4 MB, less than 67.57% of C++ online submissions for Two Sum.
// 02/25/2020
// MM/DD/YYYY

int indexOf(std::vector<int>& list, int target) {
    for (unsigned int i = 0; i < list.size(); i++) {
        if (target == list[i]) {
            return i;
        }
    }
    return -1;
}

std::vector<int> shell_sort(std::vector<int> &array) {
    int gap = array.size()/2;
    while (gap >= 1) {
        for (int i = gap; i < array.size(); i++) { //iterate through array, starting from gap
            int comparator = array[i]; //make comparisons with this
            int output; //for accessing x outside the array
            int index; //for negative indexes

            for (int x = i; x > gap-2; x -= gap) { //iterate throguh array with gap as the step
                output = x; //to access x outside the loop
                if (x-gap < 0) { //in case of negative index
                    index = array.size()-x-gap;
                } else {
                    index = x-gap;
                }

                if (array[index] <= comparator) { //break when correct spot is found
					break;
				} else { //otherwise, move elements forward to make space
					array[x] = array[index];
				}
            }
            array[output] = comparator; //insert comparator in the correct spot
        }
        gap /= 2; //increment the gap
    }
    return array;
}


class Solution {
    public:
        std::vector<int> twoSum(std::vector<int>& nums, int target) {
            std::vector<int> input_array(nums);
            std::vector<int> array = shell_sort(input_array);
            
            for (unsigned int ind = 0; ind < array.size(); ind ++) {
                for (unsigned int i = array.size()-1; i > ind; i--) {
                    if (array[ind] + array[i] == target) {
                        if (array[ind] != array[i]) {
                            std::vector<int> output = {
                                indexOf(nums, array[i]), 
                                indexOf(nums, array[ind])
                            };
                            return output;
                        } else {
                            int foo = indexOf(nums, array[ind]);
                            std::vector<int> slice = std::vector<int>(nums.begin() + foo + 1, nums.end());
                            std::vector<int> output = {
                                foo,
                                indexOf(slice, (array[i]))+foo+1
                            };
                            return output;
                        }
                    } else if (target-array[ind] > array[i]) {
                        break;
                    }
                }
            }
            std::vector<int> out = {};
            return out;
        }
};