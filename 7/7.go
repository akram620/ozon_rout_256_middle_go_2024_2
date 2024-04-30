package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main7() {
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for j := 0; j < t; j++ {
		var n, m int
		fmt.Fscan(in, &n, &m)

		origPlaces := make(map[int][]int, m)

		var a = make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &a[i])
			origPlaces[a[i]] = append(origPlaces[a[i]], i)
		}

		sort.Ints(a)

		used := make(map[int]struct{})

		var err bool

		res := make([]rune, m)

		findPlace := func(el int, r rune) {
			res[origPlaces[el][0]] = r
			origPlaces[el] = origPlaces[el][1:]
		}

		for _, el := range a {
			if _, ok := used[el-1]; !ok && el > 1 {
				used[el-1] = struct{}{}
				findPlace(el, '-')
				continue
			}

			if _, ok := used[el]; !ok {
				used[el] = struct{}{}
				findPlace(el, '0')
				continue
			}

			if _, ok := used[el+1]; !ok && el < n {
				used[el+1] = struct{}{}
				findPlace(el, '+')
				continue
			}

			err = true
			break
		}

		if err {
			fmt.Fprintln(out, "x")
		} else {
			for _, r := range res {
				fmt.Fprint(out, string(r))
			}
			fmt.Fprintln(out)
		}
	}
}
