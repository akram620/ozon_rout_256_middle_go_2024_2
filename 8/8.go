package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for j := 0; j < t; j++ {
		var n, b, r, f int
		fmt.Fscan(in, &n, &b, &r, &f)
		for i := 0; i < n; i++ {

		}
	}
}
