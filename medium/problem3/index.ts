#!/usr/bin/env deno

/**
 * Runtime: 108 ms, faster than 76.52% of TypeScript online submissions for Longest Substring Without Repeating Characters.
 * Memory Usage: 43.6 MB, less than 24.29% of TypeScript online submissions for Longest Substring Without Repeating Characters.
 * 16/08/2020
 * DD/MM/YY
 */

/**
 * Generate an array with no repeating values from arr
 * @param arr - base array
 * @returns set derived from array
 */
const set = <T>(arr: T[]): T[] => { 
    let all: T[] = []

    for (let i = 0; i < arr.length; i++) {
        if (!all.includes(arr[i])) {
            all.push(arr[i])
        }
    }

    return all
}

const lengthOfLongestSubstring = (s: string): number => {
    const letters = set(s.split(""))
    let longest = 0,
        subString: string[]
    
    if (letters.length <= 2) {
        return letters.length
    } else {
        for (let index = 0; index < s.length; index++) {
            subString = [s[index]]

            for (let subIndex = 0; subIndex < s.slice(index+1).length; subIndex++) {
                if (!subString.includes(s.slice(index + 1)[subIndex])) {
                    subString.push(s.slice(index + 1)[subIndex])
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
