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

	result := ""
outer:
	for i, v := range array {
		for j := i - 1; j > 0; j-- {
			if v <= array[j] {
				result += strconv.Itoa(j+1) + " "
				continue outer
			}
		}
		result += "0 "
	}

	writer.WriteString(result)
	writer.Flush()
}
