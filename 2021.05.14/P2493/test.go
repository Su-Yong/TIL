package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var name string
	fmt.Scan(&name)

	file, err := os.Open("./" + name + ".txt")
	output, outputErr := os.Create("./" + name + ".p2493.txt")

	if err != nil || outputErr != nil {
		fmt.Println("cannot find file \"" + name + "\"")
		return
	}

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(output)

	timestamp("input")

	var length int
	fmt.Fscan(reader, &length)

	array := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(reader, &array[i])
	}

	timestamp("start")

	// start

	result := make([]int, length)
	point := make([]int, 1)
	point[0] = 0
	last := -1
	for i, v := range array {
		if last > v {
			point = append(point, i)
		}

		last = v
	}

	timestamp("point end")

	lastPointIndex := len(point) - 1
	for i := len(array) - 1; i >= 0; i-- {
		if point[lastPointIndex] > i {
			lastPointIndex -= 1
		}

		for j := lastPointIndex; j > 0; j-- {
			fmt.Println("point check", array[i], array[point[j]], point[j])
			if array[i] < array[point[j]-1] {
				result[i] = point[j]
				break
			}
		}
	}

	timestamp("end")

	// end

	for _, v := range result {
		writer.WriteString(strconv.Itoa(v))
		writer.WriteString(" ")
	}

	writer.Flush()
}

func timestamp(str ...interface{}) {
	fmt.Println(str, "\n\ttime:", time.Now())
}
