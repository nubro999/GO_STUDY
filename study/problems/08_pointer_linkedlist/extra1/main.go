// 문제 08-extra1: 두 정렬된 연결 리스트 합치기
//
// [같은 패턴]
//   - *Node 포인터 조작, dummy head 트릭
//
// [문제]
//   두 개의 오름차순 정렬된 연결 리스트 a, b가 주어진다.
//   둘을 병합한 단일 정렬된 연결 리스트의 head를 반환하시오.
//
//   입력 노드를 그대로 재사용해도 좋다 (새 노드 할당 불필요).
//
// [예시]
//   a: 1 -> 3 -> 5
//   b: 2 -> 4 -> 6
//   →  1 -> 2 -> 3 -> 4 -> 5 -> 6
//
// [힌트]
//   - dummy head 트릭: tail 포인터를 가리키는 가상의 head 노드를 만들어두면
//     첫 노드 분기 처리가 사라짐.
//   - 한 쪽이 nil이 되면 나머지를 통째로 붙임.
package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func MergeSorted(a, b *Node) *Node {
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
		a, b []int
		want []int
	}{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{}, []int{1, 2}, []int{1, 2}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{[]int{}, []int{}, nil},
		{[]int{1, 1, 2}, []int{1, 3}, []int{1, 1, 1, 2, 3}},
	}

	for _, tc := range tests {
		got := toSlice(MergeSorted(fromSlice(tc.a), fromSlice(tc.b)))
		fmt.Printf("Merge(%v, %v) = %v  | pass=%v\n", tc.a, tc.b, got, equalInt(got, tc.want))
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
