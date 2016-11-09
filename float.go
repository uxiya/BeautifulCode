package main

import (
	"fmt"
	"unsafe"
)

func main() {
	b := float32(12.5)
	c := *(*int)(unsafe.Pointer(&b))
	for i := uint(31); i > 0; i-- {
		fmt.Printf("%d", (c>>i)&0x01)
		if i%8 == 0 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("%d", (c & 0x01))
}
