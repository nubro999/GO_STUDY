# 10. Parallel Sum (Goroutines + WaitGroup) — 정답 및 해설

## 정답 코드

```go
package main

import (
	"fmt"
	"sync"
)

func ParallelSum(nums []int, workers int) int {
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

	chunk := (n + workers - 1) / workers // ceil 분할
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
			sum := 0
			for _, v := range nums[s:e] {
				sum += v
			}
			partials[idx] = sum
		}(w, start, end)
	}
	wg.Wait()

	total := 0
	for _, p := range partials {
		total += p
	}
	return total
}
```

## 해설

### `sync.WaitGroup` 의 정확한 사용법

```go
var wg sync.WaitGroup
wg.Add(1)            // ① 고루틴 시작 전에 카운터 증가
go func() {
    defer wg.Done()  // ② 종료 시 카운터 감소 (defer로 panic에도 안전)
    // 작업
}()
wg.Wait()            // ③ 모든 Done이 호출될 때까지 블록
```

**`Add`는 반드시 `go` 문 *전에* 호출**해야 합니다. 고루틴 안에서 Add를 부르면 main이 그보다 먼저 Wait에 도달해 0으로 보고 빠져나갈 수 있습니다.

### 클로저 캡처 함정 (Go 1.21 이전)

```go
// ❌ 오래된 Go에서 버그
for w := 0; w < workers; w++ {
    go func() {
        process(w)  // 모든 고루틴이 같은 w를 봄!
    }()
}

// ✅ 인자로 넘김 (안전)
for w := 0; w < workers; w++ {
    go func(w int) {
        process(w)
    }(w)
}
```

Go 1.22부터는 루프 변수가 매 반복마다 새로 만들어지도록 **시맨틱 변경**이 있어 첫 번째 형태도 안전해졌지만, 기존 코드 호환과 명확성을 위해 인자 전달이 여전히 안전한 패턴입니다.

### 데이터 레이스 없는 결과 수집
모든 고루틴이 `partials[idx]` 의 **서로 다른 인덱스** 에만 씁니다. 같은 메모리 위치에 동시에 쓰지 않으므로 race condition 없음.

```bash
go run -race ./problems/10_goroutine_sum/
```
`-race` 플래그로 실제 race를 감지할 수 있습니다. 동시성 코드 작성 시 항상 검증하세요.

### 다른 풀이: 채널로 결과 수집

```go
results := make(chan int, workers)
for w := 0; w < workers; w++ {
    s, e := w*chunk, min((w+1)*chunk, n)
    go func() {
        sum := 0
        for _, v := range nums[s:e] {
            sum += v
        }
        results <- sum
    }()
}

total := 0
for i := 0; i < workers; i++ {
    total += <-results
}
```

- 채널은 자체로 동기화 메커니즘 — WaitGroup 불필요.
- 버퍼 크기를 워커 수와 맞추면 송신 블록 없음.

### Ceil 분할 트릭
```go
chunk := (n + workers - 1) / workers
```
정수 나눗셈에서 올림 효과 — `100 / 7 == 14` 이지만 `(100+6)/7 == 15`.
나머지 원소를 마지막 워커가 처리하도록 보장.

## 자주 하는 실수

| 실수 | 증상 |
|------|------|
| 고루틴 안에서 `wg.Add(1)` | Wait가 일찍 끝남 |
| `defer wg.Done()` 대신 함수 끝에 `wg.Done()` | panic 시 Done 호출 안 됨 → Wait 영원히 블록 |
| 같은 변수에 여러 고루틴이 `total += ...` | data race, `-race` 로 감지 가능 |
| 클로저로 `w`, `start`, `end` 직접 캡처 (구버전) | 모든 고루틴이 같은 값을 봄 |
