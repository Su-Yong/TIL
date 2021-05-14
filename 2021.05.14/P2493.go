package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var length int
	fmt.Fscan(reader, &length)

	array := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(reader, &array[i])
	}

	// start

	result := make([]int, length)
	point := make([]int, 1)
	point[0] = 0
	last := 0
	for i, v := range array {
		if last > v {
			point = append(point, i)
		}

		last = v
	}

	lastPointIndex := len(point) - 1
	for i := len(array) - 1; i >= 0; i-- {
		if point[lastPointIndex] > i {
			lastPointIndex -= 1
		}

		for j := lastPointIndex; j > 0; j-- {
			if array[i] < array[point[j]-1] {
				result[i] = point[j]
				break
			}
		}
	}

	// end

	for _, v := range result {
		writer.WriteString(strconv.Itoa(v))
		writer.WriteString(" ")
	}

	writer.Flush()
}
