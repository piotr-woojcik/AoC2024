package main

import (
	"bufio"
	"fmt"
	"os"
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

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	count := 0
	directions := [][2]int{
		{0, 1},   // E
		{1, 0},   // S
		{1, 1},   // SE
		{1, -1},  // SW
		{0, -1},  // W
		{-1, 0},  // N
		{-1, 1},  // NE
		{-1, -1}, // NW
	}

	target := "XMAS"
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if checkWord(matrix, i, j, dx, dy, target) {
					count++
				}
			}
		}
	}

	fmt.Println("Part 1:", count)
}

func checkWord(matrix [][]rune, startX, startY, dx, dy int, target string) bool {
	rows, cols := len(matrix), len(matrix[0])
	if startX < 0 || startY < 0 || startX >= rows || startY >= cols {
		return false
	}

	for i := 0; i < len(target); i++ {
		x, y := startX+i*dx, startY+i*dy
		if x < 0 || x >= rows || y < 0 || y >= cols {
			return false
		}
		if matrix[x][y] != rune(target[i]) {
			return false
		}
	}
	return true
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt: ", err)
		return
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	count := 0
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if matrix[i][j] == 'A' {
				word := []rune{matrix[i-1][j-1], matrix[i][j], matrix[i+1][j+1]}
				word2 := []rune{matrix[i-1][j+1], matrix[i][j], matrix[i+1][j-1]}

				if checkMasWord(word) && checkMasWord(word2) {
					count++
				}

			}
		}
	}

	fmt.Println("Part 2:", count)

}

func checkMasWord(word []rune) bool {
	str := string(word)

	if str == "MAS" || str == "SAM" {
		return true
	}

	return false
}
