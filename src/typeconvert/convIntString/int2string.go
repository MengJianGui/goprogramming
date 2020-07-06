package convIntString

import "strconv"

// int型convert string类型
func Int2string(i int) string {
	str := strconv.Itoa(i)
	// 等价于str := strconv.FormatInt(int64(int), 10)
	return str
}

// int64类型convert string类型
func Int64string(i int64) string {
	str := strconv.FormatInt(i, 10)
	return str
}

// int类型convert int64
func Int2int64(i int) int64 {
	return int64(i)
}

// string类型 convert int
func Str2int(str string) (int, error) {
	i, err := strconv.Atoi(str)
	return i, err
}

// string类型 convert int64
func Str2int64(str string) (int64, error) {
	// 第三个参数位大小表示期望转换的结果类型，其值可以为0, 8, 16, 32和64，分别对应 int, int8, int16, int32和int64
	int64, err := strconv.ParseInt(str, 10, 64)
	//strconv.ParseUint()
	return int64, err
}

// string类型 convert float
func Str2float(str string) (float64, error) {
	float, err := strconv.ParseFloat(str, 64)
	return float, err
}

// float类型convert string
func Float2str(f float64) string {
	return strconv.FormatFloat(f, 'E', -1, 64)
}