package numeric

import (
	"fmt"
	"math"
	"unsafe"
)

// int16
func Int16() {
	var maxInt16 int16 = math.MaxInt16
	// 占用空间
	// 2
	fmt.Println(unsafe.Sizeof(maxInt16))
	// 最小值与最大值
	// -32768 32767
	fmt.Println(math.MinInt16, math.MaxInt16)
}

// int32
func Int32() {
	var maxInt32 int32 = math.MaxInt32
	// 占用空间
	// 4
	fmt.Println(unsafe.Sizeof(maxInt32))
	// 最小值与最大值
	// -2147483648 2147483647
	fmt.Println(math.MinInt32, math.MaxInt32)
}

// int
func Int() {
	var maxInt int = math.MaxInt
	// 占用空间
	// 8
	fmt.Println(unsafe.Sizeof(maxInt))
	// 最小值与最大值
	// -9223372036854775808 9223372036854775807
	fmt.Println(math.MinInt, math.MaxInt)
}

// int64
func Int64() {
	var maxInt64 int64 = math.MaxInt64
	// 占用空间
	// 8
	fmt.Println(unsafe.Sizeof(maxInt64))
	// 最小值与最大值
	// -9223372036854775808 9223372036854775807
	fmt.Println(math.MinInt64, math.MaxInt64)
}

// float32
func Float32() {
	var maxFloat32 float32 = math.MaxFloat32
	// 占用空间
	// 4
	fmt.Println(unsafe.Sizeof(maxFloat32))
	// 最小值与最大值
	// 1.401298464324817e-45 3.4028234663852886e+38
	fmt.Println(math.SmallestNonzeroFloat32, math.MaxFloat32)
}

// float64
func Float64() {
	var maxFloat64 float64 = math.MaxFloat64
	// 占用空间
	// 4
	fmt.Println(unsafe.Sizeof(maxFloat64))
	// 最小值与最大值
	// 5e-324 1.7976931348623157e+308
	fmt.Println(math.SmallestNonzeroFloat64, math.MaxFloat64)
}
