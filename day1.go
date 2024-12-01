package main

import (
    "slices"
    "strconv"
    "strings"
    "math"
    "fmt"
    "os"
)

func createLists() ([]int, []int) {
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
    return lnums, rnums
}

func partOne() {
    lnums, rnums := createLists()

    slices.Sort(lnums)
    slices.Sort(rnums)

    var diff float64

    for i, v := range lnums {
        diff += math.Abs(float64(v - rnums[i]))
    }

    fmt.Printf("Diff: %.0f \n", diff)
}

func partTwo() {
    lnums, rnums := createLists()

    var sim int

    for _, l := range lnums {
        count := 0
        for _, r := range rnums {
            if r == l {
                count++
            }
        }
        sim += count*l
    }
    fmt.Printf("Simi: %d\n", sim)
}

func main() {
    partOne()
    partTwo()
}
