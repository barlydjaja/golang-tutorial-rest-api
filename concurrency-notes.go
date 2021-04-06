package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo(){
	start := time.Now()
	fmt.Println("running foo...")
	fmt.Println("this is foo")
	elapsed := time.Since(start)
	fmt.Println("time took to run foo: ", elapsed)
	wg.Done()
}

func bar(){
	start := time.Now()
	fmt.Println("running bar...")
	fmt.Println("this is bar")
	elapsed := time.Since(start)
	fmt.Println("time took to run bar: ", elapsed)
	wg.Done()
}

func main(){
	fmt.Println("you are using ",runtime.GOOS, " OS")
	fmt.Println("machine architecture",runtime.GOARCH)
	fmt.Println("available number of go routine",runtime.NumCPU())
	fmt.Println("number of go routine",runtime.NumGoroutine())
	wg.Add(2)
	go bar()
	go foo()
	fmt.Println("number of go routine",runtime.NumGoroutine())
	wg.Wait()
}
