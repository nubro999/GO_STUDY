# 04. Anagram — 정답 및 해설

## 정답 코드 (빈도 카운트 방식)

```go
package main

import (
	"fmt"
	"strings"
)

func IsAnagram(a, b string) bool {
	cnt := map[rune]int{}
	for _, r := range strings.ToLower(a) {
		if r == ' ' {
			continue
		}
		cnt[r]++
	}
	for _, r := range strings.ToLower(b) {
		if r == ' ' {
			continue
		}
		cnt[r]--
	}
	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
```

## 해설

### `for _, r := range string`은 rune 단위
Go의 `string`은 UTF-8 인코딩된 byte 시퀀스입니다.

```go
s := "한글"
len(s)        // 6 (byte 길이)
[]byte(s)     // [0xED 0x95 0x9C 0xEA 0xB8 0x80]
for _, r := range s { ... }  // r은 rune (한 글자씩)
```

`range` 루프가 자동으로 UTF-8 디코딩을 해주므로 **한글/이모지 같은 multi-byte 문자도 안전**합니다.

### `rune` = `int32`
`rune`은 Unicode code point를 표현하는 타입(`int32`의 별칭). 문자 단위 처리에는 항상 rune 사용.

### 한 맵으로 +/- 차감하는 트릭
A에서는 +1, B에서는 -1 하면 결국 모든 카운트가 0이어야 애너그램. 두 개의 맵을 비교하는 것보다 코드가 짧습니다.

## 다른 풀이: 정렬

```go
import "sort"

func IsAnagram(a, b string) bool {
	norm := func(s string) string {
		rs := []rune{}
		for _, r := range strings.ToLower(s) {
			if r != ' ' {
				rs = append(rs, r)
			}
		}
		sort.Slice(rs, func(i, j int) bool { return rs[i] < rs[j] })
		return string(rs)
	}
	return norm(a) == norm(b)
}
```

- 시간 복잡도 O(n log n) (정렬 vs 빈도 카운트의 O(n)).
- 코드는 더 짧고 한 번 보면 직관적.
- ASCII만 다룬다면 `[128]int` 배열로 더 빠르게 가능.

## 자주 하는 실수

| 실수 | 증상 |
|------|------|
| `for i := 0; i < len(s); i++ { c := s[i] }` | byte 단위 — 한글/이모지 깨짐 |
| 대소문자 변환 누락 | "Listen"과 "Silent" → false |
| 길이 비교 누락 (정렬 풀이) | `"a"` vs `"aa"` 같은 부분문자열 케이스에서 버그 |
