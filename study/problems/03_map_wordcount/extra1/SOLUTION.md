# 03-extra1. First Index — 정답 및 해설

## 정답 코드

```go
func FirstIndex(words []string) map[string]int {
	out := map[string]int{}
	for i, w := range words {
		if _, ok := out[w]; !ok {
			out[w] = i
		}
	}
	return out
}
```

## 해설

### 같은 패턴 — `comma-ok` idiom
WordCount 와 동일한 맵을 사용하지만 카운팅이 아닌 **"키가 이미 있는지"** 만 검사합니다.

```go
if _, ok := out[w]; !ok {
    out[w] = i
}
```

- `_` 로 값은 버리고 `ok` 로 존재 여부만 받음.
- 두 개의 인덱스(0, 1, 2, ...)와 첫 등장 단어를 동기화하는 트릭.

### `range` 의 두 변수
```go
for i, w := range words {
    // i: 인덱스, w: 값
}
```

- 인덱스만: `for i := range words`
- 값만: `for _, w := range words`
- 둘 다 무시: `for range words` (단순 N번 반복용으로 잘 안 씀)

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `out[w] = i` 만 함 (검사 누락) | 마지막 등장 인덱스가 저장됨 |
| `out[w] != 0` 으로 존재 검사 | 인덱스 0인 첫 단어가 "없음"으로 판정됨 |
