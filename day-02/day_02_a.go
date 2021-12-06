package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var x = 0
	var y = 0

	const (
		UP      = "up"
		DOWN    = "down"
		FORWARD = "forward"
		BACK    = "back"
	)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		z, _ := strconv.Atoi(split[1])
		switch split[0] {
		case UP:
			y -= z
		case DOWN:
			y += z
		case FORWARD:
			x += z
		case BACK:
			x -= z
		}
	}
	fmt.Println(x * y)
}
