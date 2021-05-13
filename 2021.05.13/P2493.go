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

	point := make([]int, 1)
	last := array[0]
	point[0] = 0
	for i, v := range array {
		if last > v {
			point = append(point, i)
		}

		last = v
	}
	if point[len(point)-1] < len(array)-1 {
		point = append(point, len(array)-1)
	}

	fmt.Println(point)

	result := make([]string, 0)
	check := false
	lastIndex := len(point) - 1
	for i := len(array) - 1; i >= 0; i-- {
		check = false

		fmt.Println("now", i, array[i], point[lastIndex])

		for k := point[lastIndex]; k >= i; k-- {
			fmt.Println("find", array[k], array[i])
			if array[k] >= array[i] {
				result = append(result, strconv.Itoa(k+1))
				check = true
				break
			}
		}

		if lastIndex > 0 && lastIndex > i {
			lastIndex -= 1
		}

		if !check {
			result = append(result, "0")
		}
	}

	for i := len(result) - 1; i >= 0; i-- {
		writer.WriteString(result[i])
		writer.WriteString(" ")
	}

	writer.Flush()
}
