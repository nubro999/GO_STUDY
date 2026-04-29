# 09. Custom Error & Wrapping — 정답 및 해설

## 정답 코드

```go
package main

import (
	"errors"
	"fmt"
)

type DivisionByZeroError struct {
	Numerator int
}

func (e *DivisionByZeroError) Error() string {
	return fmt.Sprintf("cannot divide %d by zero", e.Numerator)
}

func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionByZeroError{Numerator: a}
	}
	return a / b, nil
}

func ComputeAverage(nums []int, divisor int) (int, error) {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	avg, err := SafeDivide(sum, divisor)
	if err != nil {
		return 0, fmt.Errorf("compute average failed: %w", err)
	}
	return avg, nil
}

func main() {
	if v, err := SafeDivide(10, 2); err == nil {
		fmt.Printf("SafeDivide(10, 2) = %d  | pass=%v\n", v, v == 5)
	}

	_, err := ComputeAverage([]int{1, 2, 3, 4}, 0)
	fmt.Printf("ComputeAverage error: %v\n", err)

	var dze *DivisionByZeroError
	if errors.As(err, &dze) {
		fmt.Printf("  -> 분자값 추출 성공: %d  | pass=%v\n", dze.Numerator, dze.Numerator == 10)
	}
}
```

## 해설

### `error` 인터페이스
```go
type error interface {
	Error() string
}
```
딱 하나의 메서드. `Error() string` 을 가진 모든 타입은 자동으로 error.

### 포인터 리시버를 쓰는 이유
```go
func (e *DivisionByZeroError) Error() string  // 권장
```
Go 커뮤니티 컨벤션입니다. 이유:
1. **identity 비교**: `errors.Is` 가 포인터 동등성을 체크할 때 명확.
2. **메모리 효율**: 에러는 보통 자주 생성되지 않으므로 큰 차이는 없지만 일관성을 위해.
3. **`errors.As` 호환**: 타겟이 `**T` 형태여야 하므로 포인터 타입이 자연스러움.

### `fmt.Errorf("%w", err)` — 에러 래핑
`%w` verb는 Go 1.13+에서 추가된 특별 verb로, **에러 체인**을 만듭니다.

```go
err1 := errors.New("disk full")
err2 := fmt.Errorf("save failed: %w", err1)
err3 := fmt.Errorf("request handler: %w", err2)

// err3 메시지: "request handler: save failed: disk full"
// 체인:        err3 -> err2 -> err1
```

### `errors.Is` vs `errors.As`

| 함수 | 용도 | 예시 |
|------|------|------|
| `errors.Is(err, target)` | **특정 sentinel error 인지** | `errors.Is(err, io.EOF)` |
| `errors.As(err, &target)` | **특정 타입으로 추출** (필드 접근 가능) | `var pe *os.PathError; errors.As(err, &pe)` |

`errors.As`는 체인을 따라 내려가며 target 타입에 맞는 첫 에러를 `target`에 대입.

```go
var dze *DivisionByZeroError
if errors.As(err, &dze) {
    fmt.Println(dze.Numerator)  // 래핑된 에러에서도 필드 접근 가능
}
```

### 언제 sentinel, 언제 custom type?

| 상황 | 선택 |
|------|------|
| 동등성만 체크 (`if err == ErrNotFound`) | sentinel: `var ErrNotFound = errors.New(...)` |
| 추가 데이터 (필드) 필요 | custom type |
| API 컨트랙트의 일부 | custom type (확장 가능) |

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `fmt.Errorf("...: %v", err)` 사용 | 단순 문자열로만 변환, 체인 끊김 → `errors.Is/As` 동작 안 함 |
| `errors.As(err, dze)` (포인터 아님) | 컴파일 에러 또는 panic — 반드시 `&dze` |
| 값 리시버로 `Error()` 정의 | 동작은 하지만 Go 컨벤션 위배, `errors.As` 시그니처 안 맞을 수 있음 |
