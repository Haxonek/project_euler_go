package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {

    if len(os.Args) != 2 {
        fmt.Println("Missing input argument, must pass through file with 100 numbers.")
        os.Exit(0)
    }

    var numbers [101]int;
    var err error;
    file_name := os.Args[1]

    // read in file with 100 numbers
    file, _ := os.Open(file_name)
    input := bufio.NewScanner(file)

    i := 0;
    for input.Scan() {
        i++;
        var in string = input.Text()[40:]
        if len(in) == 0 {
            // empty line, don't read
            continue;
        }

        numbers[i], err = strconv.Atoi(in)

        if err != nil {
            fmt.Println(err)
            os.Exit(0)
        }
    }
    file.Close()

    // sum up all the values
    var sum int = 0;
    for num := range numbers {
        sum += numbers[num];
    }

    // chop off extra digits and print first ten numbers
    final := strconv.Itoa(sum)
    start := len(final) - 10
    final = final[start:]
    fmt.Println("Sum of just the first 10 digits: ", final)
}
