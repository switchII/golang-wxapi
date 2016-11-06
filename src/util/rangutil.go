package util

import (
	"math/rand"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母

)

// 字符截取
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start

	}
	end = start + length

	if start > end {
		start, end = end, start

	}

	if start < 0 {
		start = 0

	}
	if start > rl {
		start = rl

	}
	if end < 0 {
		end = 0

	}
	if end > rl {
		end = rl

	}
	return string(rs[start:end])

}

// 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))

	}
	return result
}

// func main() {
// fmt.Println("num:   " + string(Krand(16, KC_RAND_KIND_NUM)))
// fmt.Println("lower: " + string(Krand(16, KC_RAND_KIND_LOWER)))
// fmt.Println("upper: " + string(Krand(16, KC_RAND_KIND_UPPER)))
// fmt.Println("all:   " + string(Krand(16, KC_RAND_KIND_ALL)))
// }
