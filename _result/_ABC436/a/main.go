package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//入力の数が決まっており、半角スペース区切りの場合
	//// 1文字ずつデータ型を指定して受け取る
	var a, b, c int                   // int型の変数を宣言
	var s string                      // string型の変数を宣言
	fmt.Scanf("%d %d %d", &a, &b, &c) // %dでint型を代入
	fmt.Scanf("%s", &s)               // %sでstring型を代入

	// 入力の数が決まっており、区切り文字が無い場合
	//// 1文字ずつデータ型を指定して受け取る
	var d, e, f int                    // int型の変数を宣言
	fmt.Scanf("%1d%1d%1d", &d, &e, &f) // %1dでint型を1つずつ代入

	// 入力の数が決まっておらず、半角スペース区切りの場合
	sc := bufio.NewScanner(os.Stdin)        // 標準入力を受け付けるスキャナ
	sc.Scan()                               // １行分の入力を取得する
	inputs := strings.Split(sc.Text(), " ") // 半角スペース区切りでstring型として配列inputsに格納

	//// 配列inputsの中身をstring→intに変換してリストに格納
	var ps []int                   // Pnを格納する配列を宣言
	for _, input := range inputs { // 配列inputsの全ての要素について実行
		p, _ := strconv.Atoi(input) // string→intに型変換
		ps = append(ps, p)          // intに型変換した値を、Pnを格納する配列に追加
	}

	// 入力の数が決まっておらず、複数行に分かれている場合
	//// 1文字ずつデータ型を指定して受け取る
	var n int           // int型の変数を宣言
	fmt.Scanf("%d", &n) // %dでint型を代入

	sc := bufio.NewScanner(os.Stdin) // 標準入力を受け付けるスキャナ

	var us []int // uの値を格納する配列を宣言
	var vs []int // vの値を格納する配列を宣言

	for i := 0; i < n; i++ { // 行数分繰り返す
		sc.Scan()                               // １行分の入力を取得する
		inputs := strings.Split(sc.Text(), " ") // 半角スペース区切りでstring型として配列inputsに格納
		u, _ := strconv.Atoi(inputs[0])         // string→intに型変換
		v, _ := strconv.Atoi(inputs[1])         // string→intに型変換
		us = append(us, u)                      // uの配列に追加
		vs = append(vs, v)                      // vの配列に追加
	}
}
