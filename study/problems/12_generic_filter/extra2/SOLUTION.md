# 12-extra2. Generic Min / Max — 정답 및 해설

## 정답 코드

```go
import "cmp"

func Min[T cmp.Ordered](xs []T) (T, bool) {
	var zero T
	if len(xs) == 0 {
		return zero, false
	}
	m := xs[0]
	for _, v := range xs[1:] {
		if v < m {
			m = v
		}
	}
	return m, true
}

func Max[T cmp.Ordered](xs []T) (T, bool) {
	var zero T
	if len(xs) == 0 {
		return zero, false
	}
	m := xs[0]
	for _, v := range xs[1:] {
		if v > m {
			m = v
		}
	}
	return m, true
}
```

## 해설

### 같은 패턴 — Map/Filter/Reduce 의 특수화
함수 골격은 Reduce 와 같음 — 슬라이스 누적. 다만 제약이 `any` 가 아니라 `cmp.Ordered` 라 `<`, `>` 연산자 사용 가능.

### `cmp.Ordered` (Go 1.21+)
```go
package cmp

type Ordered interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~string
}
```

- `<`, `<=`, `>`, `>=` 가 가능한 타입.
- `~int` 의 `~` 는 **"int을 underlying type으로 가지는 모든 타입"** — 예: `type MyInt int` 도 포함.
- `cmp` 패키지엔 `Compare`, `Less` 도 있음.

### Comparable vs Ordered 구분

| 제약 | 가능한 연산 | 예시 |
|------|-------------|------|
| `comparable` | `==`, `!=` | map 키, struct 비교 |
| `cmp.Ordered` | `<`, `>` 등 | 정렬, min/max |
| `any` | 모두 불가 (메서드 호출만) | Map/Filter |

### 빈 슬라이스 처리: ok 패턴
```go
return zero, false
```

- `var zero T` 는 타입 T 의 zero value (int면 0, string이면 "", struct면 모든 필드가 zero).
- (값, bool) 반환은 Go 의 일반적 패턴 — comma-ok idiom 의 동기.
- 호출자는 `if v, ok := Min(xs); ok { ... }` 로 처리.

### Go 1.21 내장 `min`/`max`
실은 Go 1.21부터 **언어 자체의 내장 함수** `min(a, b)`, `max(a, b)` 가 있습니다.

```go
m := min(3, 1, 4, 1, 5)   // 가변 인자 OK
```

하지만 슬라이스를 받는 형태는 표준에 없어, 본 문제처럼 직접 작성해야 함.
또는 `slices.Min` / `slices.Max` (Go 1.21+) 가 정확히 같은 일을 함:

```go
import "slices"
v := slices.Min(xs)   // 빈 슬라이스면 panic
```

표준 함수는 panic — "Min" 의 의미상 빈 입력이 정의되지 않으니 명시적 에러 처리.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `T any` 제약 | `<` 연산자 사용 불가 → 컴파일 에러 |
| `var m T` 로 시작 후 `if v < m` | T가 정수면 0이 초기값이라 양수만 있는 슬라이스에서 0 반환 (max) — 첫 원소를 시작값으로 |
| 빈 슬라이스에서 panic | (T, bool) 패턴 또는 명시적 panic 으로 일관된 정책 |
| Go 1.20 이하에서 `cmp.Ordered` 사용 | 컴파일 에러 — `golang.org/x/exp/constraints.Ordered` 사용 |
