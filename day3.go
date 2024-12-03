package main 

import (
    "strings"
    "regexp"
    "fmt"
)

func extractMuls(l string) []string {
    re := regexp.MustCompile(`mul\(\d+,\d+\)`)
    muls := re.FindAllString(l, -1)
    return muls
}

func extractAll(l string) []string {
    re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
    instructions := re.FindAllString(l, -1)
    return instructions
}

func mulSum(m string) int {
    nums := extractNums(m)

    if len(nums) != 2 {
        return 0
    }
    return nums[0]*nums[1]
}

func calcMuls(l string) int {
    var count int
    for _, m := range extractMuls(l) {
        count += mulSum(m)
    }
    return count
}

func checkEnabled(l string, run *bool) int {
    var count int

    for _, i := range extractAll(l) {
        if strings.Contains(i, "don") {
            *run = false
            continue
        } else if strings.Contains(i, "do(") {
            *run = true
            continue
        } else if *run {
            count += mulSum(i)
        }
    }
    return count
}

func day3() {
    var memory int
    var doMemory int
    run := true
    for _,l := range extractLines("day3.txt") {
        memory += calcMuls(l)
        doMemory += checkEnabled(l, &run)
    }

    fmt.Printf("Mul num: %d\n", memory)
    fmt.Printf("Enabled mul num: %d\n", doMemory)
}
