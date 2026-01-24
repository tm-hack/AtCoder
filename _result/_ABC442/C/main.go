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
	M := nextInt()
	conflict_map := make([][]int, N+1)
	for i := 0; i < M; i++ {
		A := nextInt()
		B := nextInt()
		conflict_map[A] = append(conflict_map[A], B)
		conflict_map[B] = append(conflict_map[B], A)
	}

	res := make([]int, 0)
	for i := 1; i <= N; i++ {
		res = append(res, combination(N-len(conflict_map[i])-1, 3))
	}

	outlist(res)
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

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}

func homogeneous(n int, k int) int {
	return combination(n+k-1, k)
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
