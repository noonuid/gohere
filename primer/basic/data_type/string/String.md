# String

*string type* 表示一个字符串。字符串值是一个（可能是空的）字节序列。字节数称为字符串的长度，永远不会为负数。字符串是不可变的：一旦创建，就不可能更改字符串的内容。预先声明的字符串类型是 `string`; 它是一个 [defined type](https://go.dev/ref/spec#Type_definitions)。

字符串的长度 `s` 可以通过内置函数 [`len`](https://go.dev/ref/spec#Length_and_capacity) 获取。如果字符串是常量，则长度是编译时常量。字符串的字节可以通过整数[索引](https://go.dev/ref/spec#Index_expressions) `0` 到 `len(s)-1` 访问。取这种字节元素的地址是非法的；如果 `s[i]` 是字符串的第 `i` 个字节，则 `&s[i]` 无效。