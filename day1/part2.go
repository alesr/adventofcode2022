package adventofcode22

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part2(inputFile string) (int, error) {
	readFile, err := os.Open(inputFile)
	if err != nil {
		return 0, fmt.Errorf("could not open input file: %s", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var i int
	elfSupplies := make(map[int]int)

	for fileScanner.Scan() {
		inputLine := strings.TrimSpace(fileScanner.Text())

		if inputLine == "" {
			i++
			continue
		}

		calories, err := strconv.Atoi(inputLine)
		if err != nil {
			return 0, fmt.Errorf("could not convert line input to int: %s", err)
		}
		elfSupplies[i] += calories
	}

	calories := make([]int, len(elfSupplies))
	for _, v := range elfSupplies {
		calories = append(calories, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	var top3CaloriesSum int
	for i := 0; i < 3; i++ {
		top3CaloriesSum += calories[i]
	}

	return top3CaloriesSum, nil
}
