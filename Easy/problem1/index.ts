#!/usr/bin/env deno

/**
 * Runtime: 184 ms, faster than 22.83% of TypeScript online submissions for Two Sum.
 * Memory Usage: 37.3 MB, less than 51.18% of TypeScript online submissions for Two Sum.
 * 16/08/2019
 * DD/MM/YY
 */

const count = <T>(target: T, list: T[]): number => {
    let total = 0

    for (let i = 0; i < list.length; i++) {
        if (list[i] === target) {
            total ++
        }
    }

    return total
}

const indice = <T>(target: T, list: T[]): number[] => {
    let indices: number[] = [],
        i = -1

    while ((i = list.indexOf(target, i + 1)) !== -1) {
        indices.push(i)
    }

    return indices
}

const twoSum = (nums: number[], target: number): [number, number] | void => {
    for (let x = 0; x < nums.length; x++) {
        let i = nums[x]

        if (nums.includes(target-i)) {
            if (count(target-i, nums) >= 2) {
                let indices = indice(target-i, nums)

                return [indices[0], indices[1]]
            } else if (nums.indexOf(i) === nums.indexOf(target-i)) {
                continue
            } else {
                return [nums.indexOf(i), nums.indexOf(target-i)]
            }
        }
    }
}
