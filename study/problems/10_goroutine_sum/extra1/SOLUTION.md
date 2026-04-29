# 10-extra1. Parallel Max — 정답 및 해설

## 정답 코드

```go
func ParallelMax(nums []int, workers int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if workers <= 0 {
		workers = 1
	}
	if workers > n {
		workers = n
	}

	chunk := (n + workers - 1) / workers
	partials := make([]int, workers)
	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		start := w * chunk
		end := start + chunk
		if end > n {
			end = n
		}
		wg.Add(1)
		go func(idx, s, e int) {
			defer wg.Done()
			m := nums[s]
			for _, v := range nums[s+1 : e] {
				if v > m {
					m = v
				}
			}
			partials[idx] = m
		}(w, start, end)
	}
	wg.Wait()

	result := partials[0]
	for _, p := range partials[1:] {
		if p > result {
			result = p
		}
	}
	return result
}
```

## 해설

### 같은 패턴 — ParallelSum 의 max 변종
- 청크 분할 → 각 워커가 부분 결과 → 최종 합산.
- 합산 대신 max를 취하는 것만 다름.
- 부분 결과 슬라이스에 race 없이 쓰는 패턴 동일.

### Max의 초기값 처리
**합과 다른 점**: 합은 0으로 시작해도 OK지만, **max는 적절한 초기값이 까다롭습니다**.

```go
// ❌ 위험
m := 0
for _, v := range chunk { if v > m { m = v } }
// 음수만 있는 청크면 0 반환 — 잘못됨
```

```go
// ✅ 첫 원소를 초기값으로
m := chunk[0]
for _, v := range chunk[1:] { if v > m { m = v } }
```

`math.MinInt` 를 쓸 수도 있지만 첫 원소 사용이 더 안전.

### Reduce 관점
부분 결과 합산 단계 역시 max(reduce):
```go
result := partials[0]
for _, p := range partials[1:] { if p > result { result = p } }
```

ParallelSum 의 `total += p` 와 동일 골격이지만 연산만 바뀜. 일반화하면 임의의 결합법칙 연산(monoid)에 적용 가능.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| 초기값 0 | 음수 슬라이스에서 잘못된 max |
| 빈 청크에서 첫 원소 접근 | panic — 청크 분할 시 이미 보장되지만 직접 호출하면 위험 |
| 부분 결과 `partials[idx] = m` 대신 공유 변수 사용 | data race |
