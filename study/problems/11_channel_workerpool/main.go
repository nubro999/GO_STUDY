// 문제 11: 채널과 워커 풀
//
// [학습 포인트]
//   - 채널 송수신 (ch <- v, v := <-ch)
//   - close(ch) 와 for-range 채널
//   - 버퍼드 채널 vs 언버퍼드 채널
//   - 워커 풀 패턴 (생산자-소비자)
//
// [문제]
//   정수 슬라이스 jobs 의 각 원소에 대해 무거운 계산 f(x)를 수행한 결과를
//   "원래 입력 순서대로" 슬라이스로 반환하시오.
//   workers 개의 고루틴이 동시에 처리하여 throughput을 높여야 한다.
//
//   func Process(jobs []int, workers int, f func(int) int) []int
//
// [예시]
//   f := func(x int) int { return x * x }
//   Process([1,2,3,4,5], 3, f) -> [1, 4, 9, 16, 25]
//
// [힌트]
//   - 입력 순서를 유지하려면 (index, value) 쌍을 채널에 흘리거나,
//     결과 배열의 i번째에 직접 쓰기 (여러 워커가 서로 다른 인덱스에만 씀 → race 없음).
//   - jobs 채널을 close 하면 워커들의 for-range 루프가 자연스럽게 종료됨.
package main

import (
	"fmt"
	"sync"
)

func Process(jobs []int, workers int, f func(int) int) []int {
	// TODO: 구현하세요.
	_ = sync.WaitGroup{}
	return nil
}

func main() {
	square := func(x int) int { return x * x }

	tests := []struct {
		jobs    []int
		workers int
		f       func(int) int
		want    []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, square, []int{1, 4, 9, 16, 25}},
		{[]int{10, 20, 30}, 1, square, []int{100, 400, 900}},
		{[]int{}, 5, square, []int{}},
		{[]int{7}, 4, square, []int{49}},
	}

	for _, tc := range tests {
		got := Process(tc.jobs, tc.workers, tc.f)
		fmt.Printf("Process(%v, workers=%d) = %v  | pass=%v\n",
			tc.jobs, tc.workers, got, equalInt(got, tc.want))
	}
}

func equalInt(a, b []int) bool {
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
