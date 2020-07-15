package main

import (
	"fmt"
	"log"
	"sync"
)

/*
2. chạy đoạn chương trình dưới đây. Nếu có lỗi hãy thêm logic để nó chạy đúng.
- Lý giải nguyên nhân lỗi.

```go
func errFunc() {
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 10000; j++ {
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10
			}
		}()
	}

	log.Print("done")
}

```
*/

func errFunc() {
	m := make(map[int]int)
	//Nguyên nhân lỗi do 1000 cái goroutine cùng truy cập và gán giá trí cho map => race condition
      var mu sync.Mutex
	//Tao 1000 cai goroutine
	for i := 0; i < 1000; i++ {
        
		go func() {
            
			//Moi cai goroutine tao 9999 phan tu map
			for j := 1; j < 10000; j++ {
				//Neu m[j] != nil thi xoa
				mu.Lock()
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10
				mu.Unlock()
			}
			
		}()
	
	}
    
	log.Print("done")
	fmt.Println(len(m))

}
