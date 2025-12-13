package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	// 入力読み込み
	N := nextInt()

	// 配列の計算
	grid := make2D(N, N, -1)

	r := 0
	c := (N - 1) / 2
	grid[r][c] = 1
	for k := 1; k < N*N; k++ {
		if grid[returnMod(r-1, N)][returnMod(c+1, N)] == -1 {
			r = returnMod(r-1, N)
			c = returnMod(c+1, N)
		} else {
			r = returnMod(r+1, N)
			c = c
		}
		grid[r][c] = k + 1
	}

	// 解の出力
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, row := range grid {
		for _, val := range row {
			fmt.Fprintf(w, "%d ", val)
		}
		fmt.Fprintln(w)
	}
}

func make2D(h, w, val int) [][]int {
	res := make([][]int, h)
	for i := 0; i < h; i++ {
		res[i] = make([]int, w)
		if val != 0 { // 0以外で初期化したい場合
			for j := 0; j < w; j++ {
				res[i][j] = val
			}
		}
	}
	return res
}

func returnMod(a int, b int) int {
	ans := (a%b + b) % b
	return ans
}
