package _

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var start, end int
		var r string
		fmt.Fscan(in, &start, &end, &r)

		s = s[:start-1] + r + s[end:]
	}

	fmt.Fprintln(out, s)
}
