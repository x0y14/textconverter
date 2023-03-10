package textconverter

import (
	"strings"
)

// 桁の多い方の先頭から10出割って、あまりとなった番号を取り出す
// これマジ天才好き
// source: https://stackoverflow.com/questions/68072383/how-to-convert-an-integer-to-a-string-without-buildin-java-methods
func i2s(i int) string {
	// ゼロのときだけ反応しないっぽいので直接返してあげる
	if i == 0 {
		return "0"
	}
	result := ""
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for n := i; n > 0; n /= 10 {
		result = digits[n%10] + result
	}
	return result
}

func beki[T float64 | int](x, n T) T {
	if n == 0 {
		return 1
	}
	var result = x
	for i := n; i > 1; i-- {
		result *= x
	}
	return result
}

func beki10Minus(n int) float64 {
	if n == 0 {
		return 1
	}
	result := float64(1)
	for i := n; i < 0; i++ {
		result /= 10.0
	}
	return result
}

func stringNum2NumI(s string) int {
	var n int
	switch s {
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
	return n
}

func stringNum2NumF(s string) float64 {
	var n float64
	switch s {
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
	return n
}

// 桁の小さい方から文字列を取り出して掛けてく
// 注意 文字列アクセスが必要
func s2i(s string) int {
	result := 0
	numOfDig := 0
	ss := strings.Split(s, "")
	for len(ss) > numOfDig {
		cur := ss[len(ss)-(1+numOfDig)]
		n := stringNum2NumI(cur)
		if numOfDig != 0 {
			result += n * (beki(10, numOfDig))
		} else {
			result += n
		}
		numOfDig++
	}
	return result
}

func i2f(i int) float64 {
	// integer -> string
	// にしてからs2iと同じように解析してfloatを作り出す
	s := i2s(i)

	result := 0.0
	numOfDig := 0
	ss := strings.Split(s, "")
	for len(ss) > numOfDig {
		cur := ss[len(ss)-(1+numOfDig)]
		n := stringNum2NumF(cur)
		if numOfDig != 0 {
			result += n * beki[float64](10.0, i2f(numOfDig))
		} else {
			result += n
		}
		numOfDig++
	}
	return result
}

// 数学を使って小数点がどこにあるかを検知しなければならない
// がおそらく難しいのでチートとして切り捨てを使用する
// 常用対数を使えばいいらしい？ 高校の数学もう忘れたから無理
// むりでした
//func f2s(f float64) string {
//	// 小数点の位置の特定
//	count := 0
//	// 切り捨てをして整数部を取り出す
//	integer := math.Floor(f)
//	// 小数部を取り出す
//	var decimal float64
//	decimal = f - integer
//	for {
//		// 誤差が生まれてしまうのでどれぐらいを許容するのかはテスト等から考えてください
//		//if math.Floor(decimal*10)-decimal*10 == 0.0 ||
//		//	math.Floor(decimal*10)-decimal*10 < 0.0001 {
//		//	break
//		//}
//		x := decimal*10 - math.Floor(decimal*10)
//		if x == 0.0 || x < 0.00000001 {
//			break
//		}
//
//		decimal = decimal * 10
//		count++
//	}
//	fmt.Println(f)
//	fmt.Println("count:", count)
//	return ""
//}

func s2f(s string) float64 {
	result := 0.0
	numOfDig := 0
	ss := strings.Split(s, "")

	// 小数点の存在を確認
	pos := 0
	total := len(ss)
	hasPoint := false
	pointPos := 0
	for total > pos {
		if ss[pos] == "." {
			hasPoint = true
			pointPos = pos
		}
		pos++
	}

	// 小数点のない場合
	// float64(23)
	if !hasPoint {
		for len(ss) > numOfDig {
			cur := ss[len(ss)-(1+numOfDig)]
			n := stringNum2NumF(cur)
			if numOfDig != 0 {
				result += n * beki[float64](10.0, i2f(numOfDig))
			} else {
				result += n
			}
			numOfDig++
		}

		return result
	}

	// 整数部
	for pointPos > numOfDig {
		cur := ss[pointPos-(1+numOfDig)]
		n := stringNum2NumF(cur)
		if numOfDig != 0 {
			result += n * beki[float64](10.0, i2f(numOfDig))
		} else {
			result += n
		}
		numOfDig++
	}
	// 小数部
	numOfDig = 1
	for len(ss)-pointPos > numOfDig {
		cur := ss[pointPos+numOfDig]
		n := stringNum2NumF(cur)
		result += n * beki10Minus(-numOfDig)
		//result += n * math.Pow(10.0, float64(-numOfDig))
		numOfDig++
	}

	return result

}
