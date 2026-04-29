# 08-extra2. Middle Node (Slow/Fast Pointer) — 정답 및 해설

## 정답 코드

```go
func Middle(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
```

## 해설

### 같은 패턴 — 두 포인터 동시 진행
Reverse 의 `prev/curr` 처럼 두 포인터를 동시에 다룹니다. 이번엔 **속도 차이** 가 핵심.

### Slow/Fast (= Tortoise/Hare) 알고리즘
```
1 -> 2 -> 3 -> 4 -> 5

초기:    slow=1, fast=1
1회:     slow=2, fast=3
2회:     slow=3, fast=5  → fast.Next == nil 이라 종료
결과:    slow = 3 (정확히 중간!)
```

`fast` 가 두 배 빠르므로 `fast` 가 끝에 도달했을 때 `slow` 는 정확히 중간에 있게 됩니다.

### 짝수 길이 처리
- 노드 4개: `1->2->3->4`
- slow가 가리키는 위치: 3 (두 번째 중간)
- "첫 번째 중간" 을 원하면 종료 조건을 `fast.Next.Next != nil` 로 변경.

### 종료 조건의 의미
```go
for fast != nil && fast.Next != nil
```
- `fast == nil` 이면 더 이상 진행 불가
- `fast.Next == nil` 이면 한 칸 더 가면 nil → 두 칸 진행 불가
- 두 조건의 순서 중요: `fast` 먼저 체크해야 nil 역참조 panic 방지

### 다른 응용 — Cycle 감지
같은 slow/fast 패턴으로 **순환(cycle) 감지** 가능:
```go
for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    if slow == fast {
        return true   // 만나면 cycle 존재
    }
}
return false
```
이걸 **Floyd's cycle detection** 알고리즘이라 부릅니다.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `fast.Next != nil` 만 체크 | 빈 리스트(`fast == nil`)에서 nil 역참조 panic |
| `fast = fast.Next` (한 칸씩) | slow와 같은 속도라 끝에 도달 시 slow가 끝에 있음 |
| `for fast.Next.Next != nil` | 짝수/홀수 처리가 미묘하게 다름 — 의도에 따라 선택 |
