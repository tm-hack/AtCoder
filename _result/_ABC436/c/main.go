package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

type Point struct {
	X, Y int
}

func main() {
	_ = nextInt()
	M := nextInt()
	placed := make(map[Point]bool)
	ans := 0

	for i := 0; i < M; i++ {
		Ri := nextInt()
		Ci := nextInt()
		ok := true

		for x := Ri - 1; x <= Ri+1; x++ {
			for y := Ci - 1; y <= Ci+1; y++ {
				if placed[Point{x, y}] {
					ok = false
				}
			}
		}

		if ok {
			placed[Point{Ri, Ci}] = true
			ans++
		}
	}
	fmt.Println(ans)
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
