// 문제 01-extra2: 자릿수 합이 소수인 수의 개수
//
// [같은 패턴]
//   - for 루프 (이중), 조건 분기, 보조 함수
//
// [문제]
//   1부터 n까지의 정수 중에서, 각 자리수의 합이 소수인 수가 몇 개인지 반환하시오.
//
// [예시]
//   CountDigitSumPrime(10) -> 4   // 2,3,5,7 (1자리 소수 그대로)
//   CountDigitSumPrime(20) -> 7   // 위 4개 + 11(1+1=2), 12(1+2=3), 14(1+4=5)
//
// [힌트]
//   - 자릿수 합: while n>0 { sum += n%10; n /= 10 }
//   - 소수 판별: 2 미만 false, 2면 true, 그 외 i*i <= n 까지 시도
package main

import "fmt"

func digitSum(n int) int {
	// TODO: 구현하세요.
	return 0
}

func isPrime(n int) bool {
	// TODO: 구현하세요.
	return false
}

func CountDigitSumPrime(n int) int {
	// TODO: 구현하세요.
	return 0
}

func main() {
	tests := []struct {
		n    int
		want int
	}{
		{10, 4},
		{20, 7},
		{1, 0},
		{30, 11},
	}

	for _, tc := range tests {
		got := CountDigitSumPrime(tc.n)
		fmt.Printf("CountDigitSumPrime(%d) = %d  | pass=%v\n", tc.n, got, got == tc.want)
	}
}
