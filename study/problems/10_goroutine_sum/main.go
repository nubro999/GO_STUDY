// 문제 10: 고루틴으로 병렬 합산
//
// [학습 포인트]
//   - go 키워드로 고루틴 생성
//   - sync.WaitGroup 으로 완료 대기
//   - 결과 수집 채널 (chan int)
//   - go run -race 로 race condition 검출
//
// [문제]
//   큰 정수 슬라이스 nums와 분할 개수 workers가 주어진다.
//   슬라이스를 workers 등분해 각 청크를 별도 고루틴이 합산하고,
//   최종 합을 반환하시오.
//
//   - workers <= 0 인 경우 1로 처리
//   - len(nums) < workers 면 workers를 len(nums)로 줄임
//   - 빈 슬라이스는 0 반환
//
// [예시]
//   ParallelSum([1..100], 4) -> 5050
//
// [힌트]
//   각 고루틴이 chan int 로 부분합을 보내고, main에서 누적.
//   또는 결과를 슬라이스의 i번째 원소에 쓰는 패턴(서로 다른 인덱스에만 쓰면 race 없음).
package main

import (
	"fmt"
	"sync"
)

func ParallelSum(nums []int, workers int) int {
	// TODO: 구현하세요.
	_ = sync.WaitGroup{}
	return 0
}

func main() {
	// 1..100 합 = 5050
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i + 1
	}

	tests := []struct {
		nums    []int
		workers int
		want    int
	}{
		{nums, 4, 5050},
		{nums, 7, 5050}, // 100 % 7 != 0 인 케이스
		{nums, 1, 5050},
		{[]int{}, 4, 0},
		{[]int{42}, 10, 42},
	}

	for _, tc := range tests {
		got := ParallelSum(tc.nums, tc.workers)
		fmt.Printf("ParallelSum(len=%d, workers=%d) = %d  | pass=%v\n",
			len(tc.nums), tc.workers, got, got == tc.want)
	}
}
