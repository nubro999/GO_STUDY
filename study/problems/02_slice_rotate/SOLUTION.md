# 02. Slice Rotate — 정답 및 해설

## 정답 코드

```go
package main

import "fmt"

func Rotate(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	// 음수와 큰 값을 [0, n) 범위로 정규화
	k = ((k % n) + n) % n
	out := make([]int, n)
	for i, v := range nums {
		out[(i+k)%n] = v
	}
	return out
}
```

## 해설

### 핵심: 인덱스 매핑
"오른쪽으로 k칸 회전" = 원본 i번째 원소가 새 슬라이스의 `(i+k)%n` 위치로 이동.

```
원본:    [1, 2, 3, 4, 5]   k=2
인덱스:   0  1  2  3  4
이동후:  (i+2)%5 = 2 3 4 0 1
새:      [4, 5, 1, 2, 3]
```

### 음수 모듈러 함정
Go의 `%`는 피연산자 부호를 따릅니다 — `-1 % 5 == -1` (Python처럼 4가 아님).
`((k % n) + n) % n` 패턴은 음수 결과를 양수 범위로 끌어올리는 **표준 관용구**입니다.

### 원본 보존
문제 조건: 원본 슬라이스를 변경하지 않아야 함. 새 슬라이스 `make([]int, n)` 에 쓰면 됩니다. 만약 in-place 회전을 원한다면 **3-reverse 알고리즘**:
1. 전체 reverse
2. 앞쪽 k개 reverse
3. 나머지 reverse

이러면 추가 메모리 O(1)로 회전 완료.

## 다른 풀이: append + slicing

```go
func Rotate(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	k = ((k % n) + n) % n
	out := make([]int, 0, n)
	out = append(out, nums[n-k:]...)  // 뒤쪽 k개를 앞으로
	out = append(out, nums[:n-k]...)  // 앞쪽 n-k개를 뒤로
	return out
}
```
직관적이지만 슬라이싱 인덱스 헷갈리기 쉬움. 인덱스 매핑 방식이 더 안전합니다.

## 자주 하는 실수

- `k % n` 만 사용 → k가 음수일 때 잘못된 인덱스 (panic 가능)
- `out := nums` 후 수정 → 슬라이스는 reference type이라 원본도 변경됨
- `len(nums) == 0` 체크 누락 → `k % 0` 으로 panic
