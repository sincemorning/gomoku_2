package main

import (
	"fmt"
)

// https://gist.github.com/bsoo/bd361a7fe01fc6bff068
// ひし形を出力
func main() {
	max_width := 7
	add_size := 2
	upper := true
	count := 1
	// 回し続けて最大lenになったら減らす
	for {
		// 挿入すべきインデントの桁数を算出した上で先に出力
		insertIndent := ((max_width - count) / 2)
		if insertIndent > 0 {
			for ii := 0; ii < insertIndent; ii++ {
				fmt.Print(" ")
			}
		}
		for i := 0; i < count; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
		// 出力すべき値になったら反転するためにフラグを設定
		if count >= max_width {
			upper = false
		}

		// trueの場合は正方向に出力、falseの場合は負の方向に出力する
		if upper {
			count += add_size
		} else {
			count -= add_size
			if count <= 0 {
				break
			}
		}

	}
}
