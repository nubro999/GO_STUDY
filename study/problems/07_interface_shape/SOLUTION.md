# 07. Interface (Shape) — 정답 및 해설

## 정답 코드

```go
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Triangle struct {
	A, B, C float64
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}
func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }

func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}
```

## 해설

### Implicit interface satisfaction
Java/C#처럼 `implements` 키워드를 쓰지 않습니다. 메서드 셋만 일치하면 **자동으로** 인터페이스를 만족.

```go
var s Shape = Circle{Radius: 1}  // Circle이 Area/Perimeter 둘 다 가지므로 OK
```

이 덕분에 외부 패키지의 타입에도 인터페이스를 "사후에" 부여할 수 있어 결합도가 낮습니다 — Go의 대표적인 설계 철학.

### 값 리시버 vs 포인터 리시버
```go
func (c Circle) Area() float64       // 값 리시버
func (c *Circle) Area() float64      // 포인터 리시버
```

| 기준 | 선택 |
|------|------|
| 메서드가 필드를 변경 | 포인터 리시버 |
| struct이 큼 (수십 바이트 이상) | 포인터 리시버 (복사 비용 회피) |
| 작은 struct, 변경 없음 | 값 리시버 |
| 동일 타입의 다른 메서드와 일관성 | 통일 |

> 일관성 규칙: 한 타입의 메서드들은 **모두 값 리시버 또는 모두 포인터 리시버**로 통일하는 게 관례.

### 메서드 셋 (method set) 함정
- 값 리시버 메서드: `T`도 `*T`도 호출 가능
- 포인터 리시버 메서드: `*T`만 인터페이스 변수에 직접 대입 가능

```go
type Counter struct{ n int }
func (c *Counter) Inc() { c.n++ }

var x interface{ Inc() } = Counter{}   // ❌ 컴파일 에러
var y interface{ Inc() } = &Counter{}  // ✅ OK
```

### 다형성: `[]Shape`
서로 다른 구체 타입을 한 슬라이스에 담아 동일 인터페이스로 처리. 런타임에 동적 디스패치(가상 호출)가 일어납니다.

## 자주 하는 실수

| 실수 | 증상 |
|------|------|
| `func (Circle) Area()` 를 다른 파일에 두고 인식 안 됨 | 같은 패키지 안이어야 메서드 인식 |
| 포인터 리시버로 정의 후 `[]Shape{Circle{...}}` 시도 | 위 메서드 셋 규칙으로 컴파일 에러 |
| 헤론 공식의 `s` 계산 누락 | NaN 또는 잘못된 면적 |
