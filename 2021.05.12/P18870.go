package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

	result := ""
	for _, v := range array {
		result += strconv.Itoa(table[v]) + " "
	}
	writer.WriteString(result)
	writer.Flush()
}
