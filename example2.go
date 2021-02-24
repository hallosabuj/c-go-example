package main
/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import "fmt"
func main() {
	fmt.Println("You are inside Go code")
	go2c := "Printed from C.printf"
        var cstr *C.char = C.CString(go2c)
        //defer C.free(unsafe.Pointer(cstr))
        C.puts(cstr)
	C.fflush(C.stdout)
}

