#include <vector>

// Runtime: 20 ms, faster than 72.52% of C++ online submissions for Median of Two Sorted Arrays.
// Memory Usage: 10.6 MB, less than 21.65% of C++ online submissions for Median of Two Sorted Arrays.
// 02/23/2020
// MM/DD/YYYY
// Suprisingly, shell sort has the fastest runtime for this method out of other sorting algorithms, such as
// quicksort, merge sort, and merging (from merge sort)

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

std::vector<int> combine(std::vector<int> &array1, std::vector<int> &array2) {
    std::vector<int> array = array1;
    array.insert(array.end(), array2.begin(), array2.end());
    return array;
}

class Solution {
public:
    double findMedianSortedArrays(std::vector<int>& nums1, std::vector<int>& nums2) {
        std::vector<int> array = combine(nums1, nums2);
        std::vector<int> sorted = shell_sort(array);
        if (array.size() % 2 == 1) {
            return array[double(array.size())/2];
        } else {
            return double(array[float(array.size())/2] + array[float(array.size())/2-1]) / 2;
        }
    }
};