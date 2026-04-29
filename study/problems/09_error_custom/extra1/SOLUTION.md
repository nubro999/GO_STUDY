# 09-extra1. Multi-Error Classify — 정답 및 해설

## 정답 코드

```go
type NotFoundError struct{ Resource string }
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Resource)
}

type ValidationError struct{ Field, Reason string }
func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s: %s", e.Field, e.Reason)
}

func Classify(err error) string {
	if err == nil {
		return "none"
	}

	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return "not_found"
	}

	var ve *ValidationError
	if errors.As(err, &ve) {
		return "validation"
	}

	return "unknown"
}
```

## 해설

### 같은 패턴 — 커스텀 에러 + `errors.As`
SafeDivide 와 동일한 골격이지만 **타입이 여러 개** 입니다. `errors.As` 가 핵심.

### `errors.As` 동작 방식
```go
var target *NotFoundError
errors.As(err, &target)
```

1. 에러 체인을 따라 내려감 (`Unwrap` 반복).
2. 각 노드를 `*NotFoundError` 로 형변환 시도.
3. 성공하면 `target` 에 할당하고 `true` 반환.

따라서 래핑된 에러 `fmt.Errorf("...: %w", &NotFoundError{...})` 도 정상 분류.

### 분기 순서
```go
if err == nil { return "none" }   // 가장 먼저
// 구체 타입 검사
if errors.As(err, &nfe) { return "not_found" }
if errors.As(err, &ve) { return "validation" }
// fallback
return "unknown"
```

`nil` 체크는 가장 처음. `errors.As(nil, ...)` 는 false 반환하지만 명시적으로 처리하는 게 의도가 분명.

### Type switch 방식 (대안)
```go
switch err.(type) {
case *NotFoundError:
    return "not_found"
case *ValidationError:
    return "validation"
}
```

- 짧지만 **래핑된 에러를 못 잡음**. `fmt.Errorf("%w", ...)` 로 감싸진 에러는 type assertion 실패.
- 새 코드는 `errors.As` 권장.

### `errors.Is` vs `errors.As` 다시 정리

| 케이스 | 사용 |
|--------|------|
| sentinel error (`var ErrFoo = errors.New(...)`) | `errors.Is(err, ErrFoo)` |
| 커스텀 타입의 필드 접근 필요 | `errors.As(err, &target)` |
| 단순히 "저 타입인지" 만 확인 | `errors.As(err, &target)` (target 은 안 써도 됨) |

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `if _, ok := err.(*NotFoundError); ok` (type assertion) | 래핑된 에러 못 잡음 |
| `errors.As(err, nfe)` (포인터 아님) | 컴파일 에러 또는 panic |
| `target` 변수를 분기 밖에 선언 후 재사용 | 이전 분기에서 할당된 값이 남아있을 수 있음 — 분기마다 새 변수 |
