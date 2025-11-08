package main

import (
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1 := len(num1)
	n2 := len(num2)

	result := make([]int, n1+n2)

	for i := n1 - 1; i >= 0; i-- {
		digit1 := int(num1[i] - '0')
		for j := n2 - 1; j >= 0; j-- {
			digit2 := int(num2[j] - '0')
			mul := digit1 * digit2
			p1 := i + j
			p2 := i + j + 1
			sum := mul + result[p2]

			result[p2] = sum % 10
			result[p1] += sum / 10
		}
	}

	var sb strings.Builder
	k := 0
	for k < len(result) && result[k] == 0 {
		k++
	}
	for ; k < len(result); k++ {
		sb.WriteByte(byte(result[k] + '0'))
	}
	return sb.String()
}
