package main

import (
	"fmt"
)

/*
 * https://gist.github.com/bsoo/bd361a7fe01fc6bff068
 * 指定した値までの素数を算出する
 */

func main() {
	// 素数の値を求める数を設定
	targetSize := 100
	// 指定した数字になるまで回し続ける
	for i := 1; i <= targetSize; i++ {
		count := 0
		for j := 1; j <= targetSize; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Print(i)
			fmt.Print("　")
		}

	}
	fmt.Println("")
}
