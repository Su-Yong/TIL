package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var K, N int
	fmt.Fscan(reader, &N)
	fmt.Fscan(reader, &K)

	sum := (K + K*K) / 2
	if sum > N {
		fmt.Println("-1")
	} else if (N-K-sum)%K == 0 {
		fmt.Println(K - 1)
	} else {
		fmt.Println(K)
	}
}
