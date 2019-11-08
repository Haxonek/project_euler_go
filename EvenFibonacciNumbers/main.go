package main

import (
    "fmt"
)

func findEvenFibb(a uint32, b uint32) uint32 {
    var cur uint32 = a + b;

    if cur % 2 == 0 {
        // even number, return
        return cur;
    } else {
        // odd number, iterate through again
        if a < b {
            return findEvenFibb(b, cur)
        } else {
            return findEvenFibb(a, cur)
        }
    }
}

func main() {

    var cur uint32 = 1
    var last uint32 = 1
    var MAX_VAL uint32 = 4000000;
    count := 0

    fmt.Println("Starting fibonacci search")
    for cur < MAX_VAL {
        tmp := cur
        cur = last
        last = tmp
        cur = findEvenFibb(cur, last)
        count++;
    }

    fmt.Println("There were ",count," even fib numbers, and the last two were ", last, " and ", cur)

}
