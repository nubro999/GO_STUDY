# 03. Word Count — 정답 및 해설

## 정답 코드

```go
package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	out := map[string]int{}
	for _, w := range strings.Fields(strings.ToLower(s)) {
		out[w]++
	}
	return out
}
```

## 해설

### 핵심 관용구: `m[k]++`
Go 맵의 zero value는 자료형의 zero value입니다. `int` 맵에서 존재하지 않는 키를 읽으면 `0`이 반환되므로 `out[w]++`은 **"없으면 1, 있으면 +1"** 이 자연스럽게 됩니다.

```go
out := map[string]int{}   // 빈 맵
out["go"]++               // out["go"] == 1
out["go"]++               // out["go"] == 2
```

> ⚠️ `var out map[string]int` (nil 맵)에 쓰기 시도하면 **panic**. 반드시 `make` 또는 리터럴로 초기화.

### `strings.Fields` vs `strings.Split`
- `strings.Split(s, " ")` — 정확히 한 칸 공백으로만 분리. 연속 공백, 탭, 개행 → 빈 토큰 발생.
- `strings.Fields(s)` — **모든 종류의 공백(스페이스/탭/개행)에서 분리**, 빈 토큰 자동 제거.

문장 분할에는 거의 항상 `Fields`가 정답입니다.

### 대소문자 무시
`strings.ToLower(s)` 를 한 번 호출해 전체를 소문자로. 각 단어마다 호출하는 것보다 효율적.

## 자주 하는 실수

| 실수 | 증상 |
|------|------|
| `var out map[string]int` 후 `out[k]++` | runtime panic |
| `strings.Split(s, " ")` 사용 | 연속 공백 시 `""` 키가 카운트됨 |
| 대소문자 변환 누락 | "Go"와 "go"가 별개 키 |

## 심화: 문장 부호 제거

```go
import "unicode"

clean := strings.Map(func(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsSpace(r) {
		return r
	}
	return -1  // 제거
}, s)
```

`strings.Map`은 각 rune을 변환하는 함수형 도구입니다. `-1` 반환 = 해당 rune 삭제.
