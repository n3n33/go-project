// Go에서 최적화 하는 방법은
// 동시에 작업을 처리하는 것이다.
package main

import (
	"fmt"
	"time"
)

func main() {
	timeCount("yena")
	timeCount("kwon")
}

//Goroutines

func timeCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "counting...", i)
		time.Sleep(time.Second)
	}
}
