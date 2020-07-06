package convIntString

import (
	"fmt"
	"testing"
)

func TestFloat2str(t *testing.T) {
	str := Float2str(3.1415926)
	fmt.Printf("float to string: %T, %v", str, str)
}

func TestInt2int64(t *testing.T) {
	temp := Int2int64(16)
	fmt.Printf("int to int64: %T, %v", temp, temp)
}

func TestInt2string(t *testing.T) {
	str := Int2string(666)
	fmt.Printf("float to string: %T, %v", str, str)
}

func TestInt64string(t *testing.T) {
	str := Int2string(666)
	fmt.Printf("float to string: %T, %v", str, str)
}

func TestStr2float(t *testing.T) {
	float, err := Str2float("3.1415926")
	if err != nil {
		fmt.Printf("Str2float error{%v}!", err)
	}else {
		fmt.Printf("string to float: %T, %v", float, float)
	}
}

func TestStr2int(t *testing.T) {
	i, err := Str2int("666")
	if err != nil {
		fmt.Printf("Str2int error{%v}!", err)
	}else {
		fmt.Printf("string to int: %T, %v", i, i)
	}
}

func TestStr2int64(t *testing.T) {
	i, err := Str2int64("666")
	if err != nil {
		fmt.Printf("Str2int64 error{%v}!", err)
	}else {
		fmt.Printf("string to int64: %T, %v", i, i)
	}
}