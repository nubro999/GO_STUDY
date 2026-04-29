// 문제 02: 슬라이스 회전 (Rotate)
//
// [학습 포인트]
//   - slice 자료형의 구조 (포인터 + 길이 + 용량)
//   - slicing 표현 a[i:j], append, copy
//   - 모듈러 연산을 활용한 회전 알고리즘
//
// [문제]
//   정수 슬라이스 nums와 정수 k가 주어진다.
//   nums의 원소들을 오른쪽으로 k칸 회전시킨 새 슬라이스를 반환하시오.
//   원본 슬라이스를 변경하지 않아야 한다.
//
//   k가 len(nums)보다 클 수 있고, 음수일 수도 있음에 주의.
//
// [예시]
//   Rotate([1,2,3,4,5], 2)  -> [4,5,1,2,3]
//   Rotate([1,2,3,4,5], 7)  -> [4,5,1,2,3]   // 7 % 5 == 2
//   Rotate([1,2,3,4,5], -1) -> [2,3,4,5,1]
package main

import "fmt"

func Rotate(nums []int, k int) []int {
	// TODO: 구현하세요.
	// 힌트: k = ((k % n) + n) % n 으로 정규화하면 음수도 처리 가능.
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	k = ((k % n) + n) % n // k를 0 이상 n 미만으로 정규화
	out := make([]int, n)
	for i, v := range nums {
		out[(i+k)%n] = v
	}

	return out
}

func main() {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 7, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, -1, []int{2, 3, 4, 5, 1}},
		{[]int{1}, 100, []int{1}},
	}

	for _, tc := range tests {
		original := append([]int(nil), tc.nums...)
		got := Rotate(tc.nums, tc.k)
		mutated := !equalInt(original, tc.nums)
		fmt.Printf("Rotate(%v, %d) = %v  | pass=%v  | mutated=%v\n",
			original, tc.k, got, equalInt(got, tc.want), mutated)
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
