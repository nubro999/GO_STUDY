# 08-extra1. Merge Sorted Lists — 정답 및 해설

## 정답 코드

```go
func MergeSorted(a, b *Node) *Node {
	dummy := &Node{}
	tail := dummy

	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}
	// 남은 한쪽을 통째로 붙임
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}

	return dummy.Next
}
```

## 해설

### 같은 패턴 — `*Node` 포인터 조작
Reverse 와 동일하게 포인터 변수를 잡고 `Next` 를 재배선합니다.

### 핵심 idiom: Dummy Head
```go
dummy := &Node{}   // 가상의 head 노드
tail := dummy
```

**왜 dummy를 쓰는가?**
- "첫 번째 노드는 어떻게 처리하지?" 라는 분기 분리됨.
- `tail.Next = ...` 만으로 모든 추가 동작이 통일.
- 마지막에 `dummy.Next` 가 진짜 head.

dummy 패턴은 연결 리스트 문제에서 **거의 항상** 등장합니다. 외워두세요.

### 왜 `<=` 인가?
```go
if a.Val <= b.Val { ... }
```
같은 값일 때 `a` 를 먼저 넣음 — 안정 병합(stable merge). 정렬 안정성이 필요한 경우 차이가 납니다.

### 남은 부분 통째로 연결
```go
if a != nil {
    tail.Next = a
}
```

이미 정렬되어 있으므로 한쪽이 끝나면 나머지는 그대로 이어붙이면 끝. 모든 노드를 일일이 옮길 필요 없음.

### 시간/공간 복잡도
- 시간 O(m + n)
- 공간 O(1) — 새 노드 할당 없음 (dummy 하나만)

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| dummy 없이 첫 노드 분기 처리 | 코드가 길어지고 버그 위험 ↑ |
| 한쪽 끝났을 때 남은 부분 연결 누락 | 결과 리스트가 짧음 |
| `tail = tail.Next` 누락 | 매번 dummy.Next 만 갱신 → 마지막 노드만 연결됨 |
