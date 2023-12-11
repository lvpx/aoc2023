package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var sum int

	f, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		sum += calibrate(sc.Text())
	}

	fmt.Println(sum)
}

func calibrate(line string) int {
	var first, last int
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			first, _ = strconv.Atoi(string(ch))
			break
		}
	}

	runeLine := []rune(line)
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(runeLine[i]) {
			last, _ = strconv.Atoi(string(runeLine[i]))
			break
		}
	}

	fmt.Println(first, last)
	return first*10 + last
}
