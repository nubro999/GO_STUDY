# 03-extra2. Group By First — 정답 및 해설

## 정답 코드

```go
func GroupByFirst(words []string) map[byte][]string {
	out := map[byte][]string{}
	for _, w := range words {
		if w == "" {
			continue
		}
		k := w[0]
		out[k] = append(out[k], w)
	}
	return out
}
```

## 해설

### 같은 패턴 — 맵 값으로 슬라이스
WordCount 의 `map[string]int` 와 같은 골격이지만, 값이 슬라이스인 케이스.

### 핵심 idiom: `m[k] = append(m[k], v)`
이게 가능한 이유:
- `m[k]` 가 키 부재 시 `nil` 반환 (슬라이스의 zero value)
- `append(nil, v)` 는 새 슬라이스 `[v]` 를 생성하고 반환

따라서 **키 존재 여부를 따로 검사할 필요 없음** — Go의 깔끔한 패턴 중 하나.

> ⚠️ map의 zero value 트릭은 슬라이스/int에 적용 가능하지만 **map의 값이 map인 경우엔 안 됨**.
> ```go
> var m map[string]map[string]int   // nil
> m["a"]["b"] = 1                    // panic: assignment to entry in nil map
> ```
> nested map 은 명시적 초기화 필요.

### `string[i]` 는 `byte`
```go
w := "apple"
w[0]            // byte (= uint8) 'a' = 0x61
```
ASCII 가정 시 안전. 한글/이모지가 들어오면 multi-byte라 첫 byte만 잘리므로 `[]rune(w)[0]` 사용.

### 맵의 zero value 정리

| 값 타입 | zero value | append 안전? |
|---------|-----------|--------------|
| `int`, `float` | 0 | `m[k]++` 안전 |
| `[]T` | nil | `m[k] = append(m[k], v)` 안전 |
| `map`, `chan` | nil | 직접 쓰면 panic — `make` 또는 첫 사용 시 초기화 |
| `*T`, `interface` | nil | 메서드 호출 시 nil 체크 필요 |

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `if _, ok := out[k]; ok { append... } else { out[k] = []string{w} }` | 코드만 길어짐 — `m[k] = append(m[k], v)` 한 줄로 OK |
| 빈 문자열 처리 누락 (`w[0]` 시도) | `index out of range` panic |
| rune 가정 시 `w[0]` 사용 | 한글이면 깨진 byte로 그룹핑됨 |
