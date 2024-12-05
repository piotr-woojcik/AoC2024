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
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt:", err)
		return
	}
	defer file.Close()

	dependencies := make(map[int]map[int]bool)
	var updates [][]int
	rulesEnd := false

	var invalidUpdates [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			rulesEnd = true
			continue
		}

		if !rulesEnd {
			parts := strings.Split(line, "|")
			before, err1 := strconv.Atoi(parts[0])
			after, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing rule:", line)
				continue
			}

			if _, ok := dependencies[after]; !ok {
				dependencies[after] = make(map[int]bool)
			}
			dependencies[after][before] = true
		} else {
			var update []int
			for _, numStr := range strings.Split(line, ",") {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					fmt.Println("Error parsing update number:", numStr)
					continue
				}
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sum := 0
	for _, update := range updates {
		if isValidOrder(update, dependencies) {
			sum += update[len(update)/2]
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	fmt.Println("Part 1:", sum)

	part2(invalidUpdates, dependencies)
}

func isValidOrder(numbers []int, dependencies map[int]map[int]bool) bool {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if deps, ok := dependencies[numbers[i]]; ok {
				if deps[numbers[j]] {
					return false
				}
			}
		}
	}
	return true
}

func quickSort(arr []int, dependencies map[int]map[int]bool, low, high int) {
	if low < high {
		pivot := partition(arr, dependencies, low, high)
		quickSort(arr, dependencies, low, pivot-1)
		quickSort(arr, dependencies, pivot+1, high)
	}
}

func partition(arr []int, dependencies map[int]map[int]bool, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		shouldComeBefore := true
		if deps, ok := dependencies[arr[j]]; ok {
			if deps[pivot] {
				shouldComeBefore = false
			}
		}
		if deps, ok := dependencies[pivot]; ok {
			if deps[arr[j]] {
				shouldComeBefore = true
			}
		}

		if shouldComeBefore {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func findValidOrderQuick(numbers []int, dependencies map[int]map[int]bool) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)
	quickSort(result, dependencies, 0, len(result)-1)
	return result
}

func part2(invalidUpdates [][]int, dependencies map[int]map[int]bool) {
	sum := 0

	for _, update := range invalidUpdates {
		orderedUpdate := findValidOrderQuick(update, dependencies)
		middleNum := orderedUpdate[len(orderedUpdate)/2]
		sum += middleNum
	}

	fmt.Println("Part 2:", sum)
}
