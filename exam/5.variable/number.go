package main

import "fmt"

func main() {

	var (
		// INT
		varINT8 int8 = 127
		varINT8Minimum int8 = -128
		varINT16 int16 = 32767
		varINT16Minimum int16 = -32768
		varINT32 int32 = 2147483647
		varINT32Minimum int32 = -2147483648
		varINT64 int64 = 9223372036854775807
		varINT64Minimum int64 = -9223372036854775808

		// Unsigned INT
		varUINT8 uint8 = 255
		varUINT8Minimum uint8 = 0
		varUINT16 uint16 = 65535
		varUINT16Minimum uint16 = 0
		varUINT32 uint32 = 4294967295
		varUINT32Minimum uint32 = 0
		varUINT64 uint64 = 18446744073709551615
		varUINT64Minimum uint64 = 0

		// Alias
		varBYTE byte = 127
		varRUNE int32 = 2147483647
		varINT int = 2147483647 //int32/64 os
		varUINT uint = 4294967295 //uint32/64 os

		// Float
		varFLOAT32 float32 = 3.4028234
		varFLOAT64 float64 = 1.79769313486231570814527423731704356798070e+308
	)

	fmt.Println("INT 8\t\t\t: ", varINT8)
	fmt.Println("INT 16\t\t\t: ", varINT16)
	fmt.Println("INT 32\t\t\t: ", varINT32)
	fmt.Println("INT 64\t\t\t: ", varINT64)
	fmt.Println("=============================================================")
	fmt.Println("MINIMUM INT 8\t\t: ", varINT8Minimum)
	fmt.Println("MINIMUM INT 16\t\t: ", varINT16Minimum)
	fmt.Println("MINIMUM INT 32\t\t: ", varINT32Minimum)
	fmt.Println("MINIMUM INT 64\t\t: ", varINT64Minimum)
	fmt.Println("=============================================================")
	fmt.Println("UINT 8\t\t\t: ", varUINT8)
	fmt.Println("UINT 16\t\t\t: ", varUINT16)
	fmt.Println("UINT 32\t\t\t: ", varUINT32)
	fmt.Println("UINT 64\t\t\t: ", varUINT64)
	fmt.Println("=============================================================")
	fmt.Println("MINIMUM UINT 8\t\t: ", varUINT8Minimum)
	fmt.Println("MINIMUM UINT 16\t\t: ", varUINT16Minimum)
	fmt.Println("MINIMUM UINT 32\t\t: ", varUINT32Minimum)
	fmt.Println("MINIMUM UINT 64\t\t: ", varUINT64Minimum)
	fmt.Println("=============================================================")
	fmt.Println("BYTE AS UINT8\t\t: ", varBYTE)
	fmt.Println("RUNE AS INT32\t\t: ", varRUNE)
	fmt.Println("INT AS Minimal INT32\t: ", varINT)
	fmt.Println("UINT AS Minimal UINT32\t: ", varUINT)
	fmt.Println("=============================================================")
	fmt.Println("FLOAT32\t\t: ", varFLOAT32)
	fmt.Println("FLOAT64\t\t: ", varFLOAT64)

}