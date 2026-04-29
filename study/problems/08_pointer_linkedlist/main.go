// 문제 08: 단일 연결 리스트 뒤집기
//
// [학습 포인트]
//   - 자기참조 구조체와 *T 포인터
//   - nil 체크와 포인터 변수 재할당
//   - 반복문 vs 재귀로 푸는 방식 비교
//
// [문제]
//   단일 연결 리스트의 head가 주어졌을 때, 리스트를 뒤집은 새 head를 반환하시오.
//   in-place(추가 노드 할당 없이) 풀이를 권장.
//
//   type Node struct {
//       Val  int
//       Next *Node
//   }
//
// [예시]
//   1 -> 2 -> 3 -> nil   ===>   3 -> 2 -> 1 -> nil
//
// [힌트]
//   prev := *Node(nil); curr := head;
//   while curr != nil:
//       next := curr.Next
//       curr.Next = prev
//       prev = curr
//       curr = next
//   return prev
package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func Reverse(head *Node) *Node {
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

func toSlice(head *Node) []int {
	var out []int
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func main() {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, nil},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, tc := range tests {
		head := fromSlice(tc.input)
		got := toSlice(Reverse(head))
		fmt.Printf("Reverse(%v) = %v  | pass=%v\n", tc.input, got, equalInt(got, tc.want))
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
