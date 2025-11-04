package main

import (
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	count := "11"
	for i := 3; i <= n; i++ {
		count += "*"
		temp := ""
		counter := 1
		for j := 0; j < len(count)-1; j++ {
			if count[j+1] != count[j] {
				temp += strconv.Itoa(counter) + string(count[j])
				counter = 1
			} else {
				counter++
			}
		}
		count = temp
	}
	return count
}
