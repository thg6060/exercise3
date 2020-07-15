package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

/*
4. bài tập worker pool: tạo bằng tay file dưới. `file.txt` sau đó đọc từng dòng file này nạp dữ liệu vào 1 buffer channel có size 10, Điều kiện đọc file từng dòng. Chỉ được sử dụng 3 go routine. Kết quả xử lý xong ỉn ra màn hình + từ `xong`
```txt
"z9nnHLy8V8"
"6AVcSrDUkB"
"DezRGPwtx7"
"eSmXGjCmTq"
"9rfCMntQA5"
"Trk6xppMuM"
"2sb8BPaUsp"
"6AAh6zVFNA"
"gsY8kAuKp8"
"FQgb8QEpxg"
"hEXnKUkYrp"
"tchiG2Tiv4"
"daMPMJWaM6"
"WbBMpX89Sz"
"YVnsveajtj"
"L9TA7FE5d9"
"xBjE7UNe98"
"q6bLPeVjYr"
"oBppTK62nT"
"GxUjEDYBdG"
"ZTEpXFStLo"
"4XkynbWFvU"
"WFmmUSWzDv"
"nit8qjmvZH"
"iT8BqzHdXo"
"7N7mz3qzn2"
"KfhMZsHABi"
"M4mKWrGgDn"
"qLEduDF7so"
"YhigrGfLJr"
"f82gk2mrxv"
"q7TPNZB3Bv"
"eWLL5Yg6sG"
"GyPqxrXiUg"
"86dGJYRzPN"
"EWYtAVfXnd"
"8dNcD3F8uS"
"NLRE6LKqCt"
"UbLD2DACiB"
"JeLHTTg8vw"
```
nâng cao. In ra số lượng goroutine đã khởi tạo.
hoàn thiện để tối ưu, thu hồi channel và goroutine đã cấp phát.

- Nâng cao 1: Tạo 1 struct `Line` có trường gồm có: `số dòng hiện tại`, `giá trị` của dòng đó.
In ra màn hình cú pháp `${line_number} giá trị là: ${data}`.
- Nâng cao 2: Khi kết thúc chương trình đã cho đóng những vòng lặp vô hạn của các goroutine lại. Viết chương trình đó.
Giợi ý sử dụng biến `make([]chan bool, n)`
*/

type Line struct {
	Id    int
	Value string
}

func worker(c chan Line, done []chan bool, id int, wg *sync.WaitGroup) {
loop:
	for {
		select {

		case item := <-c:
			fmt.Printf("Line number %d have value %s\n", item.Id, item.Value)
			wg.Done()

		case result := <-done[id-1]:

			if result {
				break loop
			}

		}

	}
	fmt.Println("the loop broke")

}

func Run4() error {

	c := make(chan Line, 10)
	numofGoroutine := 3
	done := make([]chan bool, numofGoroutine)
	var wg sync.WaitGroup

	file, err := os.Open("file.txt")
	defer file.Close()

	for j := 1; j <= numofGoroutine; j++ {
		go worker(c, done, j, &wg)
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		wg.Add(1)
		i++
		line := scanner.Text()
		l := Line{Value: line, Id: i}
		c <- l

	}

	wg.Wait()

	//nhận tín hiệu true để đóng vòng lặp vô hạn sau khi đã đọc xong
	for k := range done {
		done[k] = make(chan bool, 3)
		done[k] <- true
	}

	fmt.Println("done")

	if err != nil {
		return err
	}

	return nil
}
