package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt: ", err)
		return
	}
	defer file.Close()

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var i1, i2 int
		_, _ = fmt.Fscanf(strings.NewReader(line), "%d %d", &i1, &i2)

		list1 = append(list1, i1)
		list2 = append(list2, i2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input.txt: ", err)
		return
	}

	sort.Ints(list1)
	sort.Ints(list2)

	// Part 1

	var diffList []int
	for i := 0; i < len(list1); i++ {
		if list1[i] >= list2[i] {
			diffList = append(diffList, list1[i]-list2[i])
		} else {
			diffList = append(diffList, list2[i]-list1[i])
		}
	}

	var sum int
	for _, v := range diffList {
		sum += v
	}

	fmt.Println("sum: ", sum)

	// Part 2

	var similarityScores []int

	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				similarityScores = append(similarityScores, list1[i])
			}
			if list1[i] < list2[j] {
				break
			}
		}
	}

	totalSimilarity := 0
	for _, v := range similarityScores {
		totalSimilarity += v
	}

	fmt.Println("totalSimilarity: ", totalSimilarity)
}
