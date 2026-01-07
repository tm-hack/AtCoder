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
	N := nextInt()
	S := nextString()

	ans := 0
	for i := 0; i <= 999; i++ {
		key := fmt.Sprintf("%03d", i)
		idx := 0

		for j := 0; j < N; j++ {
			if S[j] == key[idx] {
				idx = idx + 1
			}

			if idx == 3 {
				ans = ans + 1
				break
			}
		}

	}

	out(ans)
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
