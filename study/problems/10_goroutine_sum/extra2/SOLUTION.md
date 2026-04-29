# 10-extra2. Safe Counter — 정답 및 해설

## 정답 코드 (Mutex 방식)

```go
func SafeCounter(workers, perWorker int) int {
	var mu sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < perWorker; i++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return counter
}
```

## 해설

### 같은 패턴 — 고루틴 + WaitGroup
ParallelSum 과 동일하지만 **공유 가변 상태(counter)** 가 있습니다. 그래서 보호가 필요.

### 왜 Mutex 가 필요한가?

```go
counter++   // 이게 원자(atomic) 연산이 아님!
```

내부적으로 다음 3단계로 분해됩니다:
1. counter 메모리에서 값 읽기
2. +1 계산
3. counter에 다시 쓰기

두 고루틴이 동시에 1단계에서 같은 값을 읽으면, 둘 다 같은 값을 쓰기 때문에 +1 만 반영됩니다. 결과적으로 카운트가 모자람.

`go run -race` 로 실행하면 즉시 경고가 뜹니다.

### Mutex 사용 패턴
```go
mu.Lock()
defer mu.Unlock()  // 짧은 함수에선 defer 패턴이 안전
// critical section
```

또는 짧은 critical section 에서는:
```go
mu.Lock()
counter++
mu.Unlock()
```

`defer` 의 장점: panic 이 나도 unlock 보장.
단점: 함수가 끝날 때까지 잠금 유지 (긴 함수에선 비효율).

### 더 빠른 대안: `sync/atomic`

```go
import "sync/atomic"

var counter atomic.Int64
// ...
counter.Add(1)
```

- Mutex 없이 원자적 증가.
- 단일 카운터 정도라면 atomic 이 보통 더 빠름.
- 복잡한 상태(여러 필드 동시 갱신)에는 Mutex 가 더 직관적.

### 채널 방식 (Go-스타일)
```go
done := make(chan struct{}, workers*perWorker)
for w := 0; w < workers; w++ {
    go func() {
        for i := 0; i < perWorker; i++ {
            done <- struct{}{}
        }
    }()
}
counter := 0
for i := 0; i < workers*perWorker; i++ {
    <-done
    counter++
}
```

채널이 자체로 동기화되므로 race 없음. 하지만 단순 카운트엔 오버킬 — Mutex/atomic 이 적절.

### "Don't communicate by sharing memory; share memory by communicating"
Go 의 동시성 격언입니다. **공유 변수 + Mutex** 보다 **채널 + 메시지 전달** 을 선호하라는 가이드. 단, 본 문제처럼 단순 카운터는 Mutex/atomic 이 더 자연스럽습니다.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| Mutex 없이 `counter++` | data race, 결과 < 기대값 |
| `mu.Lock()` 후 `mu.Lock()` (재진입) | deadlock — Go의 sync.Mutex 는 재진입 불가 |
| `defer mu.Unlock()` 누락 후 panic | mutex 영원히 잠김 → 다른 고루틴 deadlock |
| 고루틴 안에서 `mu` 의 복사본 캡처 | 별개의 mutex로 동작 → 동기화 안 됨. 항상 포인터로 |
