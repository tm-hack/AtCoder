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
