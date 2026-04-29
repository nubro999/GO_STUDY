# 02-extra2. Chunk — 정답 및 해설

## 정답 코드

```go
func Chunk(nums []int, size int) [][]int {
	if size <= 0 {
		return [][]int{}
	}
	n := len(nums)
	out := make([][]int, 0, (n+size-1)/size) // ceil
	for i := 0; i < n; i += size {
		end := i + size //
		if end > n {
			end = n
		}
		out = append(out, nums[i:end])
	}
	return out
}
```

## 해설

### 같은 패턴 — 슬라이싱 인덱스 계산
회전 문제처럼 인덱스 산술이 핵심. `nums[i:end]` 표현으로 부분 슬라이스를 만듭니다.

### Ceil 분할 트릭 (10번 ParallelSum과 동일)
```go
(n + size - 1) / size
```
정수 나눗셈에서 올림 효과. `[]int{1,2,3,4,5}, size=2` → `(5+1)/2 == 3` 청크.

### 슬라이싱과 메모리 공유
`nums[i:end]` 는 **원본과 같은 backing array를 공유**합니다. 한 청크를 수정하면 원본도 변경됩니다.

```go
chunks := Chunk([]int{1,2,3,4}, 2)
chunks[0][0] = 99
// nums[0]도 99가 됨
```

독립된 복사본이 필요하면:
```go
sub := make([]int, end-i)
copy(sub, nums[i:end])
out = append(out, sub)
```

### `min` 내장 함수
Go 1.21+ 에서 `min(a, b)`, `max(a, b)` 가 내장 함수로 추가됨.
```go
end := min(i+size, n)
```
1.21 미만에서는 직접 if문으로.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `i + size` 가 n을 넘는데 그대로 `nums[i:i+size]` | slice bounds out of range panic |
| `size == 0` 가능성 무시 | 무한 루프 |
| 청크 수 계산을 `n / size` 로만 함 | 마지막 청크 누락 |
