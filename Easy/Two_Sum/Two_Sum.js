// Runtime: 52 ms, faster than 94.04% of JavaScript online submissions for Two Sum.
// Memory Usage: 35.8 MB, less than 16.12% of JavaScript online submissions for Two Sum.
// 02/24/2020
// MM/DD/YY

function shell_sort(array) {
  let gap = Math.floor(array.length/2)
  while (gap >= 1) {
      //console.log(array, gap)
      for (i = gap; i < array.length; i ++) { //iterate through array, starting from gap
          let comparator = array[i] //make comparisons with this
          let index //in case of negative index
          let output //for accessing x outside the array

          for (x = i; x > gap-2; x -= gap) { //iterate throguh array with gap as the step
              output = x //for accessing x outside the array
              if (x-gap < 0) { //in case of negative index
                  index = array.length-x-gap
              } else {
                  index = x-gap
              }
              
              if (array[index] <= comparator) { //break when correct spot is found
                  break 
              } else { //otherwise, move elements forward to make space
                  array[x] = array[index] 
              }
          }
          array[output] = comparator //insert comparator in the correct spot
      }
      gap = Math.floor(gap/2) //increment the gap
  }
  return array
}

var twoSum = function(nums, target) {
  let array = shell_sort(nums.slice())
  for (ind = 0; ind < array.length; ind ++) {
      for (i = array.length-1; i > ind; i--) {
          if (array[ind] + array[i] === target) {
              if (array[ind] != array[i]) {
                  return [nums.indexOf(array[i]), nums.indexOf(array[ind])]
              } else {
                  let foo = nums.indexOf(array[ind])  
                  return [foo, nums.slice(foo+1).indexOf((array[i]))+foo+1]
              }
          } s
      }
  }
}