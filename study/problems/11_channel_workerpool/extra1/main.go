// 문제 11-extra1: 3-stage Pipeline
//
// [같은 패턴]
//   - chan + close + for-range
//   - 각 stage 가 별도 고루틴, 채널로 데이터 흐름
//
// [문제]
//   다음 3-stage 파이프라인을 구현하시오.
//
//     gen(nums)        → chan int   // 입력 슬라이스를 채널로
//     square(in chan)  → chan int   // 각 값을 제곱
//     sum(in chan)     → int        // 모든 값 합산
//
//   그리고 Pipeline(nums []int) int 가 위 3 stage를 연결해 합을 반환.
//
// [예시]
//   Pipeline([1,2,3,4]) -> 30   // 1+4+9+16
//
// [힌트]
//   - 각 stage 함수는 입력 채널을 받고 출력 채널을 반환 (마지막 sum만 int 반환).
//   - 송신자가 close 책임 — gen, square 모두 고루틴 안에서 close.
package main

import "fmt"

func gen(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		// TODO: 구현
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		// TODO: 구현
	}()
	return out
}

func sum(in <-chan int) int {
	// TODO: 구현
	return 0
}

func Pipeline(nums []int) int {
	return sum(square(gen(nums)))
}

func main() {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4}, 30},     // 1+4+9+16
		{[]int{}, 0},
		{[]int{5}, 25},
		{[]int{1, 1, 1, 1}, 4},
	}

	for _, tc := range tests {
		got := Pipeline(tc.nums)
		fmt.Printf("Pipeline(%v) = %d  | pass=%v\n", tc.nums, got, got == tc.want)
	}
}
