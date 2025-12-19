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

const INF = 1_000_000_000

func main() {
	defer flush()
	N := nextInt()
	W := nextInt()
	K := nextInt()
	a := make([]int, N+1)
	for i := 1; i <= N; i++ {
		a[i] = nextInt()
	}

	dp := matrix2D(N+1, W+1, INF)
	for i := 0; i < N+1; i++ {
		dp[i][0] = 0
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if j-a[i] >= 0 {
				chmin(&dp[i][j], dp[i-1][j-a[i]]+1)
				chmin(&dp[i][j], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	if dp[N][W] <= K {
		out("OK")
	} else {
		out("NG")
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
