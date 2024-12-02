package main

import (
    "strconv"
    "regexp"
)

func extractNums(s string) []int {
    re := regexp.MustCompile(`\d+`)
    matches := re.FindAllString(s, -1)

    var nums []int
    for _, m := range matches {
        num, _ := strconv.Atoi(m)
        nums = append(nums, num)
    }

    return nums
}
