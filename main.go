package main

import (
    "slices"
    "strconv"
    "strings"
    "math"
    "fmt"
    "os"
)

func partOne() {
    fContent, _ := os.ReadFile("day1.txt")

    var lnums []int
    var rnums []int

    lines := strings.Split(string(fContent), "\n")
    for _, l := range lines {
        parts := strings.Split(l, "   ")
        if len(parts) == 2 {
            l, _ := strconv.Atoi(parts[0])
            r, _ := strconv.Atoi(parts[1])
            lnums = append(lnums, l)
            rnums = append(rnums, r)
        }
    }
    slices.Sort(lnums)
    slices.Sort(rnums)

    var diff float64

    for i, v := range lnums {
        diff += math.Abs(float64(v - rnums[i]))
    }

    fmt.Printf("Diff: %.0f \n", diff)
}

func main() {
    partOne()
}
