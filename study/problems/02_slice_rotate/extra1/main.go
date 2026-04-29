// 문제 02-extra1: 2차원 슬라이스 평탄화 (Flatten)
//
// [같은 패턴]
//   - 슬라이스 순회, append, make 용량 할당
//
// [문제]
//   2차원 정수 슬라이스 nested를 1차원 슬라이스로 평탄화하시오.
//   원본 순서 유지.
//
// [예시]
//   Flatten([][]int{{1,2},{3},{4,5,6}}) -> [1,2,3,4,5,6]
//   Flatten([][]int{})                  -> []
//   Flatten([][]int{{}, {1}, {}})       -> [1]
//
// [힌트]
//   미리 총 길이를 계산해 make([]int, 0, total)로 용량을 할당하면 효율적.
package main

import "fmt"

func Flatten(nested [][]int) []int {
	// TODO: 구현하세요.
	total := 0
	for _, inner := range nested {
		total += len(inner)
	}
	out := make([]int, 0, total)
	for _, inner := range nested {
		out = append(out, inner...)
	}
	return out
}

func main() {
	tests := []struct {
		in   [][]int
		want []int
	}{
		{[][]int{{1, 2}, {3}, {4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{[][]int{}, []int{}},
		{[][]int{{}, {1}, {}}, []int{1}},
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}},
	}

	for _, tc := range tests {
		got := Flatten(tc.in)
		fmt.Printf("Flatten(%v) = %v  | pass=%v\n", tc.in, got, equalInt(got, tc.want))
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
