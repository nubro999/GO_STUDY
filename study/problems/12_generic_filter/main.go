// 문제 12: 제네릭 함수 (Map / Filter / Reduce)
//
// [학습 포인트]
//   - 타입 파라미터 [T any], [K comparable]
//   - 제약(constraint) 인터페이스
//   - 함수형 프로그래밍 패턴 in Go
//
// [문제]
//   다음 세 가지 제네릭 함수를 구현하시오.
//
//   1) Map[T, U any](xs []T, f func(T) U) []U
//      - 각 원소를 f로 변환한 새 슬라이스 반환
//
//   2) Filter[T any](xs []T, pred func(T) bool) []T
//      - pred 가 true 인 원소만 골라 새 슬라이스 반환
//
//   3) Reduce[T, U any](xs []T, init U, f func(U, T) U) U
//      - 누적값(init부터 시작)에 f를 반복 적용한 최종값 반환
//
// [예시]
//   Map([1,2,3], func(x int) int { return x*x })       -> [1,4,9]
//   Filter([1,2,3,4], func(x int) bool { return x%2==0 }) -> [2,4]
//   Reduce([1,2,3,4], 0, func(acc, x int) int { return acc+x }) -> 10
package main

import "fmt"

// TODO: Map[T, U any] 구현
// TODO: Filter[T any] 구현
// TODO: Reduce[T, U any] 구현

func main() {
	// Map 테스트 — int → int
	// squares := Map([]int{1, 2, 3}, func(x int) int { return x * x })
	// fmt.Printf("Map squares = %v  | pass=%v\n", squares, equalInt(squares, []int{1, 4, 9}))

	// Map 테스트 — int → string  (제네릭의 강점)
	// labels := Map([]int{1, 2, 3}, func(x int) string { return fmt.Sprintf("#%d", x) })
	// fmt.Printf("Map labels  = %v  | pass=%v\n", labels, equalStr(labels, []string{"#1", "#2", "#3"}))

	// Filter 테스트
	// evens := Filter([]int{1, 2, 3, 4, 5}, func(x int) bool { return x%2 == 0 })
	// fmt.Printf("Filter evens = %v  | pass=%v\n", evens, equalInt(evens, []int{2, 4}))

	// Reduce 테스트
	// sum := Reduce([]int{1, 2, 3, 4}, 0, func(acc, x int) int { return acc + x })
	// fmt.Printf("Reduce sum  = %d  | pass=%v\n", sum, sum == 10)

	fmt.Println("위 주석을 해제하고 Map/Filter/Reduce 를 구현한 뒤 실행하세요.")
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

func equalStr(a, b []string) bool {
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
