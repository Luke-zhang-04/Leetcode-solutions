// Runtime: 104 ms, faster than 92.06% of JavaScript online submissions for Median of Two Sorted Arrays.
// Memory Usage: 39.2 MB, less than 72.34% of JavaScript online submissions for Median of Two Sorted Arrays.
// 02/14/2020
// MM/DD/YY

/*
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
*/

function merge(array1, array2) {
    let array = []
    let first = 0
    let second = 0

    while (first < array1.length && second < array2.length) {
        if (array1[first] > array2[second]) {
            array.push(array2[second])
			second++
		} else if (array1[first] < array2[second]) {
            array.push(array1[first])
			first++
		} else {
            array.push(array2[second])
            array.push(array1[first])
			second++
			first++
        }
    }

    while (first < array1.length) {
        array.push(array1[first])
        first++
    }

    while (second < array2.length) {
        array.push(array2[second])
        second++
    }

    return array
}

var findMedianSortedArrays = function(nums1, nums2) {
    let array = merge(nums1, nums2)
    if (array.length % 2 === 1) {
        return array[Math.floor(array.length/2)]
    } else {
        return (array[Math.floor(array.length/2)] + array[Math.floor(array.length/2)-1]) / 2
    }
}