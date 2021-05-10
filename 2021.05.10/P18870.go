package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Entry struct {
	value, index int
}

type EntryArray []Entry

func (entry EntryArray) Len() int {
	return len(entry)
}
func (entry EntryArray) Swap(i, j int) {
	entry[i].index, entry[j].index = entry[j].index, entry[i].index
	entry[i].value, entry[j].value = entry[j].value, entry[i].value
}
func (entry EntryArray) Less(i, j int) bool {
	return entry[i].value < entry[j].value
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var length, input int
	var result [1000000]int
	var origin []Entry

	fmt.Fscan(reader, &length)
	for i := 0; i < length; i++ {
		fmt.Fscan(reader, &input)

		origin = append(origin, Entry{input, i})
	}

	sort.Sort(EntryArray(origin))

	now := origin[0].value
	value := 0
	for _, v := range origin {
		if v.value > now {
			value += 1
			now = v.value
		}
		result[v.index] = value
	}

	for i := 0; i < length; i++ {
		fmt.Print(result[i], " ")
	}
}
