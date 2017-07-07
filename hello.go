package main

/*
#include <stdlib.h>
#include "hello.h"
*/
import "C"
import "unsafe"

func main() {
	str := C.CString("Hello,  cgo")
	defer C.free(unsafe.Pointer(str))
	C.hello(str)
}
