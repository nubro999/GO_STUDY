# 01-extra2. Digit Sum Prime — 정답 및 해설

## 정답 코드

```go
func digitSum(n int) int {
	s := 0
	for n > 0 {
		s += n % 10
		n /= 10
	}
	return s
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func CountDigitSumPrime(n int) int {
	cnt := 0
	for i := 1; i <= n; i++ {
		if isPrime(digitSum(i)) {
			cnt++
		}
	}
	return cnt
}
```

## 해설

### 같은 패턴 — 작은 함수 조합
큰 문제를 **자릿수 합 추출 + 소수 판별 + 카운트** 세 단계로 쪼갰습니다. 각 단계가 단순해서 디버깅과 테스트가 쉽습니다.

### `n % 10`, `n / 10` — 자릿수 추출 관용구
```go
1234 % 10 = 4    // 마지막 자리
1234 / 10 = 123  // 마지막 자리 제거
```
이 패턴은 자릿수를 다루는 모든 문제에서 등장.

### 소수 판별 최적화
- 6 까지의 약수가 있으면 36 미만에서 발견됨 → `i*i <= n` 까지만 검사 (O(√n))
- `for i := 2; i <= n/2` 보다 빠름

### 캐싱
같은 자릿수 합이 반복적으로 등장하면 `map[int]bool` 으로 소수 여부 캐싱 가능. 자릿수 합은 최대 9*자릿수 = 작은 수라 큰 차이는 없음.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `for i := 2; i < n; i++` 로 모든 수 검사 | O(n) 으로 느림 |
| `if n == 1 { return true }` | 1은 소수 아님 |
| `digitSum(0) == 0` 무시 | 0은 입력 안 들어오지만 경계값 인지 |
