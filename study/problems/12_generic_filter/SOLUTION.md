# 12. Generics (Map/Filter/Reduce) — 정답 및 해설

## 정답 코드

```go
package main

import "fmt"

func Map[T, U any](xs []T, f func(T) U) []U {
	out := make([]U, len(xs))
	for i, x := range xs {
		out[i] = f(x)
	}
	return out
}

func Filter[T any](xs []T, pred func(T) bool) []T {
	out := []T{}
	for _, x := range xs {
		if pred(x) {
			out = append(out, x)
		}
	}
	return out
}

func Reduce[T, U any](xs []T, init U, f func(U, T) U) U {
	acc := init
	for _, x := range xs {
		acc = f(acc, x)
	}
	return acc
}

func main() {
	squares := Map([]int{1, 2, 3}, func(x int) int { return x * x })
	fmt.Println(squares) // [1 4 9]

	labels := Map([]int{1, 2, 3}, func(x int) string { return fmt.Sprintf("#%d", x) })
	fmt.Println(labels) // [#1 #2 #3]

	evens := Filter([]int{1, 2, 3, 4, 5}, func(x int) bool { return x%2 == 0 })
	fmt.Println(evens) // [2 4]

	sum := Reduce([]int{1, 2, 3, 4}, 0, func(acc, x int) int { return acc + x })
	fmt.Println(sum) // 10
}
```

## 해설

### 타입 파라미터 문법
```go
func Map[T, U any](xs []T, f func(T) U) []U
//        ^^^^^^^^
//        대괄호 안에 타입 파라미터 선언
```

- `T`, `U`: 타입 파라미터 이름 (관례적으로 한 글자 대문자)
- `any`: 제약 조건 (`interface{}` 의 별칭, 모든 타입 허용)
- 반환 타입 `[]U` 에서 같은 이름을 재사용

### 제약 (constraint)

```go
func Sum[T int | float64](xs []T) T {
    var s T
    for _, x := range xs { s += x }
    return s
}
```

- `int | float64` → 타입 합집합 (union)
- `comparable` → `==`, `!=` 가능한 타입
- `cmp.Ordered` (Go 1.21+) → `<`, `>` 가능 (숫자, 문자열, ...)
- 직접 정의:
```go
type Number interface {
    int | int64 | float64
}
```

### 두 타입 파라미터의 의미
`Map[T, U]` 에서 T와 U를 **분리**한 이유는, 변환 함수의 입력과 출력 타입이 다를 수 있기 때문입니다.

```go
Map([]int{1,2,3}, strconv.Itoa)        // []int → []string
Map([]User{...}, func(u User) string { return u.Name })  // []User → []string
```

만약 `Map[T any](xs []T, f func(T) T) []T` 라면 같은 타입으로만 변환 가능 — 활용도가 크게 줄어듭니다.

### `Reduce` 의 시그니처가 까다로운 이유
```go
func Reduce[T, U any](xs []T, init U, f func(U, T) U) U
```

- `T`: 입력 원소 타입
- `U`: 누적값(accumulator) 타입
- f는 `(누적값, 원소) → 새 누적값`

T와 U가 다를 수 있어야 다음 같은 풀이가 가능:
```go
// []string → map[string]int (단어 빈도)
counts := Reduce(words, map[string]int{}, func(acc map[string]int, w string) map[string]int {
    acc[w]++
    return acc
})
```

### 타입 추론
Go는 인자에서 타입 파라미터를 추론합니다 — 명시 호출은 거의 필요 없음.

```go
Map([]int{1,2,3}, func(x int) int { return x*x })       // T=int, U=int 자동
Map[int, string]([]int{1,2,3}, strconv.Itoa)            // 명시 (보통 불필요)
```

추론이 실패하는 경우(예: 빈 슬라이스 + 모호한 클로저)에만 명시.

## 표준 라이브러리 활용 (Go 1.21+)

```go
import "slices"

slices.Index(xs, target)
slices.Contains(xs, x)
slices.Sort(xs)
slices.SortFunc(xs, cmp)
slices.IndexFunc(xs, pred)
```

`Map/Filter/Reduce`는 표준에 들어오지 않았지만, `slices` 패키지에 자주 쓰는 도우미가 많이 있습니다. 대부분의 코드에서는 직접 for 루프가 더 명확하다는 게 Go 커뮤니티의 일반적인 입장.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `func Map[T any](xs []T, f func(T) T)` (한 타입) | int → string 변환 불가 |
| `make([]U, 0)` 후 `out[i] = f(x)` | 길이 0 슬라이스 인덱스 접근 → panic. `make([]U, len(xs))` 이어야 |
| 제약 없이 산술 시도 (`var s T; s += x`) | `any` 에는 + 연산자 없음 → 컴파일 에러. 제약 필요 |
| 너무 일반적인 시그니처에 집착 | 단순 for 루프가 더 읽기 좋을 때가 많음 |
