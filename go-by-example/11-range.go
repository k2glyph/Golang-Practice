package main

import "fmt"

/*
range iterates over elements in a variety of data structures.
Let’s see how to use range with some of the data structures we’ve already learned.
*/
func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("Sums of all the array number: ", sum)

	// By using range
	sum = 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sums of all the  numbers: ", sum)
	dict1 := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}
	for k, v := range dict1 {
		fmt.Println("Key: ", k, " Value: ", v)
	}
	for _, v := range "golang" {
		fmt.Println(v)
	}

}
