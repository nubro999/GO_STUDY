# 09-extra2. Chain Depth — 정답 및 해설

## 정답 코드

```go
func ChainDepth(err error) int {
	depth := 0
	for err != nil {
		depth++
		err = errors.Unwrap(err)
	}
	return depth
}
```

## 해설

### 같은 패턴 — `%w` 래핑의 안쪽 들여다보기
SafeDivide 에서는 `%w` 로 에러를 래핑했고, `errors.As` 로 체인을 탐색했습니다. 이번엔 직접 `errors.Unwrap` 으로 한 단계씩 내려갑니다.

### `errors.Unwrap` 의 동작
```go
err := fmt.Errorf("outer: %w", inner)
errors.Unwrap(err)   // inner 를 반환

err2 := errors.New("plain")
errors.Unwrap(err2)  // nil — 더 이상 unwrap 할 게 없음
```

`fmt.Errorf("%w", ...)` 로 만든 에러나 `Unwrap() error` 메서드를 가진 커스텀 타입에 대해 작동.

### `%w` vs `%v` vs `%s`

| verb | 동작 |
|------|------|
| `%w` | 에러를 **래핑** (체인 형성, `errors.Is/As` 가 통과) |
| `%v` | 단순 문자열 변환, 체인 안 만듦 |
| `%s` | `%v` 와 거의 동일 (string 포맷) |

```go
e1 := fmt.Errorf("a: %w", inner)  // 체인 깊이 2
e2 := fmt.Errorf("a: %v", inner)  // 체인 깊이 1 — inner 정보는 메시지로만 들어감
```

### 다중 래핑 (Go 1.20+)
```go
e := fmt.Errorf("multi: %w and %w", err1, err2)
```
`%w` 를 여러 번 사용 가능. 이런 에러의 `Unwrap` 는 `[]error` 를 반환하는 시그니처를 갖습니다:
```go
type multiError interface {
    Unwrap() []error
}
```
본 문제의 단순 `errors.Unwrap` 은 이 케이스를 다루지 않음. `errors.Is/As` 는 다중 래핑도 처리.

### 응용 — 가장 안쪽 에러 추출
```go
func RootCause(err error) error {
    for {
        next := errors.Unwrap(err)
        if next == nil {
            return err
        }
        err = next
    }
}
```

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `for err != nil; err = errors.Unwrap(err) { depth++ }` 같은 형태 | Go에는 C-style for의 `init; cond; post` 만 있음 — for 헤더 문법 확인 |
| `%w` 와 `%v` 혼용 | 의도와 다른 깊이가 나옴 |
| 무한 루프 가능성 | 일반적인 `%w` 체인은 유한이지만, 커스텀 `Unwrap()` 이 자기 자신을 반환하면 무한 루프 (실수로 만들기 어려움) |
