package main

import "fmt"

func main() {
	/*
		go语言的数据类型分成两类：一类是基本类型、另一类是派生类型（复合类型）
		基本类型有数值类型、浮点类型、布尔类型、复数类型、字符串类型、字节类型
			数值类型：
				有符号位：int、int8、int16、int32、int64
				无符号位：uint、uint8、uint16、uint32、uint64（无符号位是不可以存储负数的）
			浮点类型：float32、float64
			复数类型：complex64、complex128
			布尔类型：bool
			字符串类型：string
			字节类型：byte、rune，其中byte等同于uint8，rune等同于int32，用于表示UTF-8字符串的Unicode码点
	*/
	// 浮点类型
	fmt.Println("浮点类型")
	var f1 float32 = 5.55
	var f2 float64 = 99.99
	fmt.Println("f1 float32类型变量值是:", f1)
	fmt.Println("f2 float64类型变量值是:", f2)
}
