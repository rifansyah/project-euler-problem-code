// Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down,
// there are exactly 6 routes to the bottom right corner.
// How many such routes are there through a 20×20 grid?

package main

import "fmt"

func main() {
	const length = 21
	const width = 21

	var grid [length][width]int

	//initiate 1 way
	for i := 0; i < length; i++ {
		grid[0][i] = 1
	}

	for i := 0; i < width; i++ {
		grid[i][0] = 1
	}

	//way from top left to bottom right
	for i := 1; i < width; i++ {
		for j := 1; j < length; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	fmt.Println(grid[length-1][width-1], "ways")
}
