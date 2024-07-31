package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"strconv"
)

func main() {
	// Basic Arithmetic Operations
	var a, b int = 15, 10
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Modulus:", a%b)

	//arithimetic by Vishnu
	n, m := 5, 10
	fmt.Println("add:", n+m)
	fmt.Println("sub:", n-m)
	fmt.Println("mul:", n*m)
	fmt.Println("dev:", n/m)

	// Working with Floats
	var c, d float64 = 3.7, 2.5
	fmt.Println("Float Addition:", c+d)
	fmt.Println("Float Subtraction:", c-d)
	fmt.Println("Float Multiplication:", c*d)
	fmt.Println("Float Division:", c/d)

	// Type Conversion
	var e int = 42
	var f float64 = float64(e)
	fmt.Println("Integer to Float:", f)

	var g float64 = 42.5
	var h int = int(g)
	fmt.Println("Float to Integer:", h)

	//vishnu float conversion
	var o float64 = float64(m)
	fmt.Println("float convresion by Visnu:", o)
	fmt.Println("float convresion type by Vishnu:", reflect.TypeOf(o))

	// Working with Complex Numbers
	var complex1 complex128 = complex(3, 4) // 3 + 4i
	var complex2 complex128 = complex(1, 2) // 1 + 2i
	fmt.Println("Complex Addition:", complex1+complex2)
	fmt.Println("Complex Subtraction:", complex1-complex2)
	fmt.Println("Complex Multiplication:", complex1*complex2)
	fmt.Println("Complex Division:", complex1/complex2)
	fmt.Println("Complex Absolute Value:", cmplx.Abs(complex1))

	// Math Functions
	fmt.Println("Power:", math.Pow(2, 3))      // 2^3
	fmt.Println("Square Root:", math.Sqrt(16)) // √16
	fmt.Println("Sin:", math.Sin(math.Pi/2))   // sin(π/2)
	fmt.Println("Cos:", math.Cos(math.Pi))     // cos(π)
	fmt.Println("Log:", math.Log(math.E))      // ln(e)

	// math by Vishnu
	fmt.Println("power math by Vishnu:", math.Pow(2, 3))
	fmt.Println("square root by Vishnu:", math.Sqrt(16))

	// String to Number Parsing
	numStr := "3.14159"
	parsedFloat, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Error parsing string to float:", err)
	} else {
		fmt.Println("Parsed Float:", parsedFloat)
	}
}
