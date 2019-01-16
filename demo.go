package main

import "fmt"

var a = 'a'
var b = "a"
var c bool = true
var d = 123
var e = 321

func main() {

	var f = d + e

	//第一个go程序
	fmt.Println(a, b, c, f)
}
