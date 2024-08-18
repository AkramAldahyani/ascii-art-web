package ascii

import (
	"bufio"
	"fmt"
	"os"
)

func Read(fileName string) map[int][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	letters := make(map[int][]string)
	row := 0
	line := 1
	scanner := bufio.NewScanner(file)
	list := make([]string, 8)
	for scanner.Scan() {
		if line == 1 {
			line++
			continue
		}
		if (line-1)%9 == 0 {
			letters[(line-1)/9-1+32] = list
			list = make([]string, 8)
			line++
			row = 0
		} else {
			list[row] = scanner.Text()
			line++
			row++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	return letters
}
