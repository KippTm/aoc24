package main 

import (
    "regexp"
    "fmt"
)

func extractMuls(l string) int {
    re := regexp.MustCompile(`mul\(\d+,\d+\)`)
    muls := re.FindAllString(l, -1)
    var count int
    for _, m := range muls {
        nums := extractNums(m)
        if len(nums) != 2 {
            return 0
        }
        count += nums[0]*nums[1]
    }
    return count
}

func day3() {
    var memory int
    for _,l := range extractLines("day3.txt") {
        memory += extractMuls(l)
    }

    fmt.Printf("Total instructions num: %d", memory)
}
