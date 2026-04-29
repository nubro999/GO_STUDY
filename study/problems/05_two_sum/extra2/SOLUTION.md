# 05-extra2. Group Anagrams — 정답 및 해설

## 정답 코드

```go
func GroupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, s := range strs {
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		key := string(bs)
		groups[key] = append(groups[key], s)
	}
	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		out = append(out, g)
	}
	return out
}
```

## 해설

### 같은 패턴 — 맵 키 = 정규화된 표현
TwoSum 에서는 "값" 자체를 키로 썼지만, 여기서는 "정렬된 글자 = 애너그램 표준형" 을 키로 사용. 같은 그룹은 같은 표준형을 가집니다.

### 정규화 키의 다양한 변형

```go
// 1) 정렬 (O(k log k))
sort.Slice(bs, ...)
key := string(bs)

// 2) 빈도 카운트 (O(k))
cnt := [26]int{}
for _, c := range s {
    cnt[c-'a']++
}
key := fmt.Sprintf("%v", cnt)
```

작은 알파벳에선 빈도 카운트가 더 빠르지만 정렬이 짧고 직관적.

### `[]byte(s)` 정렬 트릭
ASCII 문자열을 정렬할 때 `[]rune` 보다 `[]byte` 가 가벼움. multi-byte 문자가 섞이면 `[]rune` 사용.

### `m[k] = append(m[k], v)` (또 등장)
03-extra2 와 동일한 idiom — map의 nil 슬라이스에도 안전하게 append.

### 출력 슬라이스 만들기
```go
out := make([][]string, 0, len(groups))
for _, g := range groups {
    out = append(out, g)
}
```
- `len(map)` 으로 그룹 수를 알 수 있어 용량 미리 할당.
- map의 순회 순서는 **무작위** — 테스트는 집합 비교를 사용해야 함 (실제 main 코드에서도 그렇게 검증).

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `string(bs)` 대신 `fmt.Sprint(bs)` | "[97 101 116]" 같은 문자열로 키가 길어짐 (동작은 함) |
| 정렬할 때 `string` 직접 정렬 시도 | string 은 immutable — `[]byte` 또는 `[]rune` 필요 |
| 맵 순회 순서에 의존하는 테스트 | 매 실행마다 결과 순서 다름 → 집합 비교로 검증해야 |
