#Runtime: 120 ms, faster than 31.32% of Python3 online submissions for Two Sum.
#Memory Usage: 13.6 MB, less than 96.98% of Python3 online submissions for Two Sum.
#02/24/2020
#DD/MM/YY

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
            array.append(array2[second])
            second += 1
            array.append(array1[first])
            first += 1
    
    while first < len(array1):
        array.append(array1[first])
        first += 1

    while second < len(array2):
        array.append(array2[second])
        second += 1

    return array

def merge_sort(array):
    if len(array) > 1:
        #split array into to halves
        half = len(array) // 2 #halfway point
        left = array[:half]
        right = array[half:]

        #recursively call each merge sort for each half until array size is 1
        left = merge_sort(left)
        right = merge_sort(right)
        
        #merge left and right sides
        array = merge(left, right)

    return array

class Solution:
    def twoSum(self, nums: list, target: int) -> list:
        array = merge_sort(nums.copy())
        
        for ind, elem in enumerate(array):
            for i in range(len(array)-1, ind, -1):
                if elem+array[i] == target:
                    if elem != array[i]:
                        return [nums.index(array[i]), nums.index(elem)]
                    else:
                        foo = nums.index(elem)  
                        return [foo, nums[foo+1:].index(array[i])+foo+1]
                elif target-elem > array[i]: break