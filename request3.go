package main

import (
	"fmt"
	"sync"
)

/*
3. tạo 1 biến X `map[string]string` và 3 goroutine cùng thêm dữ liệu vào X.
Mỗi goroutine thêm 1000 key khác nhau.
Sao cho quá trình đủ 15 key không mất mát dữ liệu.

Lưu ý sử dụng mutex
*/

func goRoutine(sKey int, eKey int, m map[string]string, mu *sync.Mutex) {

	for j := sKey; j < eKey; j++ {

		mu.Lock()
		m[string(j)] = string(j)
		mu.Unlock()
	}

}

func Run3() {
	m := make(map[string]string)

	var mu sync.Mutex

	go goRoutine(0, 1000, m, &mu)
	go goRoutine(1000, 2000, m, &mu)
	go goRoutine(2000, 3000, m, &mu)

	fmt.Println(len(m))
}
