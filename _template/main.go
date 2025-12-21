package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {
	defer flush()

}

// ==================================================
// init
// ==================================================
func init() {
	sc.Split(bufio.ScanWords)
}

// ==================================================
// num
// ==================================================
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// ==================================================
// io
// ==================================================
func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func out(v ...interface{}) {
	fmt.Fprintln(wtr, v...)
}

func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}

// ==================================================
// matrix
// ==================================================
func matrix2D[T any](h int, w int, val T) [][]T {
	res := make([][]T, h)
	for i := 0; i < h; i++ {
		res[i] = make([]T, w)
		for j := 0; j < w; j++ {
			res[i][j] = val
		}
	}
	return res
}

// ==================================================
// dp
// ==================================================
func chmin[T cmp.Ordered](a *T, b T) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func chmax[T cmp.Ordered](a *T, b T) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

// ==================================================
// binarySearch
// ==================================================
func binarySearch(ok, ng int, f func(int) bool) int {
	if !f(ok) {
		return -1
	}
	if f(ng) {
		return ng
	}
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if f(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

func lowerBound(v int, sl []int) int {
	if len(sl) == 0 {
		return -1
	}
	idx := binarySearch(0, len(sl)-1, func(c int) bool {
		return sl[c] <= v
	})
	return idx
}

func upperBound(v int, sl []int) int {
	if len(sl) == 0 {
		return 0
	}
	idx := binarySearch(0, len(sl)-1, func(c int) bool {
		return sl[c] <= v
	})
	return idx + 1
}
