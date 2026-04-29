# 11. Worker Pool — 정답 및 해설

## 정답 코드

```go
package main

import (
	"fmt"
	"sync"
)

func Process(jobs []int, workers int, f func(int) int) []int {
	n := len(jobs)
	if n == 0 {
		return []int{}
	}
	if workers <= 0 {
		workers = 1
	}

	type task struct {
		idx int
		val int
	}
	in := make(chan task)
	out := make([]int, n)

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range in {       // 채널이 close되면 자연 종료
				out[t.idx] = f(t.val)
			}
		}()
	}

	for i, v := range jobs {
		in <- task{idx: i, val: v}
	}
	close(in)   // 더 이상 보낼 것 없음을 알림
	wg.Wait()

	return out
}
```

## 해설

### 워커 풀 패턴의 핵심

```
[main: producer]   →   [chan in]   →   [N개 워커: consumers]
                                          ↓
                                    [out 슬라이스에 결과]
```

1. 워커 N개를 미리 띄움 (`for-range chan`).
2. main이 입력을 채널에 쏟아부음.
3. `close(in)` → 워커들의 for-range가 자연스럽게 종료.
4. `wg.Wait()` 으로 모든 워커 종료 대기.

### `for v := range ch` 의 종료 조건
채널을 `range` 로 돌리면 **채널이 close되고 비워질 때까지** 수신을 계속합니다.

```go
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    close(ch)
}()
for v := range ch { fmt.Println(v) }  // 1, 2 출력 후 종료
```

`close` 안 하면 worker 고루틴들이 영원히 대기 → **goroutine leak**.

### 채널을 닫는 책임
**송신자가 닫는다.** 절대 수신자가 닫으면 안 됩니다 (이미 닫힌 채널에 송신 시 panic).

```go
// ✅ producer가 close
go func() {
    for _, v := range data {
        ch <- v
    }
    close(ch)
}()

// ❌ consumer가 close — 다른 producer가 있으면 panic
for v := range ch { ... }
close(ch)
```

### 입력 순서 보존 트릭
`(idx, val)` 쌍을 흘리고, 워커가 `out[idx]`에 직접 씀. 워커들이 완료 순서가 뒤섞여도 인덱스 기반으로 결과 위치가 고정됩니다.

각 워커가 **서로 다른 인덱스**에만 쓰므로 race condition 없음. (10번 문제와 동일한 패턴.)

### 언버퍼드 vs 버퍼드 채널

| 종류 | 동작 | 사용 시기 |
|------|------|-----------|
| `make(chan T)` | 송수신 동기화, 양쪽이 만나야 진행 | 단순 핸드오프, 자연스러운 backpressure |
| `make(chan T, N)` | N개까지 버퍼, 가득 차면 송신 블록 | 버스트 입력 흡수, 송수신 속도 차이 완화 |

워커 풀에서는 보통 **언버퍼드** 또는 워커 수만큼 작은 버퍼.

## 다른 풀이: 결과도 채널로

```go
type result struct{ idx, val int }
in := make(chan task)
res := make(chan result)

// workers
for w := 0; w < workers; w++ {
    go func() {
        for t := range in {
            res <- result{t.idx, f(t.val)}
        }
    }()
}

// producer
go func() {
    for i, v := range jobs {
        in <- task{i, v}
    }
    close(in)
}()

// collector
out := make([]int, n)
for i := 0; i < n; i++ {
    r := <-res
    out[r.idx] = r.val
}
```

장점: 결과 슬라이스에 직접 안 써서 더 분리됨.
단점: 채널 두 개 + 추가 고루틴 → 코드량 증가.

## 자주 하는 실수

| 실수 | 증상 |
|------|------|
| `close(in)` 누락 | 워커들이 영원히 대기 → goroutine leak |
| 워커 안에서 `close(in)` | producer가 살아있다면 send on closed channel → panic |
| 채널 두 번 close | 두 번째 close에서 panic |
| `wg.Wait()` 전에 결과를 읽음 | 일부 결과가 누락된 채로 반환 |
