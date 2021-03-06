package life106

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func newCells(width, height int) [][]bool {
	grid := make([][]bool, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]bool, width)
	}
	return grid
}

func ParseFile(path string) ([][]bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return [][]bool{}, err
	}
	return Parse(f), nil
}

func Parse(r io.Reader) [][]bool {
	br := bufio.NewReader(r)

	var list [][2]int
	minX, minY := 999, 999
	maxX, maxY := -999, -999
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		if line[0] == '#' {
			continue
		}

		var x, y int
		_, err = fmt.Sscanf(string(line), "%d %d", &x, &y)
		if err != nil {
			panic(err)
		}
		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}

		list = append(list, [2]int{x, y})
	}

	cells := newCells(maxX-minX+1, maxY-minY+1)
	for _, pos := range list {
		x, y := pos[0], pos[1]
		cells[y-minY][x-minX] = true
	}

	return cells
}
