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
	len := len(nums)
	if workers <= 0 {
		workers = 1
	}
	if len < workers {
		workers = len
	}
	if len == 0 {
		return 0
	}
	
	chunksize := (len + workers - 1) / workers
	partials := make([]int, workers)
	
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		start := i * chunksize
		end := start + chunksize
		if end > len {
			end = len
		}

		wg.Add(1)

		go func(i, start, end int) {
			defer wg.Done()
			max := nums[start] // 청크의 첫 원소로 초기화
			for j := start + 1; j < end; j++ {
				if nums[j] > max {
					max = nums[j]
				}
			}
			partials[i] = max
		}(i, start, end)
	}

	wg.Wait()

	// partials에서 최종 max 찾기
	finalMax := partials[0]
	for i := 1; i < workers; i++ {
		if partials[i] > finalMax {
			finalMax = partials[i]
		}
	}
	
	return finalMax
		
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
