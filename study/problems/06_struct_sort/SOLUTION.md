# 06. Struct Sort — 정답 및 해설

## 정답 코드

```go
package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

func SortUsers(users []User) {
	sort.Slice(users, func(i, j int) bool {
		if users[i].Age != users[j].Age {
			return users[i].Age < users[j].Age
		}
		return users[i].Name < users[j].Name
	})
}
```

## 해설

### `sort.Slice` 의 시그니처
```go
sort.Slice(slice any, less func(i, j int) bool)
```
- `slice`: 정렬할 슬라이스 (in-place)
- `less(i, j)`: i번째가 j번째보다 **앞에 와야 하면** true

less 함수가 비교 로직 전체를 결정합니다.

### 다중 키 정렬 패턴
```go
if a.Key1 != b.Key1 {
    return a.Key1 < b.Key1   // 1차 키
}
return a.Key2 < b.Key2       // tie-break
```

3개 이상의 키가 필요하면 같은 패턴을 중첩:
```go
if a.K1 != b.K1 { return a.K1 < b.K1 }
if a.K2 != b.K2 { return a.K2 < b.K2 }
return a.K3 < b.K3
```

### `sort.Slice` vs `sort.SliceStable`
- `sort.Slice`: **불안정 정렬** (같은 키의 원소들 간 원래 순서 보장 안 됨)
- `sort.SliceStable`: 안정 정렬 (원래 순서 유지)

다중 키 정렬을 less 안에서 모두 처리한다면 안정성이 필요 없습니다. 그러나 1차 키만 비교하고 2차 키는 "원래 순서 유지"하고 싶다면 `SliceStable`을 쓰세요.

### Go 1.21+: `slices.SortFunc`
표준 패키지 `slices` 가 추가되면서 새 idiom이 생겼습니다:

```go
import (
	"cmp"
	"slices"
)

slices.SortFunc(users, func(a, b User) int {
    if c := cmp.Compare(a.Age, b.Age); c != 0 {
        return c
    }
    return cmp.Compare(a.Name, b.Name)
})
```

- `cmp.Compare`: -1 / 0 / +1 반환
- 두 원소를 직접 받으므로 인덱스 헷갈릴 일이 없음
- 새 코드에서는 이쪽이 권장

## 자주 하는 실수

| 실수 | 증상 |
|------|------|
| `return users[i].Age <= users[j].Age` | `<=`는 less 정의에 어긋남, 동등 시 비교 함수가 비대칭이 되어 panic 가능 |
| 1차 키만 비교하고 끝 | 같은 Age에서 Name 정렬 누락 |
| `sort.Sort(byAge(users))` 로 인터페이스 직접 구현 | 가능하지만 보일러플레이트가 많음 — `sort.Slice`가 훨씬 간결 |
