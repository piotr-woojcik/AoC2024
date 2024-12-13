package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	part1()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	//var matrix [][]rune
	//
	//scanner := bufio.NewScanner(file)
	//for scanner.Scan() {
	//	line := scanner.Text()

	//}

	fmt.Println("Part 1:", sum)
}
