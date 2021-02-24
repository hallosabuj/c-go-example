package main
/*
#include<stdio.h>
void callFromC(void){
	printf("In c code\n");
    	fflush(stdout);
}
*/
import "C"
import "fmt"
func main() {
	fmt.Println("Inside GO code")
	C.callFromC()
}
