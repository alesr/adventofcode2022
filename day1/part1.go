package adventofcode22

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(inputFile string) (int, error) {
	readFile, err := os.Open(inputFile)
	if err != nil {
		return 0, fmt.Errorf("could not open input file: %s", err)
	}
	defer readFile.Close()

	var (
		highestCalories int
		calories        int
	)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		inputLine := strings.TrimSpace(fileScanner.Text())

		if inputLine == "" {
			calories = 0
			continue
		}

		lineValue, err := strconv.Atoi(inputLine)
		if err != nil {
			return 0, fmt.Errorf("could not convert line input to int: %s", err)
		}

		calories += lineValue

		if calories > highestCalories {
			highestCalories = calories
		}
	}
	return highestCalories, nil
}
