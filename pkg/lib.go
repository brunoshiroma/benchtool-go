package main

import "C"
import "fmt"

//export benchtoolGoCall
func benchtoolGoCall(iteration, repeat, benchType int) *C.char {
	return C.CString(fmt.Sprintf("%d %d %d", iteration, repeat, benchType))
}

func main() {

}
