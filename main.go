// Go에서 최적화 하는 방법은
// 동시에 작업을 처리하는 것이다.
package main

import (
	"fmt"
	"time"
)

// 메인 함수는 고루틴을 기다려주지않아,
//Goroutines

func main() {
	c := make(chan bool)
	people := [2]string{"yena", "kwon"}
	//go timeCount("yena")
	//go timeCount("kwon")
	for _, person := range people {
		//result := go timeCount(person)
		go timeCount(person, c)
	}
	result := <-c
	fmt.Println(result)
	fmt.Println(result)
	fmt.Println(result)
	fmt.Println(result)
	//time.Sleep(time.Second * 10)
	//return true
}

// Channel 은 go-routine과 메인함수에 사이에 정보를 전달해주는 방법
// go-routine 과 다른 go-routine 으로 커뮤니케이션 하는것도 가능함
// Channel은 Pipe 같은 것임

func timeCount(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	//return true
	c <- true
	//for i := 0; i < 10; i++ {
	//	fmt.Println(person, "counting...", i)
	//	time.Sleep(time.Second * 5)
	//}
}
