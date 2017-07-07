package main

/*
#include <stdlib.h>
#include "hello.h"
*/
import "C"

import (
	"unsafe"
	"fmt"
	"os"
	"io/ioutil"
	"image"
	_ "image/png"
	_ "image/jpeg"	
)

const BUFSIZE = 1024

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
		func() {
			file, err := os.Open(dir + "/" + f.Name())
			defer file.Close()
			if err != nil {
				panic(err)
			}

			for {
				imageData, imageType, err := image.Decode(file)
				if err != nil {
					break
				}
				fmt.Println(imageData)
				fmt.Println(imageType)
			}
		}()
	}
}
