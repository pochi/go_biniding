package main

/*
#include <stdlib.h>
#include "hello.h"
*/
import "C"
import "unsafe"
import "fmt"
import "io/ioutil"

/*
  - TODO 1: loop for specific dir - 
  TODO 2: load image and transfer to c program
*/
func main() {
	str := C.CString("Hello,  cgo")
	defer C.free(unsafe.Pointer(str))
	C.hello(str)

	dir := "./images"
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
