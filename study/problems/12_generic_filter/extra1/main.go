// 문제 12-extra1: 제네릭 Set 자료구조
//
// [같은 패턴]
//   - 타입 파라미터, 제약(constraint)
//   - comparable 제약 (== 가능한 타입만)
//
// [문제]
//   임의의 비교 가능한 타입 T에 대한 Set 자료구조를 만드시오.
//
//     Set[T comparable]   // 타입
//     NewSet[T]() *Set[T]
//     (s *Set[T]) Add(v T)
//     (s *Set[T]) Remove(v T)
//     (s *Set[T]) Contains(v T) bool
//     (s *Set[T]) Size() int
//
// [예시]
//   s := NewSet[int]()
//   s.Add(1); s.Add(2); s.Add(1)
//   s.Size()       // 2
//   s.Contains(1)  // true
//   s.Remove(1)
//   s.Contains(1)  // false
//
// [힌트]
//   내부 자료구조는 map[T]struct{} 가 표준 — 빈 struct는 0바이트.
package main

import "fmt"

// TODO: Set[T comparable] 정의 및 메서드 구현

func main() {
	// 아래 주석을 해제하고 Set 을 구현한 뒤 실행하세요.
	//
	// s := NewSet[int]()
	// s.Add(1)
	// s.Add(2)
	// s.Add(1)
	// fmt.Printf("Size = %d (want 2)  | pass=%v\n", s.Size(), s.Size() == 2)
	// fmt.Printf("Contains(1) = %v (want true)\n", s.Contains(1))
	// fmt.Printf("Contains(3) = %v (want false)\n", s.Contains(3))
	// s.Remove(1)
	// fmt.Printf("After Remove(1), Contains(1) = %v (want false)\n", s.Contains(1))
	// fmt.Printf("Size = %d (want 1)  | pass=%v\n", s.Size(), s.Size() == 1)
	//
	// // 다른 타입에서도 동작
	// ss := NewSet[string]()
	// ss.Add("go"); ss.Add("rust"); ss.Add("go")
	// fmt.Printf("string Set Size = %d (want 2)\n", ss.Size())

	fmt.Println("Set 을 구현한 뒤 main()의 주석을 해제하세요.")
}
