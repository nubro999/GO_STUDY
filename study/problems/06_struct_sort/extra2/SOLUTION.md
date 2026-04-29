# 06-extra2. Sort Events — 정답 및 해설

## 정답 코드

```go
func SortEvents(events []Event) {
	sort.Slice(events, func(i, j int) bool {
		a, b := events[i], events[j]
		if a.Start != b.Start {
			return a.Start < b.Start
		}
		if a.End != b.End {
			return a.End < b.End
		}
		return a.Title < b.Title
	})
}
```

## 해설

### 같은 패턴 — 다중 키 비교 (3개)
키가 늘어나도 같은 패턴을 중첩하면 됩니다.

```go
if a.K1 != b.K1 { return a.K1 < b.K1 }
if a.K2 != b.K2 { return a.K2 < b.K2 }
return a.K3 < b.K3
```

### 지역 변수로 단축
```go
a, b := events[i], events[j]
```
`events[i].Field` 가 길어지면 가독성이 떨어지므로 미리 변수에 담기. **struct 복사가 일어나지만** 작은 struct이면 무시할 수준.

큰 struct이거나 핵심 경로면 포인터로:
```go
a, b := &events[i], &events[j]
```

### Go 1.21+: `slices.SortFunc` + `cmp.Or`
```go
import (
    "cmp"
    "slices"
)

slices.SortFunc(events, func(a, b Event) int {
    return cmp.Or(
        cmp.Compare(a.Start, b.Start),
        cmp.Compare(a.End, b.End),
        cmp.Compare(a.Title, b.Title),
    )
})
```

`cmp.Or` 는 0이 아닌 첫 결과를 반환 → 다중 키 비교를 1줄로. Go 1.22+ 권장 스타일.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| 1차/2차 키 모두 같은데 tie-break 없음 | 정렬 결과가 매 실행마다 달라질 수 있음 |
| `<=` 사용 | strict less 가 아니라 `sort.Slice` 가 panic 가능 |
| `cmp.Or` 안에 boolean 표현 섞기 | `cmp.Or` 는 int(-1/0/1) 만 다룸 — `cmp.Compare` 필수 |
