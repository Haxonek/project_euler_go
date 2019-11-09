package main

import (
    "fmt"
    "strconv"
)

func returns89(a int) (bool) {
    var str string = strconv.Itoa(a)
    sum := 0

    for i := 0; i < len(str); i++ {
        val, _ := strconv.Atoi(str[i:i+1])
        sum += val*val;
    }

    if sum == 89 {
        return true;
    } else if sum == 1 {
        return false;
    } else {
        return returns89(sum)
    }
}

func main() {
    fmt.Printf("Starting Square Digit Chains\n")

    var MAX_NUM int = 10000000
    count := 0

    for i := 2; i <= MAX_NUM; i++ {
        if returns89(i) {
            count++;
        }
    }

    fmt.Println(count, "numbers below 10 mill end at 89, or", 100*(float32(count)/float32(MAX_NUM)), "percent")
}
