package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func readInMatrix(f *os.File, matrix *[][]int) {
    // var matrix [15][15]int;
    input := bufio.NewScanner(f);

    for x := 0; input.Scan(); x++ {
        y := 0;
        for _, number := range strings.Split(string(input.Text()), " ") {
            if (len(number) == 0) {
                // just an empty space, iterate over it
                continue;
            }
            (*matrix)[x][y], _ = strconv.Atoi(number);
            y++
        }
    }

}

func main() {
    // read in matrix from file
    if len(os.Args) != 3 {
        fmt.Println("You must include both the file name of the square matrix,",
            "plus its length as a command line argument. An example using the",
            "example file is shown below.")
        fmt.Println("Example: go run main.go matrix.txt 15")
        os.Exit(0)
    }

    var file_name string = os.Args[1]
    file, er := os.Open(file_name)
    defer file.Close() // closes the file once it's no longer needed

    if er != nil {
        fmt.Fprintf(os.Stderr, "%v\n", er)
        os.Exit(0)
    }

    n, _ := strconv.Atoi(os.Args[2]); // size of array, write func to calcualte later
    var matrix [][]int = make([][]int, n)
    for i := range matrix {
        matrix[i] = make([]int, n)
    }
    fmt.Println(len(matrix), len(matrix[0][:]))

    readInMatrix(file, &matrix)


    // testing
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Print(matrix[i][j], " ")
        }
    fmt.Println("")
    }
}
