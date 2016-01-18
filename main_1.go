package main

import (
	"fmt"
)

// https://gist.github.com/bsoo/bd361a7fe01fc6bff068
func main() {
	max_width := 7
	add_size := 2
	upper := true
	count := 1
	// 回し続けて最大lenになったら減らす
	for {
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
		// 必要な分を吐き切ったら、次の行のために設定情報を更新する
		if count >= max_width {
			upper = false
		}

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
