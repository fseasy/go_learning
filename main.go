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
	producer := concurrencyread.LineProducer{f: f}
	producer.Produce()
	for ! producer.
}

func main(){

}