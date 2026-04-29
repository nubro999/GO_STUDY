# 07-extra2. Storage Interface — 정답 및 해설

## 정답 코드

```go
type MemoryStorage struct {
	data map[string]string
}

func (m *MemoryStorage) Save(key, value string) error {
	if m.data == nil {
		m.data = map[string]string{}
	}
	m.data[key] = value
	return nil
}

func (m *MemoryStorage) Load(key string) (string, error) {
	if v, ok := m.data[key]; ok {
		return v, nil
	}
	return "", ErrNotFound
}

type NullStorage struct{}

func (NullStorage) Save(key, value string) error { return nil }
func (NullStorage) Load(key string) (string, error) {
	return "", ErrNotFound
}

func SaveAll(s Storage, items map[string]string) error {
	for k, v := range items {
		if err := s.Save(k, v); err != nil {
			return err
		}
	}
	return nil
}
```

## 해설

### 같은 패턴 — 인터페이스로 의존성 추상화
실무에서 인터페이스를 가장 많이 쓰는 이유는 **구현체 교체** 입니다.

- 운영: PostgresStorage / RedisStorage / S3Storage
- 테스트: MemoryStorage / NullStorage / 가짜(mock)

`SaveAll` 함수는 어떤 구현체가 들어올지 모르고도 동작 — 결합도가 낮음.

### 포인터 리시버를 쓰는 이유
```go
func (m *MemoryStorage) Save(...)
```
- `m.data = ...` 처럼 필드를 변경 → 포인터 리시버 필수
- 값 리시버였다면 호출 측의 사본만 변경되어 데이터 유실

이게 Shape 문제(값 리시버)와 다른 점: **상태를 가진 타입은 포인터 리시버**.

### `lazy init` 패턴
```go
if m.data == nil {
    m.data = map[string]string{}
}
```
`MemoryStorage` 의 zero value 를 그대로 사용 가능하도록 함. 호출자가 `&MemoryStorage{}` 만 써도 동작.

명시적 생성자 패턴도 흔함:
```go
func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{data: map[string]string{}}
}
```

### Sentinel Error
```go
var ErrNotFound = errors.New("key not found")
```
- 패키지 레벨 변수.
- `errors.Is(err, ErrNotFound)` 로 비교 — 직접 `==` 비교도 되지만 래핑된 에러도 처리하려면 `errors.Is` 사용.

### NullStorage 패턴
"아무것도 안 하는 구현체" 를 두면 호출자가 nil 체크 없이 항상 인터페이스를 통해 호출 가능. **Null Object 패턴** 이라 부릅니다. 테스트와 기본값에 유용.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| 값 리시버로 `Save` 정의 | 데이터 변경이 호출자에게 보이지 않음 |
| `m.data` 초기화 누락 | `assignment to entry in nil map` panic |
| 인터페이스를 만족하지 않은 채 넘김 | 컴파일 에러. `var _ Storage = (*MemoryStorage)(nil)` 으로 컴파일 시점 체크 가능 |
| `if err == ErrNotFound` 만 사용 | 래핑된 에러는 놓침 — `errors.Is` 권장 |

### 컴파일-타임 인터페이스 체크
```go
var _ Storage = (*MemoryStorage)(nil)
var _ Storage = NullStorage{}
```
이 두 줄을 두면 메서드 시그니처가 어긋났을 때 **컴파일 시점에 즉시 발견**. 큰 코드베이스에서 권장.
