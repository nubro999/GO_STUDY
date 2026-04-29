# 11-extra2. Fan-In — 정답 및 해설

## 정답 코드

```go
func FanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
```

## 해설

### 같은 패턴 — 멀티 producer
워커 풀에선 **여러 consumer 가 하나의 입력 채널을 공유**했습니다.
Fan-In 은 그 반대 — **여러 producer 가 하나의 출력 채널에 송신**.

### 핵심 idiom: WaitGroup + 별도 close 고루틴
```go
go func() {
    wg.Wait()   // 모든 forwarder 종료까지 대기
    close(out)  // 그 후 안전하게 close
}()
```

**왜 별도 고루틴인가?**
- `FanIn` 함수는 즉시 `out` 을 반환해야 함 (블록되면 안 됨).
- `wg.Wait()` 은 모든 forwarder 가 끝나기 전엔 블록.
- 따라서 close 처리를 별도 고루틴에 위임해서 `FanIn` 은 비동기로 즉시 리턴.

### 가변 인자 + 단방향 채널
```go
func FanIn(channels ...<-chan int) <-chan int
```
- `...` 가변 인자 + `<-chan int` 수신 전용 채널.
- 호출 측은 `FanIn(ch1, ch2, ch3)` 또는 `FanIn(channels...)` 형태로 사용.

### 클로저 캡처 함정 (또 등장)
```go
for _, ch := range channels {
    go func(c <-chan int) {  // 인자로 넘김
        for v := range c { out <- v }
    }(ch)
}
```

루프 변수 `ch` 를 직접 캡처하면 (Go 1.21 이전) 모든 고루틴이 같은 ch 를 보게 됨. 인자 전달이 안전.

### 송신 충돌 없음
여러 forwarder가 동시에 `out <- v` 해도 **채널 자체가 동기화**되어 race 없음. Mutex 불필요.
이게 "share memory by communicating" 의 실제 예.

### Select 기반 다른 풀이

```go
go func() {
    defer close(out)
    cases := make([]reflect.SelectCase, len(channels))
    for i, ch := range channels {
        cases[i] = reflect.SelectCase{
            Dir: reflect.SelectRecv,
            Chan: reflect.ValueOf(ch),
        }
    }
    remaining := len(cases)
    for remaining > 0 {
        i, v, ok := reflect.Select(cases)
        if !ok {
            cases[i].Chan = reflect.Value{}  // 비활성화
            remaining--
            continue
        }
        out <- v.Interface().(int)
    }
}()
```

reflect 기반은 **유연하지만 느리고 복잡** — 일반적인 경우 WaitGroup 방식이 명료.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `close(out)` 를 forwarder 안에서 호출 | 첫 채널이 close 되면 out 도 닫힘 — 다른 forwarder들이 send on closed channel 로 panic |
| `wg.Wait()` 을 main 고루틴에서 호출 | `FanIn` 이 블록되어 호출자가 채널을 받을 수 없음 → deadlock |
| `wg.Add(len(channels))` 후 `wg.Add(1)` 도 호출 | 카운터 두 배 증가 |
