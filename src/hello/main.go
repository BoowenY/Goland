package main

import "fmt"

const (
	_ = iota;
	n1, n2 = iota + 1, iota + 2;
	n3, n4 = iota + 3, iota + 5;
	n5 = iota;
)

func main () {
	var a int = 10;
	var b int = 11;
	var c string = "abc";
	fmt.Println("hello world", a, b ,c, n1, n2, n3, n4, n5);
}