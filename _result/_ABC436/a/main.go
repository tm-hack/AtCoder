package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {
	N := nextInt()
	S := nextString()

	prefix := strings.Repeat("o", N-len(S))
	result := prefix + S

	fmt.Println(result)
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
