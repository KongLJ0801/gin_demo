/*

   @author #Kkk
   @File testCode
   @Description:
   @version 0.1
   @date 2022/8/24 19:36

*/

package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func main() {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		log.Fatalln("binary.Read failed:", err)
	}
	fmt.Println(pi)

	var reader *bufio.Reader

	earl, _ := reader.Peek(4)
	fmt.Println(earl)

}
