package main

import (
	"fmt"
	"strings"
)
func Index(strs []string, t string) int {
	for i,str:=range strs {
		if str==t {
			return i
		}
	}
	return -1
}
func IsExists(strs []string, t string) bool {
	return Index(strs, t)>=0
}
func Any(strs []string, f func(string)bool)bool {
	for _,str:=range strs{
		if f(str) {
			return true
		}
	}
	return false
}
func All(strs []string, f func(string) bool) bool{
	for _,str:=range strs {
		if !f(str){
			return false
		}
	}
	return true
}
func Filter(strs []string, f func(string)bool) []string {
	strf:=make([]string,0)
	for _,str:=range strs {
		if f(str){
			strf = append(strf, str)
		}
	}
	return strf
}
func Map(strs []string, f func(string) string) [] string{
	mstrs:=make([]string,len(strs))
	for i,str:=range strs {
		mstrs[i]=f(str)
	}
	return mstrs
}
func main() {
	var strs=[]string {"apple", "google", "microsoft", "google"}
	fmt.Println(Index(strs,"google"))
	fmt.Println(IsExists(strs,"google"))
	fmt.Println(Any(strs, func(v string)bool{
		return strings.HasPrefix(v, "g")
	}))
	fmt.Println(All(strs, func(v string)bool{
		return strings.HasPrefix(v, "g")
	}))
	fmt.Println(Filter(strs, func(str string)bool{
		return strings.Contains(str, "googl")
	}))
	fmt.Println(Map(strs, func(str string) string {
		return strings.ToUpper(str)
	}))
	fmt.Println(Map(strs, strings.ToUpper))

}