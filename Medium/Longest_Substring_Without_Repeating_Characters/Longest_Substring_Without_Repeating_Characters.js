// Runtime: 100 ms, faster than 62.03% of JavaScript online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 41.5 MB, less than 29.20% of JavaScript online submissions for Longest Substring Without Repeating Characters.
// 13/02/2020
// DD/MM/YY

/*
 * @param {string} s
 * @return {number}
 */

function set(arr) { 
    let all = []
    for (i = 0; i < arr.length; i++) {
        if (!all.includes(arr[i])) {
            all.push(arr[i])
        }
    }
    return all
}

var lengthOfLongestSubstring = function(s) {
    let letters = set(s.split(""))
    let longest = 0
    let subString
    if (letters.length <= 2) {
        return letters.length
    } else {
        for (index = 0; index < s.length; index++) {
            subString = [s[index]]
            for (subIndex = 0; subIndex < s.slice(index+1).length; subIndex++) {
                if (!subString.includes(s.slice(index+1)[subIndex])) {
                    subString.push(s.slice(index+1)[subIndex])
                } else {
                    break
                }
            }
            if (subString.length > longest) {
                    longest = subString.length
            }
            if (longest === letters.length || index >= s.length-longest) {
                break
            }
        }
    }
    return longest
}