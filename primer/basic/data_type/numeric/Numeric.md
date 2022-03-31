# Numeric

*numeric type* 表示整数或浮点值的集合。预先声明的与架构无关的数字类型是：

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

*n* 位整数的值是 *n* 位宽，并使用 [二进制补码算法表示](https://en.wikipedia.org/wiki/Two's_complement)。

还有一组预先声明的数字类型具有特定于实现的大小：

```
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

为了避免可移植性问题，所有数值类型都是 [defined types](https://go.dev/ref/spec#Type_definitions)，因此互相之间都是不同的，除了 `byte` 和 `rune`，`byte` 是 `uint8` 的[别名](https://go.dev/ref/spec#Alias_declarations)，`rune` 是 `int32` 的别名。当在表达式或赋值中混合不同的数字类型时，需要显式转换。例如，`int32` 和 `int` 不是相同的类型，即使它们在某个特定的架构上可能具有相同的数据位数。

