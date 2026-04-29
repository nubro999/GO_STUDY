# 04-extra1. Palindrome — 정답 및 해설

## 정답 코드 (투포인터)

```go
import "unicode"

func IsPalindrome(s string) bool {
	rs := []rune{}
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			rs = append(rs, unicode.ToLower(r))
		}
	}
	i, j := 0, len(rs)-1
	for i < j {
		if rs[i] != rs[j] {
			return false
		}
		i++
		j--
	}
	return true
}
```

## 해설

### 같은 패턴 — `[]rune` + 정규화
Anagram 문제와 동일하게 `range string` 으로 rune 단위 순회.

### `unicode` 패키지의 도우미
- `unicode.IsLetter(r)` — 글자(국제 문자 포함)
- `unicode.IsDigit(r)` — 숫자
- `unicode.IsSpace(r)` — 공백
- `unicode.ToLower(r)` — 소문자 변환

`strings.ToLower` 는 문자열 전체를, `unicode.ToLower` 는 단일 rune을 변환. rune 처리할 때는 후자를 씁니다.

### 투포인터 (two pointers)
```go
i, j := 0, len(rs)-1
for i < j {
    if rs[i] != rs[j] { return false }
    i++; j--
}
```
- 문자열 양 끝에서 시작해 가운데로 좁히는 패턴.
- 회문, 정렬된 배열 합 검색, 슬라이딩 윈도우 등에서 자주 등장.
- 공간 O(n) (rune 슬라이스), 시간 O(n).

### 더 짧게: 슬라이스 한 번 만들기 + reverse 비교
```go
rev := make([]rune, len(rs))
for i, r := range rs {
    rev[len(rs)-1-i] = r
}
return string(rs) == string(rev)
```
가독성은 좋지만 메모리 2배.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `s[i] != s[len(s)-1-i]` byte 비교 | multi-byte 문자(한글) 깨짐 |
| 정규화 누락 | "A man" 의 'A'와 'n' 비교 → false |
| `i <= j` 로 비교 | 가운데 글자가 자기 자신과 비교 (영향 없지만 불필요) |
