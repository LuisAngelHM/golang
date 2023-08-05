package main

import "fmt"


func jumpingOnClouds(c []int32) int32 {
    // Write your code here
    var jumps int32
    jumps = 0
    index := 0
    for true {
        if (index == len(c)-1) {
            return jumps
        } 

        if ((index + 2) < len(c) && c[index + 2] != 1) {
            index = index + 2
        } else {
            index = index + 1
        }

        jumps = jumps + 1
    }
    
    return jumps
}

func main() {
    slice :=[]int32{0,0,0,1,0,0}
    fmt.Println(jumpingOnClouds(slice))
}