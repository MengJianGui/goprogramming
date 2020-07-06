package convByteString

import (
	"fmt"
	"testing"
)

func TestByte2str(t *testing.T) {
	byte := []uint8{109, 101, 110, 103, 106, 105, 97, 110, 103, 117, 105}
	str := Byte2str(byte)
	fmt.Printf("[]byte to string：%T, %v", str, str)
}

func TestStr2byte(t *testing.T) {
	slice := Str2byte("mengjiangui")
	fmt.Printf("string to []byte：%T, %v", slice, slice)
}