/*
 * Go生成动态库的命令:go build -o hello.so -buildmode=c-shared hello.go
 */
package main

/*
  #include <stdlib.h>

  typedef void (*HelloCB)(void);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export HelloGolang
func HelloGolang(str string) *C.char {
	cs := C.CString("Hello " + str)
	C.free(unsafe.Pointer(cs))
	return cs
}

//export HelloWorld
func HelloWorld() *C.char {
	cs := C.CString("Hello Japan")
	C.free(unsafe.Pointer(cs))
	return cs
}

//HelloSetCallback
func HelloSetCallback() {

}

func main() {
	fmt.Println("main")
}
