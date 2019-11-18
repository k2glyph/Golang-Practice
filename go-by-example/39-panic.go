package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("test")
	//panic("a problem occur")
	_,e:=os.Create("/tmp/file")
	if e!=nil{
		panic(e)
	}
}