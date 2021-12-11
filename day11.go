package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"bufio"
	"log"
)

var size int
var totalFlashes int

type Matrix [10][10]int

func flash(grid *Matrix, i int, j int) {

	if grid[i][j] > 9 {
		totalFlashes++
		grid[i][j] = 0 // reset ourself

		if i-1 > 0 && j-1 > 0 { 
			grid[i-1][j-1]++ 
			flash(grid,i-1,j-1)
		}
		
		if i - 1 > 0 {
			grid[i-1][j]++
			flash(grid,i-1,j)
		}
		
		if i - 1 > 0 && j + 1 < size {
			grid[i-1][j+1]++
			flash(grid,i-1,j+1)
		}

		if j - 1 > 0 {
			grid[i][j-1]++
			flash(grid,i,j-1)
		}
		
		if j + 1 < size {
			grid[i][j+1]++
			flash(grid,i,j+1)
		}

		if i + 1 < size && j - 1 > 0 {
			grid[i+1][j-1]++
			flash(grid,i+1,j-1)
		}
		
		if i + 1 < size {
			grid[i+1][j]++
			flash(grid,i+1,j)
		}

		if i + 1 < size && j + 1 < size {
			grid[i+1][j+1]++
			flash(grid,i+1,j+1)
		}
		
	}
}

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
	numTimesToRun := 100
	totalFlashes = 0
	var grid Matrix

	// initialize
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for x, line := range lines {
		for y, num := range line {
			grid[x][y] = int(num)
		}
	}

	size := len(grid[0])

	var pa *Matrix = &grid

	for x := 0; x < numTimesToRun; x++ {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		// Increase all by +1
		var i, j int
		for i = 0; i < size; i++ {
			for j = 0; j < size; j++ {
				grid[i][j]++;
			}
		}

		// Second loop
		for i = 0; i < size; i++ {
			for j = 0; j < size; j++ {
				flash(pa, i, j)
			}
		}

		// Third loop to print it out
		for i = 0; i < size; i++ {
			for j = 0; j < size; j++ {
				fmt.Printf("%d", grid[i][j])
			}
			fmt.Printf("\n")
		}
		
		fmt.Println("")
		time.Sleep(time.Second / 5)
	}
	fmt.Printf("total flashes: %d\n", totalFlashes)
	
}