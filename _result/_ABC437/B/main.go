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
	H := nextInt()
	W := nextInt()
	N := nextInt()
	A := matrix2D(H, W, 0)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			A[i][j] = nextInt()
		}
	}
	B := make([]int, N)
	for i := 0; i < N; i++ {
		B[i] = nextInt()
	}

	res := 0
	tmp := 0
	for i := 0; i < H; i++ {
		tmp = 0
		for j := 0; j < N; j++ {
			for k := 0; k < W; k++ {
				if B[j] == A[i][k] {
					tmp = tmp + 1
				}
			}
		}
		chmax(&res, tmp)
	}

	out(res)
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
