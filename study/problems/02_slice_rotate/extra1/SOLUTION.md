# 02-extra1. Flatten — 정답 및 해설

## 정답 코드

```go
func Flatten(nested [][]int) []int {
	total := 0
	for _, sub := range nested {
		total += len(sub)
	}
	out := make([]int, 0, total)
	for _, sub := range nested {
		out = append(out, sub...)
	}
	return out
}
```

## 해설

### 같은 패턴 — 용량 미리 할당
회전 문제와 동일한 idiom: 결과 길이를 미리 알 수 있을 땐 `make([]T, 0, total)` 로 용량을 잡아 재할당을 피합니다.

### `append(out, sub...)` — 가변 인자 + spread
```go
append(out, sub...)
```
- `...` 는 슬라이스를 가변 인자로 풀어서 전달.
- 안쪽에서 `for _, v := range sub { out = append(out, v) }` 와 동일하지만 더 짧고 컴파일러 최적화에 유리.

### 빈 슬라이스 vs nil
- 빈 입력 `[][]int{}` 에 대해 `nil` 을 반환할지 `[]int{}` 를 반환할지는 시그니처 차이.
- `make([]int, 0, 0)` 은 빈 non-nil 슬라이스. 대부분의 경우 둘 다 같이 동작하지만 `reflect.DeepEqual` 이나 JSON 인코딩에서 차이남 (`null` vs `[]`).
- 본 풀이는 `make` 사용으로 빈 슬라이스를 일관되게 반환.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `append(out, sub)` (`...` 누락) | `[]int{1,2,3}` 을 한 원소로 끼우려 함 → 컴파일 에러 |
| 길이 계산 없이 `make([]int, 0)` | 큰 입력에서 여러 번 재할당 (성능만의 문제) |
