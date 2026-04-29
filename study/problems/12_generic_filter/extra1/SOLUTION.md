# 12-extra1. Generic Set — 정답 및 해설

## 정답 코드

```go
type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: map[T]struct{}{}}
}

func (s *Set[T]) Add(v T)              { s.m[v] = struct{}{} }
func (s *Set[T]) Remove(v T)           { delete(s.m, v) }
func (s *Set[T]) Contains(v T) bool    { _, ok := s.m[v]; return ok }
func (s *Set[T]) Size() int            { return len(s.m) }
```

## 해설

### 같은 패턴 — 제네릭 함수에서 제네릭 타입으로 확장
Map/Filter/Reduce 는 함수에 타입 파라미터를 적용했고, Set 은 **타입(struct)** 자체에 타입 파라미터를 적용합니다.

### 제약 `comparable` 의 의미
```go
type Set[T comparable] struct { ... }
```

- `comparable` 은 `==`, `!=` 가 가능한 타입.
- map의 키로 쓸 수 있는 타입이 정확히 comparable.
- int, string, struct(필드들이 모두 comparable인 경우), pointer, channel 등.
- **slice, map, function 은 comparable 아님** — 이들은 Set 원소가 될 수 없음.

### `map[T]struct{}` — Go 표준 Set 패턴
```go
m[key] = struct{}{}    // "있음" 표시
delete(m, key)
_, ok := m[key]        // 존재 여부
```

- `struct{}` 은 **0 바이트** — 메모리 효율 최고.
- 값을 비교할 일이 없으니 자료형은 의미 없고, 표시자(presence marker) 역할.
- `bool` 을 쓰면 1 바이트 차지하므로 `struct{}` 가 관례.

### 메서드에 타입 파라미터 재선언?
```go
func (s *Set[T]) Add(v T) { ... }
//        ^^^
//        타입 파라미터 이름을 메서드 리시버에서 다시 명시
```

이건 "재선언" 이 아니라 **참조** — Set 의 T를 메서드 안에서 사용 가능하게 함. 새 타입 파라미터를 메서드에서 추가할 수 없음 (Go 의 의도적 제약):

```go
// ❌ 컴파일 에러
func (s *Set[T]) Map[U any](f func(T) U) *Set[U] { ... }
```

이런 변환은 free 함수로 빼야 함:
```go
func MapSet[T, U comparable](s *Set[T], f func(T) U) *Set[U] { ... }
```

### 타입 추론 vs 명시
```go
s := NewSet[int]()           // 명시 (필요 — 빈 호출이라 추론 불가)
s2 := NewSet[string]()
```

호출 인자에서 타입을 추론할 수 없으면 명시 필요. 본 문제처럼 인자 없는 생성자는 **명시 필수**.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `Set[T any]` 로 정의 | 키로 쓸 수 없는 타입(slice 등)을 받게 되어 map 사용 시 컴파일 에러 |
| `m[v] = true` (bool 값) | 동작은 하지만 메모리 낭비 — 관례는 `struct{}{}` |
| 메서드를 값 리시버 `func (s Set[T]) Add(v T)` | Add가 호출자에게 반영 안 됨 — 포인터 리시버 필수 |
| `NewSet()` (타입 인자 누락) | 컴파일 에러 — 추론 못함 |
