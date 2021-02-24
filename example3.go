package main
/*
#include<a.c>
*/
import "C"
import "fmt"
func main() {
	fmt.Println("You are inside Go file")
	C.print_hello()
}
