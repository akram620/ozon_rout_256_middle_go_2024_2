package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type directory struct {
	Dir     string      `json:"dir"`
	Files   []string    `json:"files"`
	Folders []directory `json:"folders"`
}

func main5() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	res := make([]int, t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		in.ReadString('\n')
		rows := make([]string, n)
		for j := 0; j < n; j++ {
			rows[j], _ = in.ReadString('\n')
		}
		rawSting := strings.Join(rows, "")

		// unmarshal json
		var dir directory
		json.Unmarshal([]byte(rawSting), &dir)

		c1 := checkFiles(&dir)
		res[i] = c1
	}

	for _, r := range res {
		fmt.Println(r)
	}
}

func checkFiles(dic *directory) int {
	var count int
	for _, f := range dic.Files {
		if strings.HasSuffix(f, ".hack") {
			return fileCount(dic)
		}
	}

	for _, f := range dic.Folders {
		count += checkFiles(&f)
	}

	return count
}

func fileCount(dic *directory) int {
	count := len(dic.Files)
	for _, f := range dic.Folders {
		count += fileCount(&f)
	}
	return count
}
