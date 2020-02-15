#Runtime: 68 ms, faster than 53.34% of Python3 online submissions for Longest Substring Without Repeating Characters.
#Memory Usage: 12.8 MB, less than 100.00% of Python3 online submissions for Longest Substring Without Repeating Characters.
#13/02/2020
#DD/MM/YY

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        letters = set(list(s))        
        if len(letters) <= 2: return len(letters)
        else:
            longest = 0
            for index, letter in enumerate(s):
                subString = [letter]
                for subLetter in s[index+1:]:
                    if subLetter not in subString:
                        subString.append(subLetter)
                    else:
                        break
                if len(subString) > longest:
                    longest = len(subString)
                if longest == len(letters) or index >= len(s)-longest:
                    break
        return longest