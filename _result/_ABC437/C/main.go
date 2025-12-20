package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {
	defer flush()
	T := nextInt()

	for i := 0; i < T; i++ {
		N := nextInt()
		W := make([]int, N)
		P := make([]int, N)
		WP := make([]int, N)

		sumP := 0
		for j := 0; j < N; j++ {
			W[j] = nextInt()
			P[j] = nextInt()
			WP[j] = W[j] + P[j]
			sumP = sumP + P[j]
		}
		sort.Ints(WP)

		res := 0
		for j := 0; j < N; j++ {
			if sumP-WP[j] >= 0 {
				res = res + 1
				sumP = sumP - WP[j]
			}
		}

		out(res)
	}

}

// ==================================================
// init
// ==================================================
func init() {
	sc.Split(bufio.ScanWords)
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
// binary_search
// ==================================================
func binary_search[T cmp.Ordered](key T, a []T) int {
	left := 0
	right := len(a) - 1
	if right >= left {
		mid := left + (right-left)/2
		if a[mid] == key {
			return mid
		} else if a[mid] > key {
			right = mid - 1
		} else if a[mid] < key {
			left = mid + 1
		}
	}
	return -1
}
