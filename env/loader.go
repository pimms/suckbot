package env

import (
	"bufio"
	"os"
)

func LoadMap(filename string) [MAX_SIZE][MAX_SIZE]bool {
	if filename == "" {
		return loadDefault()
	}

	lines, err := readLines(filename)
	if err != nil {
		return loadDefault()
	}

	var v [MAX_SIZE][MAX_SIZE]bool
	
	var nrows = min(len(lines), MAX_SIZE)
	for row:=0; row<nrows; row++ {
		var ncols = min(len(lines[row]), MAX_SIZE)
		for col:=0; col<ncols; col++ {
			if lines[row][col] != ' ' {
				v[row][col] = true
			}
		}
	}

	return v
}


// Simple 4 by 4 environment
func loadDefault() [MAX_SIZE][MAX_SIZE]bool {
	var v [MAX_SIZE][MAX_SIZE]bool

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			v[x][y] = true
		}
	}

	return v
}

// Return the contents of a file as a slice of strings
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

// Return "b" if "b < a", return "a"
// otherwise.
func min(a, b int) int {
	if b < a {
		return b
	}

	return a
}
