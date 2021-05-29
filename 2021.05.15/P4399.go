package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	// writer := bufio.NewWriter(os.Stdout)

	array := make([]string, 0)
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}

	// start

	fmt.Println("array", array)

	// end
}
