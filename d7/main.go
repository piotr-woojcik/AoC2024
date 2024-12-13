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
	part1()
	part2()
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		testValue, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Fatal(err)
		}

		elements := strings.Fields(strings.TrimSpace(parts[1]))
		var nums []int
		for _, element := range elements {
			num, err := strconv.Atoi(element)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}

		var operatorCombinations [][]byte
		generateOperators(len(nums)-1, []byte{}, &operatorCombinations)

		for _, operators := range operatorCombinations {
			if evaluateExpression(nums, operators) == testValue {
				sum += testValue
				break
			}
		}
	}

	fmt.Println("Part 1:", sum)
}

func evaluateExpression(nums []int, operators []byte) int {
	result := nums[0]
	for i := 0; i < len(operators); i++ {
		if operators[i] == '+' {
			result += nums[i+1]
		} else {
			result *= nums[i+1]
		}
	}
	return result
}

func generateOperators(n int, current []byte, result *[][]byte) {
	if len(current) == n {
		combination := make([]byte, n)
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	operators := []byte{'+', '*'}
	for _, op := range operators {
		generateOperators(n, append(current, op), result)
	}
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		testValue, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Fatal(err)
		}

		elements := strings.Fields(strings.TrimSpace(parts[1]))
		var nums []int
		for _, element := range elements {
			num, err := strconv.Atoi(element)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}

		var operatorCombinations [][]byte
		generateOperatorsWithConcat(len(nums)-1, []byte{}, &operatorCombinations)

		for _, operators := range operatorCombinations {
			if evaluateExpressionWithConcat(nums, operators) == testValue {
				sum += testValue
				break
			}
		}
	}

	fmt.Println("Part 2:", sum)
}

func concatenate(a, b int) int {
	bStr := strconv.Itoa(b)
	result, _ := strconv.Atoi(strconv.Itoa(a) + bStr)
	return result
}

func generateOperatorsWithConcat(n int, current []byte, result *[][]byte) {
	if len(current) == n {
		combination := make([]byte, n)
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	operators := []byte{'+', '*', '|'} // Added concatenation operator
	for _, op := range operators {
		generateOperatorsWithConcat(n, append(current, op), result)
	}
}

func evaluateExpressionWithConcat(nums []int, operators []byte) int {
	result := nums[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case '+':
			result += nums[i+1]
		case '*':
			result *= nums[i+1]
		case '|':
			result = concatenate(result, nums[i+1])
		}
	}
	return result
}
