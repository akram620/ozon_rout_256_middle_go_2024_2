package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main6() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	if n > m || n == m {
		fmt.Fprintln(out, -1)
		return
	}

	origPlaces := make(map[int][]int, n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		origPlaces[a[i]] = append(origPlaces[a[i]], i)
	}

	res := make([]int, n)
	sort.Ints(a)

	curCounter := a[0]
	for i := 0; i < n; i++ {
		el := a[i] + 1
		if el > curCounter {
			curCounter = el
		} else {
			curCounter++
		}

		if curCounter > m {
			fmt.Fprintln(out, -1)
			return
		}

		fmt.Fprint(out, curCounter, " ")

		res[origPlaces[a[i]][0]] = curCounter
		origPlaces[a[i]] = origPlaces[a[i]][1:]

	}
	fmt.Fprintln(out)
}
