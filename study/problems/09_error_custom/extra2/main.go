// 문제 09-extra2: 에러 체인 깊이 측정
//
// [같은 패턴]
//   - fmt.Errorf("%w", err), errors.Unwrap
//
// [문제]
//   ChainDepth(err error) int 함수를 작성:
//     - err 이 nil 이면 0
//     - 그 외에는 errors.Unwrap 으로 체인을 따라가며 깊이 카운트
//     - 가장 안쪽 에러(더 이상 Unwrap 안 됨)도 1로 카운트
//
// [예시]
//   ChainDepth(nil)                                       -> 0
//   ChainDepth(errors.New("a"))                           -> 1
//   ChainDepth(fmt.Errorf("b: %w", errors.New("a")))      -> 2
//   ChainDepth(fmt.Errorf("c: %w", fmt.Errorf("b: %w",
//                          errors.New("a"))))             -> 3
//
// [힌트]
//   for err != nil {
//       depth++
//       err = errors.Unwrap(err)
//   }
package main

import (
	"errors"
	"fmt"
)

func ChainDepth(err error) int {
	// TODO: 구현하세요.
	return 0
}

func main() {
	tests := []struct {
		err  error
		want int
	}{
		{nil, 0},
		{errors.New("a"), 1},
		{fmt.Errorf("b: %w", errors.New("a")), 2},
		{fmt.Errorf("c: %w", fmt.Errorf("b: %w", errors.New("a"))), 3},
		{fmt.Errorf("no wrap %v", errors.New("a")), 1}, // %v 는 래핑 아님
	}

	for _, tc := range tests {
		got := ChainDepth(tc.err)
		fmt.Printf("ChainDepth(%v) = %d  | pass=%v\n", tc.err, got, got == tc.want)
	}
}
