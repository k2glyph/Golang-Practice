package main

import (
	"fmt"
	"os"
)

func main() {
	f:=CreateFile("/tmp/test.txt")
	defer CloseFile(f)
	WriteFile(f)
}

func CreateFile(path string) *os.File{
	fmt.Println("Creating file")
	f,e:=os.Create(path)
	if e!=nil {
		panic(e)
	}
	return f
}

func CloseFile(file *os.File) {
	fmt.Println("closing")
	file.Close()
}
func WriteFile(file *os.File){
	fmt.Println("Writing inside file")
	file.WriteString("This is a example of test write")
}