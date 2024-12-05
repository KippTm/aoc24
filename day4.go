package main 

import (
    "fmt"
)

func createArr(lines []string) [][]string {
    arr := make([][]string, len(lines))
    for i := range arr {
        arr[i] = make([]string, len(lines[0]))
    }
    for i, l := range lines {
        for j, c := range l {
            arr[i][j] = string(c)
        }
    } 
    return arr
}

func findMas() {
    xmap := createArr(extractLines("day4.txt"))
    var xmas int

    for i, l := range xmap {
        for j, c := range l {
            if c != "X" {
               continue 
            }
            //Horizontal forward
            if j <= (len(l)-4) &&  l[j+1] == "M" && l[j+2] == "A" && l[j+3] == "S" {
                xmas++
            }
            //Horizontal backwards
            if j >= 3 && l[j-1] == "M" && l[j-2] == "A" && l[j-3] == "S" {
                xmas++
            }
            //Vertical down
            if i <= (len(xmap)-4) && xmap[i+1][j] == "M" && xmap[i+2][j] == "A" && xmap[i+3][j] == "S" {
                xmas++
            }

            //Vertical up
            if i >= 3 && xmap[i-1][j] == "M" && xmap[i-2][j] == "A" && xmap[i-3][j] == "S" {
                xmas++
            }

            //Diagonal up right
            if i >= 3 && j <= (len(l)-4) && xmap[i-1][j+1] == "M" && xmap[i-2][j+2] == "A" && xmap[i-3][j+3] == "S" {
                xmas++
            }

            //Diagonal up left
            if i >= 3 && j >= 3 && xmap[i-1][j-1] == "M" && xmap[i-2][j-2] == "A" && xmap[i-3][j-3] == "S" {
                xmas++
            }

            //Diagonal down right
            if i <= (len(xmap)-4) && j <= (len(l)-4) && xmap[i+1][j+1] == "M" && xmap[i+2][j+2] == "A" && xmap[i+3][j+3] == "S" {
                xmas++
            }

            //Diagonal down left
            if i <= (len(xmap)-4) && j >= 3 && xmap[i+1][j-1] == "M" && xmap[i+2][j-2] == "A" && xmap[i+3][j-3] == "S" {
                xmas++
            }
        }
    }
    fmt.Printf("Xmax count: %d\n", xmas)
}

func findXMs() {
    xmap := createArr(extractLines("day4.txt"))
    var xMAS int

    for i, l := range xmap {
        for j, c := range l {
            if c == "A" {
                if j < 1 || j > len(xmap[0])-2 || i < 1 || i > len(xmap)-2{
                    continue
                }
                /* M . M
                     A
                   S . S */
                if xmap[i-1][j-1] == "M" && xmap[i-1][j+1] == "M" && xmap[i+1][j-1] == "S" && xmap[i+1][j+1] == "S" {
                    xMAS++
                }

                /* M . S
                     A
                   M . S */
                if xmap[i-1][j-1] == "M" && xmap[i-1][j+1] == "S" && xmap[i+1][j-1] == "M" && xmap[i+1][j+1] == "S" {
                    xMAS++
                }

                /* S . M
                     A
                   S . M */
                if xmap[i-1][j-1] == "S" && xmap[i-1][j+1] == "M" && xmap[i+1][j-1] == "S" && xmap[i+1][j+1] == "M" {
                    xMAS++
                }

                /* S . S
                     A
                   M . M */
                if xmap[i-1][j-1] == "S" && xmap[i-1][j+1] == "S" && xmap[i+1][j-1] == "M" && xmap[i+1][j+1] == "M" {
                    xMAS++
                }
            }
        }
    }
    fmt.Printf("X-MAS count: %d\n", xMAS)
}
