// 문제 01: FizzBuzz
//
// [학습 포인트]
//   - for 루프, if/switch, 나머지 연산자(%)
//   - fmt.Println / fmt.Sprintf
//
// [문제]
//   1부터 n까지의 정수를 다음 규칙에 따라 슬라이스로 반환하시오.
//     - 3의 배수면 "Fizz"
//     - 5의 배수면 "Buzz"
//     - 3과 5의 공배수면 "FizzBuzz"
//     - 그 외에는 숫자 자체를 문자열로
//
// [예시]
//   FizzBuzz(5)  -> ["1", "2", "Fizz", "4", "Buzz"]
//   FizzBuzz(15) -> [..., "13", "14", "FizzBuzz"]
package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) []string {
	// TODO: 구현하세요.
	out := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			out = append(out, "FizzBuzz")
		case i%3 == 0:
			out = append(out, "Fizz")
		case i%5 == 0:
			out = append(out, "Buzz")
		default:
			out = append(out, strconv.Itoa(i))
		}
	}
	return out
}

func main() {
	tests := []struct {
		n    int
		want []string
	}{
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{
			"1", "2", "Fizz", "4", "Buzz",
			"Fizz", "7", "8", "Fizz", "Buzz",
			"11", "Fizz", "13", "14", "FizzBuzz",
		}},
	}

	for _, tc := range tests {
		got := FizzBuzz(tc.n)
		ok := equal(got, tc.want)
		fmt.Printf("FizzBuzz(%d) = %v  | pass=%v\n", tc.n, got, ok)
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
