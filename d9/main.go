package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

	var line string
	var input []string

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line = scanner.Text()

	isEmpty := false
	currNum := 0

	for _, c := range line {

		num := int(c - '0')

		if isEmpty {
			for i := 0; i < num; i++ {
				input = append(input, ".")
			}
			isEmpty = false
		} else {
			for i := 0; i < num; i++ {
				input = append(input, strconv.Itoa(currNum))
			}
			isEmpty = true
			currNum += 1
		}
	}

	var inputSlice []string

	replaceIdx := len(input) - 1
	for i := 0; i < len(input); i++ {

		if input[i] == "." {
			for input[replaceIdx] == "." {
				replaceIdx -= 1

				if replaceIdx <= i {
					break
				}
			}
			inputSlice = append(inputSlice, input[replaceIdx])
			replaceIdx -= 1
		} else {
			inputSlice = append(inputSlice, input[i])
		}

		if replaceIdx <= i {
			break
		}
	}

	var sum int

	idx := 0
	for idx < len(inputSlice) && inputSlice[idx] != "." {
		number, _ := strconv.Atoi(inputSlice[idx])
		sum += number * idx
		idx += 1
	}

	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", part2(line))
}

type File struct {
	id    int
	size  int
	start int
}

func part2(line string) int {

	var files []File
	var disk []string
	pos := 0
	fileId := 0
	isEmpty := false

	for _, c := range line {
		size := int(c - '0')
		if isEmpty {
			for i := 0; i < size; i++ {
				disk = append(disk, ".")
			}
			pos += size
		} else {
			files = append(files, File{
				id:    fileId,
				size:  size,
				start: pos,
			})
			for i := 0; i < size; i++ {
				disk = append(disk, strconv.Itoa(fileId))
			}
			pos += size
			fileId++
		}
		isEmpty = !isEmpty
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].id > files[j].id
	})

	for _, f := range files {
		freeSpaceStart := -1
		consecutiveFree := 0
		for i := 0; i < len(disk); i++ {
			if disk[i] == "." {
				if consecutiveFree == 0 {
					freeSpaceStart = i
				}
				consecutiveFree++
				if consecutiveFree >= f.size {
					break
				}
			} else {
				consecutiveFree = 0
				freeSpaceStart = -1
			}
		}

		if freeSpaceStart != -1 && freeSpaceStart < f.start {
			for i := f.start; i < f.start+f.size; i++ {
				disk[i] = "."
			}
			for i := 0; i < f.size; i++ {
				disk[freeSpaceStart+i] = strconv.Itoa(f.id)
			}
		}
	}

	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] != "." {
			id, _ := strconv.Atoi(disk[i])
			sum += id * i
		}
	}

	return sum
}
