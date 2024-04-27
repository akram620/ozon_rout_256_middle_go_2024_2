package _

import (
	"bufio"
	"fmt"
	"os"
)

func main3() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)

	var c = 0
	var lastGlobal = 0
	mes := make(map[int]int)

	for i := 0; i < q; i++ {
		var t, id int
		fmt.Fscan(in, &t, &id)

		if t == 1 {
			c++

			if id == 0 {
				mes = make(map[int]int)
				lastGlobal = c
				continue
			}

			mes[id] = c
			continue
		}

		if t == 2 {
			cc, ok := mes[id]
			if ok {
				fmt.Fprintln(out, cc)
				continue
			}
			fmt.Fprintln(out, lastGlobal)
		}
	}
}
