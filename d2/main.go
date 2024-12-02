package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	part1()
	part2()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt: ", err)
		return
	}
	defer file.Close()

	var safeLines int
	var lineList []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineList = strings.Split(line, " ")

		safe := isLineSafe(lineList)

		if safe {
			safeLines++
		}
	}

	fmt.Println("Part 1: ", safeLines)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt: ", err)
		return
	}
	defer file.Close()

	var safeLines int
	var lineList []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineList = strings.Split(line, " ")

		safe := isLineSafe(lineList)

		if safe {
			safeLines++
		} else {
			for i := 0; i < len(lineList); i++ {
				var lineTest []string
				for j := 0; j < len(lineList); j++ {
					if i != j {
						lineTest = append(lineTest, lineList[j])
					}
				}

				safe = isLineSafe(lineTest)

				if safe {
					safeLines++
					break
				}
			}
		}
	}

	fmt.Println("Part 2: ", safeLines)
}

func isLineSafe(line []string) bool {

	decreasing := false
	var diff int

	for i := 1; i < len(line); i++ {
		i1, _ := strconv.Atoi(line[i-1])
		i2, _ := strconv.Atoi(line[i])

		if i == 1 {
			decreasing = i1 > i2
		}

		if decreasing {
			diff = i1 - i2
		} else {
			diff = i2 - i1
		}

		if diff < 1 || diff > 3 {
			return false
		}

	}

	return true
}
