package main

import (
	"fmt"
	"os"
	"go_learning/rwfile"
	"go_learning/goroutines"
	"go_learning/readconcurrency"
)


func testRWFile() {
	rwfile.RWFile()
}

func testGoroutines() {
	goroutines.Goroutines()
	for {
		<- goroutines.RetValue
	}
}

func testProducer() {
	fmt.Println("start test producer")
	f, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer func(){
		if err := f.Close(); err != nil{
			panic(err)
		} 
	}()
	producer := concurrencyread.NewLineProducer(f, 100)
	producer.Produce()
	const ConcurrentNum = 2
	sentinel := make(chan bool, ConcurrentNum)
	for i := 0; i < ConcurrentNum; i++{
		go func(){
			for {
				num, sent, ok := producer.GetNumberedLine(); 
				if !ok {
					break	
				}
				fmt.Printf("%d: %s\n", num, sent)
			}
			sentinel <- true
		}()
	}
	for i := 0; i < ConcurrentNum; i++{
		<-sentinel
	}
}

func main(){
	testProducer()
}