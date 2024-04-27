package _

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main4() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)

		origPlates := make(map[int][]int)

		var times = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &times[j])
			origPlates[times[j]] = append(origPlates[times[j]], j)
		}

		sort.Ints(times)

		place := 1
		curPlace := 1

		sortedPlates := make(map[int]int)
		sortedPlates[times[0]] = place

		for j := 1; j < n; j++ {
			if times[j] == times[j-1] || times[j]-times[j-1] == 1 || times[j]-times[j-1] == -1 {
				sortedPlates[times[j]] = place
				curPlace++
				continue
			}

			curPlace++
			place = curPlace

			sortedPlates[times[j]] = curPlace
		}

		res := make([]int, n)
		for tt, pl := range sortedPlates {
			for _, pos := range origPlates[tt] {
				res[pos] = pl
			}
		}

		for _, r := range res {
			fmt.Fprint(out, r, " ")
		}
		fmt.Fprintln(out)
	}
}
