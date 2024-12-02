package main 

import (
    "strings"
    "fmt"
    "os"
)

func isSafe(line []int) bool {
    inc, dec := false, false
    
    for i := 1; i < len(line); i++ {
        diff := line[i] - line[i-1]

        if (diff > 0) {
            inc = true
        } else if diff < 0 {
            dec = true
        } else {
            return false 
        }

        if inc && dec {
            return false
        }
        if diff > 3 || diff < -3 {
            return false
        }
    }
    return true
}


func day2() {
    input, _ := os.ReadFile("day2.txt")

    lines := strings.Split(string(input), "\n")
    var safe int
    for _, l := range lines {
        nums := extractNums(l)
        if isSafe(nums) && len(nums) != 0 {
            safe++
        }
    }

    fmt.Printf("Safe: %d", safe)
}
