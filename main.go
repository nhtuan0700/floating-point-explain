// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

func main() {
	binary := "00111101110011001100110011001101" // 0.1
	// binary = "01000010101010100100000000000000" // 85.125

	// chia ra các phần
	signBit := binary[0]       // 1 bit: 1
	exponentBit := binary[1:9] // 8 bit: 2 -> 9
	mantissaBit := binary[9:]  // 23 bit: 10 -> 32

	var sign, exponent, mantissa float64
	// Dùng để làm biến tạm thời tính từ binary -> decimal
	var decimal float64

	// sign: (-1)^signBit
	if signBit == '0' {
		sign = math.Pow(-1, 0)
	} else {
		sign = math.Pow(-1, 1)
	}

	fmt.Println("sign: ", sign)

	// exponent: 2^biasedExponent
	decimal = 0
	for i, b := range exponentBit {
		if b == '1' {
			decimal = decimal + math.Pow(2, float64(7-i))
		}
	}
	// Theo chuẩn IEEE 754, phần mũ thực sự (biased exponent) được tính bằng cách trừ đi giá trị bias là 127 (cho dạng float32):
	bias := 127
	biasedExponent := decimal - float64(bias)
	exponent = math.Pow(2, biasedExponent)
	fmt.Println("biasedExponent: ", biasedExponent)
	fmt.Println("exponent: ", exponent)

	// mantissa
	decimal = 0
	for i, b := range mantissaBit {
		if b == '1' {
			fmt.Println(-(i+1)+int(biasedExponent))
			decimal = decimal + math.Pow(2, float64(-(i+1)+int(biasedExponent)))
		}
	}
	mantissa = decimal
	fmt.Println("mantissa: ", mantissa)


	floatResult := sign * (mantissa + exponent)
	fmt.Println("floating result: ", floatResult)
}
