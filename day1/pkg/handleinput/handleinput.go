package handleinput

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func ReadInput(filename string) (col1, col2 []int) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		location, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		col1 = append(col1, location)

		scanner.Scan()
		location, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		col2 = append(col2, location)
	}

	sort.Ints(col1)
	sort.Ints(col2)

	return col1, col2
}
