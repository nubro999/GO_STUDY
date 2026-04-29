# 08. Reverse LinkedList — 정답 및 해설

## 정답 코드 (반복문)

```go
package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func Reverse(head *Node) *Node {
	var prev *Node
	curr := head
	for curr != nil {
		next := curr.Next   // ① 다음 노드 미리 저장
		curr.Next = prev    // ② 현재 노드의 화살표를 뒤로 돌림
		prev = curr         // ③ prev를 한 칸 전진
		curr = next         // ④ curr를 한 칸 전진
	}
	return prev
}
```

## 해설

### 세 포인터 패턴
연결 리스트 뒤집기는 **prev / curr / next** 세 개의 포인터로 진행합니다.

```
초기:     nil   1 -> 2 -> 3 -> nil
         prev  curr

1회 후:   nil <- 1     2 -> 3 -> nil
                prev  curr

2회 후:   nil <- 1 <- 2     3 -> nil
                     prev  curr

3회 후:   nil <- 1 <- 2 <- 3
                          prev (=새 head)
```

### 순서가 매우 중요
```go
next := curr.Next       // 먼저 다음 노드 백업 (안 그러면 잃어버림)
curr.Next = prev        // 화살표 뒤집기
prev = curr             // 그 다음에 prev 이동
curr = next             // 마지막에 curr 이동
```

이 4줄은 외워두는 게 좋습니다. 순서가 바뀌면 다음 노드를 영영 잃어버립니다.

### 빈 리스트 / 단일 노드
- `head == nil`: 루프가 한 번도 실행 안 됨, `prev == nil` 반환 → 정답
- 단일 노드: 한 번 돌고 `prev == 그 노드`, `Next == nil` → 정답

별도 분기 없이 자연스럽게 처리됩니다.

## 다른 풀이: 재귀

```go
func Reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := Reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
```

- 코드는 짧지만 **스택 깊이가 리스트 길이만큼** 쌓임 → 긴 리스트에서 stack overflow.
- Go의 기본 스택은 늘어나지만 시스템 리소스 한계 존재.
- 면접에서 "재귀로도 풀 수 있나요?" 질문에 대비해 알아두면 좋음.

## 자주 하는 실수

| 실수 | 증상 |
|------|------|
| `next := curr.Next` 누락하고 `curr.Next = prev` 먼저 | 다음 노드 잃음 → 무한 루프 또는 nil 참조 |
| `for curr.Next != nil` | 마지막 노드가 뒤집히지 않음 |
| 재귀에서 `head.Next = nil` 누락 | 마지막 노드(=원래 head)가 자기 자신을 가리켜 cycle 형성 |
