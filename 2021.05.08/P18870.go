package main

import "fmt"

func main() {
	var length, input int
	var origin []int
	table := make(map[int]int)

	fmt.Scan(&length)

	for i := 0; i < length; i++ {
		fmt.Scan(&input)
		origin = append(origin, input)

		_, check := table[input]

		if !check {
			table[input] = 0
		}

		if !include(input, origin, i) {
			value := 0
			for key := range table {
				if key > input {
					table[key] += 1
				} else if key < input {
					value += 1
				}
			}

			table[input] = value
		}
	}

	for _, v := range origin {
		fmt.Print(table[v], " ")
	}
}

func include(value int, array []int, except int) bool {
	for index, compare := range array {
		if compare == value && except != index {
			return true
		}
	}

	return false
}
