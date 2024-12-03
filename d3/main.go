package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	var results []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		for _, match := range matches {
			var i1, i2 int
			_, _ = fmt.Sscanf(match, "mul(%d,%d)", &i1, &i2)
			results = append(results, i1*i2)
		}
	}

	final := 0
	for _, result := range results {
		final += result
	}
	fmt.Println("Part 1:", final)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt: ", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(content)
	var results []int

	enabled := true
	pos := 0

	for pos < len(text) {
		switch {
		case strings.HasPrefix(text[pos:], "do()"):
			enabled = true
			pos += 4
		case strings.HasPrefix(text[pos:], "don't()"):
			enabled = false
			pos += 7
		case strings.HasPrefix(text[pos:], "mul("):
			if enabled {
				endPos := strings.Index(text[pos:], ")")
				if endPos != -1 {
					mulStr := text[pos : pos+endPos+1]
					var i1, i2 int
					if n, err := fmt.Sscanf(mulStr, "mul(%d,%d)", &i1, &i2); err == nil && n == 2 {
						results = append(results, i1*i2)
						pos += len(mulStr)
						continue
					}
				}
			}
			pos++
		default:
			pos++
		}
	}

	final := 0
	for _, result := range results {
		final += result
	}
	fmt.Println("Part 2:", final)
}
