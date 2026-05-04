// 문제 02-extra2: 슬라이스를 일정 크기로 분할 (Chunk)
//
// [같은 패턴]
//   - 슬라이싱 a[i:j], 길이 계산, 모듈러
//
// [문제]
//
//	정수 슬라이스 nums와 양의 정수 size가 주어진다.
//	nums를 size 크기의 청크들로 나눈 [][]int 를 반환하시오.
//	마지막 청크는 size보다 작을 수 있다.
//
// [예시]
//
//	Chunk([1,2,3,4,5], 2) -> [[1,2], [3,4], [5]]
//	Chunk([1,2,3,4],   4) -> [[1,2,3,4]]
//	Chunk([1,2,3,4],   5) -> [[1,2,3,4]]
//	Chunk([],          3) -> []
//
// [힌트]
//   - i가 0, size, 2*size, ... 로 진행
//   - end := min(i+size, len(nums))
package main

import "fmt"

func Chunk(nums []int, size int) [][]int {
	// TODO: 구현하세요.
	if size <= 0 {
		return [][]int{}
	}

	n := len(nums)
	out := make([][]int, 0, (len(nums)+size-1)/size) // size로 나눈 올림
	for i := 0; i < len(nums); i += size {
		end := i + size
		if end > n {
			end = n
		}
		out = append(out, nums[i:end])
	}

	return out
}

func main() {
	tests := []struct {
		nums []int
		size int
		want [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{[]int{1, 2, 3, 4}, 4, [][]int{{1, 2, 3, 4}}},
		{[]int{1, 2, 3, 4}, 5, [][]int{{1, 2, 3, 4}}},
		{[]int{}, 3, [][]int{}},
	}

	for _, tc := range tests {
		got := Chunk(tc.nums, tc.size)
		fmt.Printf("Chunk(%v, %d) = %v  | pass=%v\n", tc.nums, tc.size, got, equal2D(got, tc.want))
	}
}

func equal2D(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
