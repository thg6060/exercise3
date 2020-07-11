package main

import (
	"log"
	"sync"
	"time"
)

/*
1. dùng kiến thức về go routine và chan đề func dưới in ra đủ 3 message.

```go
func chanRoutine() {
    log.Print("hello 1")
    go func() {
        time.Sleep(1 * time.Second)
        log.Print("hello 3")
    }
    log.Print("hello 2")
}

```

-- nâng cao. In ra các message theo thứ tự.
-- In ra message 3 trước message 2.
Sử dụng 3 cách để làm( gợi ý: sử dụng mutex, chan, waitGroup)
*/

func chanRoutine1() {

	var mu sync.Mutex

	log.Print("hello 1")
	go func() {

		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		mu.Unlock()

	}()
	mu.Lock()
	mu.Lock() //In 3 truoc 2

	log.Print("hello 2")
	//mu.Lock() //In theo thu tu

}

func chanRoutine2() {
	c := make(chan bool)

	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		c <- true

	}()
	//<-c //In ra 3 truoc in ra 2
	log.Print("hello 2")
	//<-c //In ra dung thu tu

}
func chanRoutine3(wg *sync.WaitGroup) {

	log.Print("hello 1")
	wg.Add(1)
	go func() {

		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		wg.Done()
	}()
	//wg.Wait() //In 3 truoc in 2 sau
	log.Print("hello 2")
	//wg.Wait() //In theo thu tu

}

func Run1() {
	//var wg sync.WaitGroup
	chanRoutine1()
	//chanRoutine2()
	//chanRoutine3(&wg)
}
