package numeric

func ExampleInt16() {
	Int16()
	// Output:
	// 2
	// -32768 32767
}

func ExampleInt32() {
	Int32()
	// Output:
	// 4
	// -2147483648 2147483647
}

func ExampleInt() {
	Int()
	// Output:
	// 8
	// -9223372036854775808 9223372036854775807
}

func ExampleInt64() {
	Int64()
	// Output:
	// 8
	// -9223372036854775808 9223372036854775807
}

func ExampleFloat32() {
	Float32()
	// Output:
	// 4
	// 1.401298464324817e-45 3.4028234663852886e+38
}

func ExampleFloat64() {
	Float64()
	// Output:
	// 8
	// 5e-324 1.7976931348623157e+308
}
