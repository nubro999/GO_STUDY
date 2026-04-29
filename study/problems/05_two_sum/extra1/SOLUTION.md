# 05-extra1. Longest Unique Substring — 정답 및 해설

## 정답 코드

```go
func LongestUnique(s string) int {
	last := map[byte]int{} // 문자 -> 마지막으로 본 인덱스
	left, best := 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if idx, ok := last[c]; ok && idx >= left {
			left = idx + 1
		}
		last[c] = i
		if i-left+1 > best {
			best = i - left + 1
		}
	}
	return best
}
```

## 해설

### 같은 패턴 — TwoSum 의 "이전에 본 값 캐싱"
TwoSum 은 "값 → 인덱스" 맵을 유지했고, 본 문제는 "문자 → 마지막 인덱스" 맵을 유지합니다. 둘 다 **단일 패스 + map** 패턴.

### 슬라이딩 윈도우
```
abcabcbb
^^^         left=0, i=0..2 → "abc" (길이 3)
   ^^^      a 중복 발견 → left=1, "bca"
      ^^^   b 중복 발견 → left=2, "cab"
         ...
```

핵심 관계식:
- 윈도우 길이 = `i - left + 1`
- 중복 발견 시 `left = (마지막 등장 위치) + 1`

### 조건의 핵심: `idx >= left`
```go
if idx, ok := last[c]; ok && idx >= left { ... }
```

`idx < left` 면 그 문자는 **현재 윈도우 밖**에 있으므로 무시. 이 조건 없으면 `left` 가 거꾸로 움직여 윈도우가 깨집니다.

### 시간/공간 복잡도
- 시간 O(n) — 각 문자 한 번씩
- 공간 O(min(n, σ)) — σ 는 문자 종류 수 (ASCII면 128)

### 한글/multi-byte 대응
`s[i]` 는 byte. multi-byte 문자가 있으면 깨지므로 `[]rune(s)` 으로 변환 후 동일 로직.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `idx >= left` 체크 누락 | "abba" 에서 left가 거꾸로 → 잘못된 길이 |
| 매 반복마다 `for i := left; i < ...` 중첩 검사 | O(n²), 슬라이딩 윈도우 의미 상실 |
| `best = max(best, i-left+1)` 갱신 누락 | 0 반환 |
