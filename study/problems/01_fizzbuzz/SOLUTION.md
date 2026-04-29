# 01. FizzBuzz — 정답 및 해설

## 정답 코드

```go
package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) []string {
	out := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			out = append(out, "FizzBuzz")
		case i%3 == 0:
			out = append(out, "Fizz")
		case i%5 == 0:
			out = append(out, "Buzz")
		default:
			out = append(out, strconv.Itoa(i))
		}
	}
	return out
}
```

## 해설

### 핵심 패턴: `switch` true-form
조건식 없는 `switch { case <bool>: ... }` 형태는 **첫 번째로 true가 되는 case**가 실행됩니다. `if/else if` 사다리보다 가독성이 좋습니다.

### 검사 순서가 중요
`i%15`(공배수)를 **먼저** 검사해야 합니다. `i%3`을 먼저 두면 15는 "Fizz"로 분류됩니다. switch는 위에서부터 매칭하므로 더 강한 조건(=더 좁은 조건)을 위에 둡니다.

### 용량 미리 할당
`make([]string, 0, n)` — 길이 0, 용량 n. n번 `append`해도 재할당이 일어나지 않아 효율적입니다. 슬라이스의 길이/용량 차이는 Go의 핵심 개념입니다.

### 정수 → 문자열
- `strconv.Itoa(i)` — 정수 전용, 빠름.
- `fmt.Sprint(i)` — 임의 타입에 동작하지만 reflection 비용.
코딩테스트에서는 `strconv` 쪽을 선호합니다.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `i%3` 케이스를 위에 둠 | 15가 "Fizz"로 분류됨 |
| `make([]string, n)` 후 `append` | 길이 n인 빈 슬라이스 뒤에 append → 길이 2n |
| `for i := 0; i < n` | 0~n-1 출력, 문제는 1~n |
