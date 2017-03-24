package goroutines

import "fmt"

// RetValue ... 
// 	channel for 
var RetValue = make(chan int)

func echo(i int) {
	fmt.Printf("%d\n", i)
	RetValue <- 1
}

// Goroutines ...
//	test goroutines
func Goroutines() {
	for i := 0; i < 10; i++ {
		fmt.Println("call goroutinues")
		go echo(i)
	}

}