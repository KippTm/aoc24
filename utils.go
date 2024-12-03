package main

import (
    "strconv"
    "strings"
    "regexp"
    "os"
)

func extractLines(filePath string) []string {
    contents, _ := os.ReadFile(filePath)
    lines := strings.Split(string(contents), "\n")
    return lines
}

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

func removeIdx(nums []int, idx int) []int {
    copyNums := make([]int, len(nums))
    copy(copyNums, nums)
    copyNums = append(copyNums[:idx], copyNums[idx+1:]...)

    return copyNums
}
