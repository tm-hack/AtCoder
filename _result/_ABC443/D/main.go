package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {
	defer flush()
	T := nextInt()
	for i := 0; i < T; i++ {
		N := nextInt()
		R := make([]int, N)
		pointed := make([][]int, N)
		for i := 0; i < N; i++ {
			R[i] = nextInt() - 1
			pointed[R[i]] = append(pointed[R[i]], i)
		}

		out(pointed)
		out(R)

		res := 0
		flag := make([]bool, N)
		for i := 0; i < N; i++ {
			for j := 0; j < len(pointed[i]); j++ {
				if !flag[pointed[i][j]] {
					flag[pointed[i][j]] = true

					if pointed[i][j] > 0 {
						res = R[pointed[i][j]-1] - i + 1
						flag[pointed[i][j]-1] = true
					}

					if pointed[i][j] < N-1 {
						res = R[pointed[i][j]+1] - i + 1
						flag[pointed[i][j]+1] = true
					}

					res = pointed[i][j]
				}
			}
		}
	}

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

func mod(a int, b int) int {
	ans := (a%b + b) % b
	return ans
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sqrt(a int) int {
	if a < 0 {
		return -1
	}

	res := int(math.Sqrt(float64(a)))
	if res*res == a {
		return res
	}

	return -1
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

func itoa(i int) string {
	return strconv.Itoa(i)
}

func atoi(a string) int {
	i, _ := strconv.Atoi(a)
	return i
}

func out(v ...interface{}) {
	fmt.Fprintln(wtr, v...)
}

func outlist(sl []int) {
	r := make([]string, len(sl))
	for i, v := range sl {
		r[i] = itoa(v)
	}
	out(strings.Join(r, " "))
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
