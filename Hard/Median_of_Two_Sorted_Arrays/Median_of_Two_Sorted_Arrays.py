#Runtime: 92 ms, faster than 84.09% of Python3 online submissions for Median of Two Sorted Arrays.
#Memory Usage: 12.7 MB, less than 100.00% of Python3 online submissions for Median of Two Sorted Arrays.
#02/14/2020
#MM/DD/YY

class Solution:
    #from merge sort
    @staticmethod
    def merge(array1, array2): #merge two arrays
        array = [] #merged array
        first = second = 0 #starting points

        while first < len(array1) and second < len(array2):
            if array1[first] > array2[second]:
                array.append(array2[second])
                second += 1
            elif array1[first] < array2[second]:
                array.append(array1[first])
                first += 1
            else:
                array.append(array1[first])
                array.append(array2[second])
                first += 1
                second += 1
        
        while first < len(array1):
            array.append(array1[first])
            first += 1

        while second < len(array2):
            array.append(array2[second])
            second += 1

        return array
    
    def findMedianSortedArrays(self, nums1: list, nums2: list) -> float:
        array = self.merge(nums1, nums2)
        if len(array) % 2 == 1:
            return float(array[len(array)//2])
        else:
            return (array[len(array)//2] + array[len(array)//2-1]) / 2