package main

func divide(dividend int, divisor int) int {
	const INT_MAX = 1<<31 - 1
	const INT_MIN = -1 << 31

	// オーバーフロー防止
	if dividend == INT_MIN && divisor == -1 {
		return INT_MAX
	}

	signPositive := !((dividend < 0) != (divisor < 0))

	absDividend := getAbs(dividend)
	absDivisor := getAbs(divisor)
	output := 0

	for absDividend >= absDivisor {
		tempDivisor := absDivisor
		multiple := 1

		for absDividend >= (tempDivisor << 1) {
			// オーバーフロー防止
			if tempDivisor<<1 < 0 {
				break
			}
			tempDivisor <<= 1
			multiple <<= 1
		}

		absDividend -= tempDivisor
		output += multiple
	}

	if !signPositive {
		output = -output
	}

	if output > INT_MAX {
		return INT_MAX
	}
	if output < INT_MIN {
		return INT_MIN
	}

	return output
}

func getAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
