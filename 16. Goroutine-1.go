package main

import (
	"fmt"
	"time"
)

func main() {
	exampleGoroutine()
}

func exampleGoroutine(){
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
	time.Sleep(2 * time.Second)
}