package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func decomposeMatrix(matrix [][]int) int {

    max_val := 0
    max_x := 0
    max_y := 0

    // find the max value
    for x := 0; x < len(matrix[0][:]); x++ {
        for y := 0; y < len(matrix[:]); y++ {
            if matrix[x][y] > max_val {
                max_val = matrix[x][y];
                max_x = x;
                max_y = y;
            }
        }
    }

    if max_val == 0 {
        // end condition
        return 0;
    }

    // set row to 0
    for x := 0; x < len(matrix); x++ {
        matrix[x][max_y] = 0;
    }

    // set col to 0
    for y := 0; y < len(matrix); y++ {
        matrix[max_x][y] = 0;
    }

    return max_val + decomposeMatrix(matrix);
}

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

    matrix_length, _ := strconv.Atoi(os.Args[2]);
    var matrix [][]int = make([][]int, matrix_length)
    for i := range matrix {
        matrix[i] = make([]int, matrix_length)
    }
    fmt.Println(len(matrix), len(matrix[0][:]))

    readInMatrix(file, &matrix)



    sum := decomposeMatrix(matrix)
    fmt.Println("Matrix sum is: ", sum)

    // testing
    // for i := 0; i < matrix_length; i++ {
    //     for j := 0; j < matrix_length; j++ {
    //         fmt.Print(matrix[i][j], " ")
    //     }
    //     fmt.Println("")
    // }
}
