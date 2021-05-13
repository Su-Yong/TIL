package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var result strings.Builder
	check := false
	for i, v := range array {
		check = false
		for j := i - 1; j >= 0; j-- {
			if v <= array[j] {
				result.WriteString(strconv.Itoa(j + 1))
				result.WriteString(" ")
				check = true
				break
			}
		}

		if !check {
			result.WriteString("0 ")
		}
	}

	writer.WriteString(result.String())
	writer.Flush()
}
