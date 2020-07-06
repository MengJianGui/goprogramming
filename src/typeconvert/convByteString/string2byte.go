package convByteString

// string类型 convert byte切片
func Str2byte(str string) []byte {
	return []byte(str)
}

// []byte切片 convert string
// string不能直接和[]byte互换
func Byte2str(slice []byte) string {
	return string(slice[:])
}