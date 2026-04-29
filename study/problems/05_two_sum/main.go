// 문제 05: Two Sum
//
// [학습 포인트]
//   - map을 활용한 O(n) 알고리즘 패턴
//   - "이미 본 값" 캐시로 단일 패스 풀이
//   - 다중 반환값
//
// [문제]
//   정수 슬라이스 nums와 정수 target이 주어진다.
//   합이 target이 되는 두 인덱스 (i, j)를 반환하시오. (i < j)
//   답이 존재하지 않으면 (-1, -1) 반환.
//   각 입력에는 정답이 최대 1개라고 가정.
//
// [예시]
//   TwoSum([2,7,11,15], 9)  -> (0, 1)   // 2 + 7 = 9
//   TwoSum([3,2,4], 6)      -> (1, 2)   // 2 + 4 = 6
//   TwoSum([3,3], 6)        -> (0, 1)
//   TwoSum([1,2,3], 100)    -> (-1, -1)
//
// [힌트]
//   nums[i]를 순회하며 (target - nums[i]) 가 이전에 본 값인지 map에서 확인.
package main

import "fmt"

func TwoSum(nums []int, target int) (int, int) {
	// TODO: 구현하세요.
	return -1, -1
}

func main() {
	tests := []struct {
		nums   []int
		target int
		wantI  int
		wantJ  int
	}{
		{[]int{2, 7, 11, 15}, 9, 0, 1},
		{[]int{3, 2, 4}, 6, 1, 2},
		{[]int{3, 3}, 6, 0, 1},
		{[]int{1, 2, 3}, 100, -1, -1},
	}

	for _, tc := range tests {
		i, j := TwoSum(tc.nums, tc.target)
		pass := i == tc.wantI && j == tc.wantJ
		fmt.Printf("TwoSum(%v, %d) = (%d, %d)  | pass=%v\n", tc.nums, tc.target, i, j, pass)
	}
}
