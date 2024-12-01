package main

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    arr1, arr2 := make([]int, 0, 1000), make([]int, 0, 1000)
    for scanner.Scan() {
        text := scanner.Text()
        if text == "" { break }
        splits := strings.Split(text, " ")
        ai, _ := strconv.Atoi(splits[0])
        bi, _ := strconv.Atoi(splits[1])
        arr1 = append(arr1, ai)
        arr2 = append(arr2, bi)
    }

    slices.Sort(arr1)
    slices.Sort(arr2)

    total := 0
    for i := 0; i < len(arr1); i++ {
        total += int(math.Abs(float64(arr1[i] - arr2[i])))
    }

    println(total)
}
