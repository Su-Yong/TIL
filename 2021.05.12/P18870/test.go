package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	fmt.Scan(&name)

	file, err := os.Open("./" + name + ".txt")
	output, outputErr := os.Create("./" + name + ".p18870.txt")

	if err != nil || outputErr != nil {
		fmt.Println("cannot find file \"" + name + "\"")
		return
	}

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(output)

	timestamp("start")

	var length int
	fmt.Fscan(reader, &length)
	timestamp("array length")

	array := make([]int, length)
	sorted := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(reader, &array[i])
		sorted[i] = array[i]
	}

	timestamp("save array")

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

	timestamp("sort array")

	var result strings.Builder
	for _, v := range array {
		result.WriteString(strconv.Itoa(table[v]))
		result.WriteString(" ")
	}

	timestamp("solved problem")

	writer.WriteString(result.String())
	writer.Flush()

	timestamp("end")
}

func timestamp(str ...interface{}) {
	fmt.Println(str, "\n\ttime:", time.Now())
}
