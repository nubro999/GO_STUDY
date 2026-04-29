// 문제 10-extra1: 병렬 Max 찾기
//
// [같은 패턴]
//   - 고루틴 + WaitGroup + 부분 결과 슬라이스
//
// [문제]
//   정수 슬라이스 nums를 workers 개로 나눠 각 청크에서 max를 찾고,
//   최종 max를 반환. 빈 슬라이스는 0 반환.
//
//   ParallelSum (10번 문제) 과 동일한 골격으로 작성.
//
// [예시]
//   ParallelMax([3,1,4,1,5,9,2,6,5,3,5,8,9,7,9,3], 4) -> 9
package main

import (
	"fmt"
	"sync"
)

func ParallelMax(nums []int, workers int) int {
	// TODO: 구현하세요.
	_ = sync.WaitGroup{}
	return 0
}

func main() {
	tests := []struct {
		nums    []int
		workers int
		want    int
	}{
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3}, 4, 9},
		{[]int{-5, -3, -10, -1}, 2, -1},
		{[]int{42}, 8, 42},
		{[]int{}, 4, 0},
	}

	for _, tc := range tests {
		got := ParallelMax(tc.nums, tc.workers)
		fmt.Printf("ParallelMax(%v, %d) = %d  | pass=%v\n", tc.nums, tc.workers, got, got == tc.want)
	}
}
