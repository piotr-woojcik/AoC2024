package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Direction represents the guard's current facing direction
type Direction int
type Position struct {
	x, y int
	dir  Direction
}

const (
	Up Direction = iota
	Right
	Down
	Left
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

	var matrix [][]rune
	var visited [][]bool
	var count int

	startX, startY := -1, -1
	var startDir = Up

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
		visited = append(visited, make([]bool, len(line)))

		if startX < 0 && startY < 0 {
			for i := 0; i < len(line); i++ {
				if line[i] == '^' {
					startX = i
					startY = len(matrix) - 1
					visited[startY][startX] = true
					break
				}
			}
		}
	}

	if startX < 0 || startY < 0 {
		log.Fatal("start not found")
	}

	fmt.Println("Start position:", startX, startY)

	count = walk(startX, startY, startDir, matrix, visited, 1)
	fmt.Println("Part 1:", count)
}

func walk(x, y int, dir Direction, matrix [][]rune, visited [][]bool, count int) int {
	for {
		if x < 0 || x >= len(matrix[0]) || y < 0 || y >= len(matrix) {
			return count
		}

		if !visited[y][x] {
			visited[y][x] = true
			count++
		}

		nextX, nextY := getNextPosition(x, y, dir)
		if isOutOfBounds(nextX, nextY, matrix) {
			return count
		} else if isObstacle(nextX, nextY, matrix) {
			dir = (dir + 1) % 4
		} else {
			x, y = nextX, nextY
		}
	}
}

func getNextPosition(x, y int, dir Direction) (int, int) {
	switch dir {
	case Up:
		return x, y - 1
	case Right:
		return x + 1, y
	case Down:
		return x, y + 1
	case Left:
		return x - 1, y
	default:
		return x, y
	}
}

func isOutOfBounds(x, y int, matrix [][]rune) bool {
	return x < 0 || x >= len(matrix[0]) || y < 0 || y >= len(matrix)
}

func isObstacle(x, y int, matrix [][]rune) bool {
	return matrix[y][x] == '#'
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var matrix [][]rune
	startX, startY := -1, -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))

		if startX < 0 && startY < 0 {
			for i := 0; i < len(line); i++ {
				if line[i] == '^' {
					startX = i
					startY = len(matrix) - 1
					break
				}
			}
		}
	}

	if startX < 0 || startY < 0 {
		log.Fatal("start not found")
	}

	loopPositions := findLoopPositions(matrix, startX, startY)
	fmt.Println("Part 2:", len(loopPositions))
}

func findLoopPositions(matrix [][]rune, startX, startY int) []Position {
	var loopPositions []Position

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[y][x] == '#' || (x == startX && y == startY) {
				continue
			}

			tempMatrix := copyMatrix(matrix)
			tempMatrix[y][x] = '#'

			if hasLoop(tempMatrix, startX, startY) {
				loopPositions = append(loopPositions, Position{x, y, 0})
			}
		}
	}

	return loopPositions
}

func hasLoop(matrix [][]rune, startX, startY int) bool {
	visited := make(map[Position]int)
	x, y := startX, startY
	dir := Up

	for step := 0; ; step++ {
		if x < 0 || x >= len(matrix[0]) || y < 0 || y >= len(matrix) {
			return false
		}

		pos := Position{x, y, dir}
		if _, exists := visited[pos]; exists {
			return true
		}
		visited[pos] = step

		nextX, nextY := getNextPosition(x, y, dir)
		if isOutOfBounds(nextX, nextY, matrix) {
			return false
		} else if isObstacle(nextX, nextY, matrix) {
			dir = (dir + 1) % 4
		} else {
			x, y = nextX, nextY
		}
	}
}

func copyMatrix(matrix [][]rune) [][]rune {
	newMatrix := make([][]rune, len(matrix))
	for i := range matrix {
		newMatrix[i] = make([]rune, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}
	return newMatrix
}
