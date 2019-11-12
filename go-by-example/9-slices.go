package main

import (
	"fmt"
)

func main() {
	s:=make([]string,3)
	fmt.Println(s)
	s[0]="a"
	s[1]="b"
	s[2]="c"
	fmt.Println(s)
	fmt.Println("length of slic: ", len(s))
	s=append(s,"d")
	s=append(s, "e","f")
	fmt.Println("length of slice: ", len(s), "Elements in slice: ", s)
	c:=make([]string, len(s))
	copy(c,s)
	fmt.Println("Copy slice: ",c)
	fmt.Println("Slice value by range: ", s[2:5])
	fmt.Println("Slice upto: ", s[:5])
	fmt.Println("Slice begin from : ", s[2:])
	t:= []string {"g","h","i"}
	fmt.Println("New slice: ", t)
	r :=make([][]int,3)
	fmt.Println(r)
	for i:=0;i<3;i++ {
	   innerLen :=i+1
	   r[i] = make([]int, innerLen)
	   for j:=0;j<innerLen;j++ {
		 r[i][j]=i+j
           }
	}
       fmt.Println("Printing r slice: ", r)

}
