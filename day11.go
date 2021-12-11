package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

var size int
var totalFlashes int

type Matrix [10][10]int

func flash(grid *Matrix, i int, j int) {

	if grid[i][j] > 9 {
		totalFlashes++ // increase our total flashes for metrics
		grid[i][j] = 0 // reset our current location

		// top left
		if i-1 > 0 && j-1 > 0 { 
			grid[i-1][j-1]++ 
			flash(grid,i-1,j-1)
		}
		
		// top center
		if i - 1 > 0 {
			grid[i-1][j]++
			flash(grid,i-1,j)
		}
		
		// top right
		if i - 1 > 0 && j + 1 < size {
			grid[i-1][j+1]++
			flash(grid,i-1,j+1)
		}

		// middle left
		if j - 1 > 0 {
			grid[i][j-1]++
			flash(grid,i,j-1)
		}
		
		// middle right
		if j + 1 < size {
			grid[i][j+1]++
			flash(grid,i,j+1)
		}

		// bottom left
		if i + 1 < size && j - 1 > 0 {
			grid[i+1][j-1]++
			flash(grid,i+1,j-1)
		}
		
		// bottom center
		if i + 1 < size {
			grid[i+1][j]++
			flash(grid,i+1,j)
		}

		// bottom right
		if i + 1 < size && j + 1 < size {
			grid[i+1][j+1]++
			flash(grid,i+1,j+1)
		}
		
	}
}

// https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	// storage
	totalFlashes = 0
	var grid Matrix

	// read input file into our grid
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for x, line := range lines {
		for y, num := range line {
			grid[x][y] = int(num)
		}
	}

	// calculate size and pointer to grid
	size := len(grid[0])
	var pa *Matrix = &grid

	for x := 0; x < 100; x++ {
		// Increase all by +1
		var i, j int
		for i = 0; i < size; i++ {
			for j = 0; j < size; j++ {
				grid[i][j]++;
			}
		}

		// output
		for i = 0; i < size; i++ {
			for j = 0; j < size; j++ {
				// flash our octopus if over energy level 9
				flash(pa, i, j)
				fmt.Printf("%d", grid[i][j])
			}
			fmt.Printf("\n")
		}
		
		fmt.Println("")
	}
	fmt.Printf("total flashes: %d\n", totalFlashes)
	
}