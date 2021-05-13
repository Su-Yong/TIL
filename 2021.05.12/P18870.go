package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var length int
	fmt.Fscan(reader, &length)

	array := make([]int, length)
	sorted := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(reader, &array[i])
		sorted[i] = array[i]
	}

	sort.Sort(sort.IntSlice(sorted))

	check := -1
	table := make(map[int]int)
	i := 0
	for _, v := range sorted {
		if check != v {
			check = v
			table[v] = i
			i += 1
		}
	}

	var result strings.Builder
	for _, v := range array {
		result.WriteString(strconv.Itoa(table[v]))
		result.WriteString(" ")
	}
	writer.WriteString(result.String())
	writer.Flush()
}
