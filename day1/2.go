package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    arr := make([]int, 0, 1000)
    counts := map[int]int{}
    for scanner.Scan() {
        text := scanner.Text()
        if text == "" { break }
        splits := strings.Split(text, "   ")
        ai, err1 := strconv.Atoi(splits[0])
        bi, err2 := strconv.Atoi(splits[1])
        if err1 != nil {
            panic(err1)
        }
        if err2 != nil {
            panic(err2)
        }

        arr = append(arr, ai)
        counts[bi]++
    }

    sim := 0

    for _, a := range arr {
        sim += a * counts[a]
    }

    println(sim)
}
