package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var length int
	fmt.Fscanln(reader, &length)

	var table1, table2 map[rune]int
	var str1, str2 string
	for i := 0; i < length; i += 1 {
		table1 = make(map[rune]int)
		table2 = make(map[rune]int)
		str1 = ""
		str2 = ""
		fmt.Fscanln(reader, &str1)
		fmt.Fscanln(reader, &str2)

		result := len(str1) + len(str2)

		for _, s1 := range str1 {
			table1[s1] += 1
		}

		for _, s2 := range str2 {
			table2[s2] += 1
		}

		for _, key := range "abcdefghijklmnopqrstuvwxyz" {
			if table1[key] < table2[key] {
				result -= table1[key] * 2
			} else {
				result -= table2[key] * 2
			}
		}

		fmt.Fprint(writer, "Case #", i+1, ": ", result, "\n")
		writer.Flush()
	}
}
