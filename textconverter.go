package textconverter

import (
	"strings"
)

// 桁の多い方の先頭から10出割って、あまりとなった番号を取り出す
// これマジ天才好き
// source: https://stackoverflow.com/questions/68072383/how-to-convert-an-integer-to-a-string-without-buildin-java-methods
func i2s(i int) string {
	result := ""
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for n := i; n > 0; n /= 10 {
		result = digits[n%10] + result
	}
	return result
}

func beki(x, n int) int {
	if n == 0 {
		return 1
	}
	var result = x
	for i := n; i > 1; i-- {
		result *= x
	}
	return result
}
func s2i(s string) int {
	result := 0
	numOfDig := 0
	//digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	ss := strings.Split(s, "")

	for len(s) > numOfDig {
		cur := ss[len(ss)-(1+numOfDig)]
		var n int
		switch cur {
		case "0":
			n = 0
		case "1":
			n = 1
		case "2":
			n = 2
		case "3":
			n = 3
		case "4":
			n = 4
		case "5":
			n = 5
		case "6":
			n = 6
		case "7":
			n = 7
		case "8":
			n = 8
		case "9":
			n = 9
		}
		if numOfDig != 0 {
			result += n * (beki(10, numOfDig))
		} else {
			result += n
		}
		numOfDig++
	}

	return result
}
