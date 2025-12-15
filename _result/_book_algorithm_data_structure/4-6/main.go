package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {
	defer flush()
	N := nextInt()
	W := nextInt()
	An := make([]int, N)
	for i := 0; i < N; i++ {
		An[i] = nextInt()
	}

	memo := make2D(W+1, N+1, -1)
	if dfs(N, W, An, memo) == 1 {
		out("True")
	} else {
		out("False")
	}
}

func dfs(i int, w int, An []int, memo [][]int) int {
	// basecase
	if i == 0 {
		if w == 0 {
			return 1
		} else {
			return 0
		}
	}

	if w < 0 {
		return 0
	}

	if memo[w][i] != -1 {
		return memo[w][i]
	}

	// recursion
	if dfs(i-1, w-An[i-1], An, memo) == 1 {
		memo[w][i] = 1
	} else if dfs(i-1, w, An, memo) == 1 {
		memo[w][i] = 1
	} else {
		memo[w][i] = 0
	}
	return memo[w][i]
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
func make2D(h, w, val int) [][]int {
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
