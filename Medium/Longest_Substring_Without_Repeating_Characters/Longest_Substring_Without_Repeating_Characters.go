// Runtime: 20 ms, faster than 31.05% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 6.4 MB, less than 20.00% of Go online submissions for Longest Substring Without Repeating Characters.
// 13/02/2020
// DD/MM/YY

package lengthOfLongestSubstring

import "strings"

func contains(target string, list []string) bool {
    for _, i := range list {
        if target == i {
            return true
        }
    }
  return false
}

func set(arr []string) []string {
    all := make([]string, 0, 0)
    for _, elem := range arr {
        if !contains(elem, all) {
            all = append(all, elem)
        }
    }
    return all
}


func lengthOfLongestSubstring(s string) int {
    letters := set(strings.Split(s, ""))
    longest := 0
    if len(letters) <= 2 {
        return len(letters)
    } else {
        for index, letter := range s {
            subString := append(make([]string, 0, 0), string(letter))
            for _, subLetter := range s[index+1:]{
                if !contains(string(subLetter), subString) {
                    subString = append(subString, string(subLetter))
                } else {
                    break
                }
            }
            if len(subString) > longest {
                longest = len(subString)
            }
            if longest == len(letters) || index >= len(s)-longest {
                break
            }
        }
    }
    return longest
}