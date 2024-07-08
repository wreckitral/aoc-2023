package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solve(input []string) {

    re := regexp.MustCompile(`\d`)

    sum := 0

    for _, in := range input {
        subMatch := re.FindAllString(in, -1)

        n := 0

        if len(subMatch) < 2 {
            comb := subMatch[0] + subMatch[0]
            n, _ = strconv.Atoi(comb)
            fmt.Println(n)
        } else {
            comb := subMatch[0] + subMatch[len(subMatch) - 1]
            n, _ = strconv.Atoi(comb)
            fmt.Println(n)
        }

        sum += n
    }

    fmt.Println(sum)
}

func main() {
    d, _ := os.ReadFile("input.txt")
    data := strings.TrimSpace(string(d))
    input := strings.Split(data, "\n")

    solve(input)
}
