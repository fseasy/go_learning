package main

import (
	"fmt"
	"go_learning/rwfile"
	"go_learning/goroutines"
)

func main(){
	fmt.Printf("%s", "Hello World")
	rwfile.RWFile()
	goroutines.Goroutines()
	for {
		<- goroutines.RetValue
	}
}