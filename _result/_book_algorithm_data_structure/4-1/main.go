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
	out(tribo(N))
}

func tribo(N int) int {
	// ベースケース
	if N == 0 {
		return 0
	} else if N == 1 {
		return 0
	} else if N == 2 {
		return 1
	}

	// 再帰呼び出し
	return tribo(N-1) + tribo(N-2) + tribo(N-3)
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
