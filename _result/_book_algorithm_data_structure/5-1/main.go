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
	N := nextInt()
	v := matrix2D(3, N, 0)
	for i := 0; i < N; i++ {
		v[0][i] = nextInt()
		v[1][i] = nextInt()
		v[2][i] = nextInt()
	}

	dp := matrix2D(3, N+1, 0)
	for i := 1; i < N+1; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j != k {
					chmax(&dp[j][i], dp[k][i-1]+v[k][i-1])
				}
			}
		}
	}

	res := 0
	for i := 0; i < 3; i++ {
		chmax(&res, dp[i][N])
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
func matrix2D(h, w, val int) [][]int {
	res := make([][]int, h)
	for i := 0; i < h; i++ {
		res[i] = make([]int, w)
		if val != 0 {
			for j := 0; j < w; j++ {
				res[i][j] = val
			}
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
