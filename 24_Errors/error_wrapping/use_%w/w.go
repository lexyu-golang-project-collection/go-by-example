package main

import (
	"errors"
	"fmt"
)

var ErrSentinel = errors.New("the underlying sentinel error")

func main() {
	// 使用 %w 包裝錯誤，保留原始錯誤
	err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
	err2 := fmt.Errorf("wrap err1: %w", err1)

	println(err2 == ErrSentinel) // false，直接比較不相等

	// 使用 errors.Is 檢查包裝鏈中是否包含特定錯誤
	if errors.Is(err2, ErrSentinel) {
		println("err2 is ErrSentinel")
		return
	}
	println("err2 is not ErrSentinel")
}
