package lines

import (
	"bufio"
	"os"
)

func Count(file *os.File) (int, error) {
	var count int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count++
	}

	err := scanner.Err()
	if err != nil {
		return count, err
	}

	return count, nil
}
