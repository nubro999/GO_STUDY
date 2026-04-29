# 04-extra2. RLE Compress — 정답 및 해설

## 정답 코드

```go
import (
	"fmt"
	"strings"
)

func Compress(s string) string {
	if s == "" {
		return ""
	}
	rs := []rune(s)

	var b strings.Builder
	count := 1
	for i := 1; i <= len(rs); i++ {
		if i == len(rs) || rs[i] != rs[i-1] {
			fmt.Fprintf(&b, "%c%d", rs[i-1], count)
			count = 1
		} else {
			count++
		}
	}

	if b.Len() >= len(s) {
		return s
	}
	return b.String()
}
```

## 해설

### 같은 패턴 — `[]rune` 변환 + 순회
Anagram 과 동일하게 multi-byte 안전 처리를 위해 `[]rune` 으로 변환.

### 핵심 idiom: `strings.Builder`
```go
var b strings.Builder
b.WriteString("hello")
b.WriteByte('!')
fmt.Fprintf(&b, "x=%d", 42)
result := b.String()
```

`+=` 로 문자열 연결하면 매번 새 문자열을 할당하지만 (O(n²)),
`strings.Builder` 는 내부 buffer에 누적해 마지막에 한 번만 string 화 (O(n)).

코딩테스트 / 실무 모두에서 **반복적인 문자열 조립의 표준** 입니다.

### `fmt.Fprintf(&b, ...)`
`Fprintf` 는 첫 인자로 `io.Writer` 를 받는 fprintf 변형. `strings.Builder` 가 `Write([]byte)` 메서드를 가지므로 `io.Writer` 인터페이스를 만족 → 직접 넘길 수 있음.

### 경계 처리 트릭: `i <= len(rs)`
루프를 한 칸 더 돌리고 `i == len(rs)` 일 때 마지막 그룹을 flush. "마지막 글자 처리 누락" 버그를 방지하는 일반 패턴.

### 길이 비교 시 `b.Len()` vs `len(b.String())`
- `b.Len()` — O(1), 내부 byte 길이 즉시 반환
- `len(b.String())` — 문자열 변환 비용 발생 (단, Go 컴파일러가 최적화하기도 함)

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `s += fmt.Sprintf(...)` 누적 | O(n²), 큰 입력에서 매우 느림 |
| 마지막 그룹 flush 누락 | "aabbb" → "a2" 로 끝남 |
| 길이 비교를 rune 수 vs byte 수로 섞음 | 압축 여부 판정 오류 (ASCII만 다루면 무관) |
