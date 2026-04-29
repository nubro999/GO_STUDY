# 01-extra1. Collatz — 정답 및 해설

## 정답 코드

```go
func Collatz(n int) int {
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps
}
```

## 해설

### 같은 패턴 — `for` + `if/else` 분기
FizzBuzz와 동일한 골격: 한 변수를 반복적으로 갱신하며 조건에 따라 다른 분기를 탑니다. 차이는 종료 조건이 **횟수가 아니라 상태(`n != 1`)** 라는 점.

### `n /= 2` 와 `n = 3*n + 1`
- `/=` 는 단축 대입 연산자. `n = n / 2` 와 동일.
- 정수 나눗셈은 자동으로 floor (양수에서). 6/2 == 3, 7/2 == 3.

### `switch` true-form 으로도 가능
```go
switch {
case n%2 == 0:
    n /= 2
default:
    n = 3*n + 1
}
```
2-way 분기에서는 `if/else` 가 보통 더 짧지만, case가 늘면 switch가 깔끔.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `for n > 1` 으로 쓰고 `n == 1` 체크 누락 | n=1 입력 시 동작은 하지만 의미 모호 |
| 무한히 큰 수에서 overflow | 큰 n에서는 `int64` 또는 unbounded — 일반적인 입력에선 문제 없음 |
