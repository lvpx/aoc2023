package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int

	f, err := os.Open("2.txt")
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(f)

	resultSet := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for sc.Scan() {
		s := sc.Text()
		game := strings.Split(s, ":")
		id, _ := strconv.Atoi(strings.Fields(game[0])[1])
		draws := strings.Split(game[1], ";")
		for _, draw := range draws {
			cubes := strings.Split(draw, ",")
			hash := make(map[string]int)
			for _, cube := range cubes {
				values := strings.Fields(cube)
				key := values[1]
				num, _ := strconv.Atoi(values[0])
				hash[key] += num
			}
			if feasible(resultSet, hash) {
				fmt.Println(id)
				sum += id
			}
		}
	}

	fmt.Println(sum)
}

func feasible(resultSet map[string]int, hash map[string]int) bool {
	fmt.Println(resultSet, hash)
	for k, v := range hash {
		if v > resultSet[k] {
			return false
		}
	}

	return true
}
