// 문제 01-extra1: Collatz 수열 길이
//
// [같은 패턴]
//   - for 루프, if/else 분기, 정수 연산
//
// [문제]
//   양의 정수 n에서 시작하여 다음 규칙을 반복한다.
//     - n이 짝수면 n = n/2
//     - n이 홀수면 n = 3*n + 1
//   1에 도달하기까지 걸리는 단계 수를 반환하시오.
//
// [예시]
//   Collatz(1)  -> 0   // 이미 1
//   Collatz(6)  -> 8   // 6→3→10→5→16→8→4→2→1
//   Collatz(27) -> 111
package main

import "fmt"

func Collatz(n int) int {
	// TODO: 구현하세요.
	return 0
}

func main() {
	tests := []struct {
		n    int
		want int
	}{
		{1, 0},
		{6, 8},
		{27, 111},
		{19, 20},
	}

	for _, tc := range tests {
		got := Collatz(tc.n)
		fmt.Printf("Collatz(%d) = %d  | pass=%v\n", tc.n, got, got == tc.want)
	}
}
