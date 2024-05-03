//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//var (
//	in  = bufio.NewReader(os.Stdin)
//	out = bufio.NewWriter(os.Stdout)
//)
//
//func main() {
//	defer out.Flush()
//
//	var t int
//	fmt.Fscan(in, &t)
//
//	for j := 0; j < t; j++ {
//		var n, b, r, f int
//		fmt.Fscan(in, &n, &b, &r, &f)
//
//		blueList := make([]string, 0, b)
//		redList := make([]string, 0, r)
//
//		forbidden := make(map[string]struct{})
//
//		cacheB := make(map[string]uint)
//		cacheR := make(map[string]uint)
//
//		for i := 0; i < n; i++ {
//			var w string
//			fmt.Fscan(in, &w)
//
//			forbidden[w] = struct{}{}
//
//			if i == f-1 {
//				for k1 := 0; k1 < len(w); k1++ {
//					for k2 := k1 + 1; k2 < len(w)+1; k2++ {
//						forbidden[w[k1:k2]] = struct{}{}
//					}
//				}
//			}
//
//			if i < b {
//				blueList = append(blueList, w)
//			}
//
//			if i < b+r && i >= b {
//				redList = append(redList, w)
//			}
//		}
//
//		var resWord string
//		var resCount uint
//
//		for _, w := range blueList {
//			for k1 := 0; k1 < len(w); k1++ {
//				for k2 := k1 + 1; k2 < len(w)+1; k2++ {
//					str := w[k1:k2]
//					if _, ok := forbidden[str]; ok {
//						continue
//					}
//
//					inB := getCountInBlueList(&blueList, &cacheB, str)
//					inR := getCountInRedList(&redList, &cacheR, str)
//
//					cacheB[str] = inB
//					cacheR[str] = inR
//
//					if inB > inR {
//						dif := inB - inR
//						if dif == resCount && len(str) > len(resWord) {
//							resWord = str
//							continue
//						}
//
//						if dif > resCount {
//							resCount = dif
//							var prev = resWord
//
//							resWord = str
//
//							if len(str) >= len(resWord) {
//								resWord = str
//							} else {
//								resWord = prev
//							}
//						}
//					}
//				}
//			}
//		}
//
//		if resCount == 0 {
//			fmt.Fprintln(out, "tkhapjiabb", 0)
//			continue
//		}
//
//		fmt.Fprintln(out, resWord, resCount)
//	}
//}
//
//func getCountInBlueList(blueList *[]string, c *map[string]uint, str string) uint {
//	if val, ok := (*c)[str]; ok {
//		return val
//	}
//	count := uint(0)
//	for _, w := range *blueList {
//		if w == str {
//			continue
//		}
//		ok := strings.Contains(w, str)
//		if ok {
//			count++
//		}
//	}
//
//	return count
//}
//
//func getCountInRedList(redList *[]string, c *map[string]uint, str string) uint {
//	if val, ok := (*c)[str]; ok {
//		return val
//	}
//
//	count := uint(0)
//	for _, w := range *redList {
//		if w == str {
//			continue
//		}
//		ok := strings.Contains(w, str)
//		if ok {
//			count++
//		}
//	}
//
//	return count
//}

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
	for i := 0; i < t; i++ {
		processTest()
	}
}

func processTest() {
	var n, b, r, f int
	fmt.Fscan(in, &n, &b, &r, &f)

	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &words[i])
	}
	wordsSet := arrToSet(words)
	blueSubWordsMap := getSubWordsMap(words[:b])
	blackWord := words[f-1]
	blackSubWords := subWordsMap(blackWord)
	withoutRight(blueSubWordsMap, blackSubWords)
	withoutRight(blueSubWordsMap, wordsSet)
	redSubWordsMap := getSubWordsMap(words[b : b+r])
	withoutRight(redSubWordsMap, blackSubWords)
	withoutRight(redSubWordsMap, wordsSet)
	subWord, bestBR := bestResult(blueSubWordsMap, redSubWordsMap)
	fmt.Fprintln(out, subWord, bestBR)
}

func bestResult(b map[string]int, r map[string]int) (string, int) {
	subWord := "tkhapjiabb"
	var bestBR int
	for k, v := range b {
		newBestBR := v - r[k]
		if newBestBR > bestBR {
			subWord = k
			bestBR = newBestBR
		}
	}
	return subWord, bestBR
}

func subWordsMap(word string) map[string]bool {
	n := len(word)
	rez := make(map[string]bool, n*(n+1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			rez[word[i:j]] = true
		}
	}
	return rez
}

func arrToSet(words []string) map[string]bool {
	rez := make(map[string]bool, len(words))
	for _, word := range words {
		rez[word] = true
	}
	return rez
}

func getSubWordsMap(words []string) map[string]int {
	rez := make(map[string]int)
	for _, w := range words {
		subW := subWordsMap(w)
		for k := range subW {
			if _, ok := rez[k]; ok {
				rez[k]++
			} else {
				rez[k] = 1
			}
		}
	}
	return rez
}

func withoutRight[K comparable, V1 any, V2 any](mapL map[K]V1, mapR map[K]V2) {
	for k := range mapL {
		if _, ok := mapR[k]; ok {
			delete(mapL, k)
		}
	}
}
