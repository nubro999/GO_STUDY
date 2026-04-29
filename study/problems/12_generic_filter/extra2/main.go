// 문제 12-extra2: 제네릭 Min / Max
//
// [같은 패턴]
//   - 타입 파라미터 + 비교 가능한 타입 제약 (cmp.Ordered)
//
// [문제]
//   다음 두 함수를 작성하시오.
//
//     Min[T cmp.Ordered](xs []T) (T, bool)
//     Max[T cmp.Ordered](xs []T) (T, bool)
//
//   - 슬라이스가 비어있으면 zero value, false 반환
//   - 그 외엔 (min/max 값, true) 반환
//
// [예시]
//   Min([3,1,4,1,5,9,2,6])         -> 1, true
//   Max([3,1,4,1,5,9,2,6])         -> 9, true
//   Min([]int{})                   -> 0, false
//   Min([]string{"go","rust","c"}) -> "c", true
//
// [참고]
//   cmp.Ordered 는 Go 1.21+ 표준. 정수/실수/문자열 등 < 가능한 타입을 포함.
package main

import (
	"cmp"
	"fmt"
)

// TODO: Min[T cmp.Ordered] 구현
// TODO: Max[T cmp.Ordered] 구현

func main() {
	// 아래 주석을 해제하고 Min/Max를 구현한 뒤 실행하세요.
	//
	// {
	// 	v, ok := Min([]int{3, 1, 4, 1, 5, 9, 2, 6})
	// 	fmt.Printf("Min(int) = %d, %v  | pass=%v\n", v, ok, v == 1 && ok)
	// }
	// {
	// 	v, ok := Max([]int{3, 1, 4, 1, 5, 9, 2, 6})
	// 	fmt.Printf("Max(int) = %d, %v  | pass=%v\n", v, ok, v == 9 && ok)
	// }
	// {
	// 	v, ok := Min([]int{})
	// 	fmt.Printf("Min(empty) = %d, %v  | pass=%v\n", v, ok, v == 0 && !ok)
	// }
	// {
	// 	v, ok := Min([]string{"go", "rust", "c"})
	// 	fmt.Printf("Min(string) = %q, %v  | pass=%v\n", v, ok, v == "c" && ok)
	// }
	// {
	// 	v, ok := Max([]float64{3.14, 2.71, 1.41})
	// 	fmt.Printf("Max(float) = %v, %v  | pass=%v\n", v, ok, v == 3.14 && ok)
	// }

	_ = cmp.Compare[int]
	fmt.Println("Min/Max 를 구현한 뒤 main()의 주석을 해제하세요.")
}
