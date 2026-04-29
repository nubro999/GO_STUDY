// 문제 08-extra2: 연결 리스트의 중간 노드 찾기 (slow/fast pointer)
//
// [같은 패턴]
//   - 두 포인터가 다른 속도로 진행, *Node 순회
//
// [문제]
//   단일 연결 리스트의 head가 주어졌을 때 중간 노드를 반환.
//   노드 수가 짝수면 두 번째 중간(=오른쪽)을 반환.
//   빈 리스트는 nil.
//
// [예시]
//   1 -> 2 -> 3 -> 4 -> 5         → 3
//   1 -> 2 -> 3 -> 4              → 3 (두 번째 중간)
//   1                             → 1
//   nil                           → nil
//
// [힌트]
//   slow는 한 칸씩, fast는 두 칸씩 진행:
//     for fast != nil && fast.Next != nil {
//         slow = slow.Next
//         fast = fast.Next.Next
//     }
//   fast 가 끝에 도달하면 slow는 중간.
package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func Middle(head *Node) *Node {
	// TODO: 구현하세요.
	return nil
}

func fromSlice(vals []int) *Node {
	var head *Node
	for i := len(vals) - 1; i >= 0; i-- {
		head = &Node{Val: vals[i], Next: head}
	}
	return head
}

func main() {
	tests := []struct {
		input []int
		want  int  // 찾고자 하는 중간 노드의 값. 빈 리스트는 -1로 표시
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4}, 3},
		{[]int{1}, 1},
		{[]int{}, -1},
		{[]int{1, 2}, 2},
	}

	for _, tc := range tests {
		mid := Middle(fromSlice(tc.input))
		var got int
		if mid == nil {
			got = -1
		} else {
			got = mid.Val
		}
		fmt.Printf("Middle(%v) = %d  | pass=%v\n", tc.input, got, got == tc.want)
	}
}
