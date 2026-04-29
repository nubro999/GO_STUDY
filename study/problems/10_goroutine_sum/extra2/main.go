// 문제 10-extra2: 안전한 병렬 카운터 (sync.Mutex)
//
// [같은 패턴]
//   - 고루틴 + WaitGroup
//   - 공유 자원 보호: sync.Mutex 또는 sync/atomic
//
// [문제]
//   N개의 고루틴이 각각 perWorker번씩 카운터를 증가시킨다.
//   최종 카운터 값은 N * perWorker 이어야 한다.
//   sync.Mutex로 보호하여 race condition 없이 구현하시오.
//
//   func SafeCounter(workers, perWorker int) int
//
// [예시]
//   SafeCounter(10, 1000) -> 10000
//
// [실행 시]
//   `go run -race ./problems/10_goroutine_sum/extra2/`
//   로 race 가 없는지 검증해보세요.
package main

import (
	"fmt"
	"sync"
)

func SafeCounter(workers, perWorker int) int {
	// TODO: 구현하세요. sync.Mutex 사용.
	_ = sync.Mutex{}
	return 0
}

func main() {
	tests := []struct {
		workers   int
		perWorker int
		want      int
	}{
		{10, 1000, 10000},
		{1, 5, 5},
		{100, 100, 10000},
		{0, 100, 0},
	}

	for _, tc := range tests {
		got := SafeCounter(tc.workers, tc.perWorker)
		fmt.Printf("SafeCounter(%d, %d) = %d  | pass=%v\n",
			tc.workers, tc.perWorker, got, got == tc.want)
	}
}
