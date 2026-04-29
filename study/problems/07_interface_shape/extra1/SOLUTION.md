# 07-extra1. Animal Interface — 정답 및 해설

## 정답 코드

```go
type Dog struct{ Nickname string }
func (d Dog) Sound() string { return "Woof" }
func (d Dog) Name() string  { return d.Nickname }

type Cat struct{ Nickname string }
func (c Cat) Sound() string { return "Meow" }
func (c Cat) Name() string  { return c.Nickname }

type Cow struct{ Nickname string }
func (c Cow) Sound() string { return "Moo" }
func (c Cow) Name() string  { return c.Nickname }

func Chorus(animals []Animal) []string {
	out := make([]string, 0, len(animals))
	for _, a := range animals {
		out = append(out, fmt.Sprintf("%s: %s", a.Name(), a.Sound()))
	}
	return out
}
```

## 해설

### 같은 패턴 — Shape 와 동일 골격
- 인터페이스 1개 → 구체 타입 N개 구현
- `[]Interface` 슬라이스 → 다형적 순회

### 공통 필드를 공유하고 싶다면 임베딩
모든 동물이 `Nickname` 을 갖는다면 다음처럼 추출 가능:

```go
type Named struct{ Nickname string }
func (n Named) Name() string { return n.Nickname }

type Dog struct{ Named }
func (Dog) Sound() string { return "Woof" }

type Cat struct{ Named }
func (Cat) Sound() string { return "Meow" }
```

- `Dog{Named: Named{"Rex"}}` 또는 `Dog{Named{"Rex"}}` 로 생성
- `Dog.Name()` 호출 가능 — 임베딩 필드의 메서드는 자동 promotion

이게 Go의 "상속 대신 합성(composition)" 철학.

### 메서드가 받는 값을 안 쓰면 `_`
```go
func (Dog) Sound() string { return "Woof" }
```
리시버 변수가 안 쓰이면 이름 생략 가능. 작은 코드 정리 팁.

### 여러 인터페이스 결합

```go
type Animal interface {
    Sound() string
    Name() string
}

// 동등하게:
type Sounder interface{ Sound() string }
type Namer   interface{ Name() string }
type Animal  interface{ Sounder; Namer }
```
Go 1.18+ 에서 인터페이스 임베딩이 더 강력해졌습니다.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| 한 메서드만 구현 (`Sound` 만) | 인터페이스 미만족 → `[]Animal{Dog{...}}` 컴파일 에러 |
| 포인터 리시버로 정의 후 값으로 슬라이스 생성 | 메서드 셋 불일치 → 컴파일 에러 |
| `Dog{Nickname}` 만 쓰고 필드명 생략 | 키워드 인자(`{Nickname: "Rex"}`) 가 더 안전, 필드 추가 시 깨짐 방지 |
