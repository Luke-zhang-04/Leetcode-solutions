def twoSum(self, nums, target):
  for i in nums:
      if target-i in nums:
        if nums.count(target-i) >= 2:
          indices = [index for index, x in enumerate(nums) if x == target-i]
          return [indices[0], indices[1]]
        elif nums.index(i) == nums.index(target-i): continue
        else: return [nums.index(i), nums.index(target-i)]

#Runtime: 804 ms, faster than 30.58% of Python3 online submissions for Two Sum.
#Memory Usage: 13.8 MB, less than 72.79% of Python3 online submissions for Two Sum.