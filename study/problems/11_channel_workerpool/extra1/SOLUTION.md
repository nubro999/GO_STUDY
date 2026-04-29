# 11-extra1. 3-Stage Pipeline — 정답 및 해설

## 정답 코드

```go
func gen(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func sum(in <-chan int) int {
	total := 0
	for n := range in {
		total += n
	}
	return total
}

func Pipeline(nums []int) int {
	return sum(square(gen(nums)))
}
```

## 해설

### 같은 패턴 — 워커 풀의 일반화
워커 풀(11번 문제)에서는 한 단계의 분산 처리였다면, 파이프라인은 **여러 단계의 직렬 연결**입니다. 채널 + close + for-range 의 골격은 동일.

### 단방향 채널 타입
```go
<-chan int   // 수신 전용
chan<- int   // 송신 전용
chan int     // 양방향 (기본)
```

함수 시그니처에서 **방향을 명시**하는 게 권장됩니다:
- 호출자가 잘못된 방향으로 사용하는 걸 컴파일러가 잡음
- 함수의 의도가 명확

### 송신자의 close 책임 (반복)
각 stage 마다 **자기가 만드는 출력 채널은 자기가 close**.

```go
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)            // 이 함수가 만든 채널이므로 책임
        for n := range in { ... }   // 입력 채널의 close 는 호출자/이전 stage 가 처리
    }()
    return out
}
```

- `in` 은 받기만 하므로 close 안 함 (이미 단방향 타입이라 close 시도해도 컴파일 에러).
- `defer close(out)` 으로 panic / 정상 종료 모두 안전.

### 자동 backpressure
언버퍼드 채널이라 **다음 stage 가 받기 전까지 송신이 블록**됩니다. 이 자체가 backpressure — 빠른 stage 가 느린 stage를 압도하지 않음.

큰 throughput 이 필요하면 채널에 작은 버퍼를 줄 수 있음:
```go
out := make(chan int, 100)
```

### 메모리 사용량
파이프라인은 **중간 결과를 슬라이스에 다 모으지 않으므로** 큰 입력에서도 메모리 효율적. stream 처리에 적합.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| 송신 stage 에서 `close` 누락 | 다음 stage 의 for-range 가 영원히 대기 → goroutine leak |
| `gen` 이 `nums` 슬라이스 변경 후 송신 시도 | 데이터 race |
| `for-range in` 대신 `for { v := <-in }` 무한 루프 | close 후에도 zero value 를 영원히 받음 — for-range 권장 |
| 양방향 채널 `chan int` 만 사용 | 잘못된 방향 사용 시 컴파일 시점에 못 잡음 |
