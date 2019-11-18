package main

import (
	"fmt"
	"sort"
)

func main() {
	ints:=[]int {1,7,5,6}
	fmt.Println("Unsorted integers", ints)
	sort.Ints(ints)
	fmt.Println("Sorting, unsorted integers:", ints)
	strs:=[]string{"d","a","b","c"}
	fmt.Println("Unsorted strings", strs)
	sort.Strings(strs)
	fmt.Println("sorted strings", strs)



}