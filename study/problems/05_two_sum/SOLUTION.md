# 05. Two Sum — 정답 및 해설

## 정답 코드

```go
package main

import "fmt"

func TwoSum(nums []int, target int) (int, int) {
	seen := map[int]int{} // value -> earliest index
	for i, v := range nums {
		if j, ok := seen[target-v]; ok {
			return j, i
		}
		seen[v] = i
	}
	return -1, -1
}
```

## 해설

### 핵심 패턴: comma-ok idiom
```go
j, ok := seen[target-v]
```
맵 인덱싱은 두 개의 값을 반환할 수 있습니다.
- `j`: 키에 해당하는 값 (없으면 zero value)
- `ok`: 키 존재 여부 (`bool`)

존재 여부를 체크하지 않으면 "값이 0인 키"와 "키가 없음"을 구분할 수 없습니다. 이 idiom은 **type assertion(`v, ok := x.(T)`)**, **채널 수신(`v, ok := <-ch`)** 에서도 동일하게 등장 — Go의 핵심 패턴입니다.

### 단일 패스로 O(n)
naive 풀이는 이중 루프 O(n²):
```go
for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ { ... }
}
```

해시맵을 쓰면 **"이미 본 값"** 을 캐싱하고, 현재 값과 짝이 맞는지만 확인하면 됩니다. 한 번의 패스로 끝나 O(n).

### 인덱스 순서가 자동으로 맞음
`seen`에는 **이전에 본** 값들만 들어있으므로, 매치 시 `j`(맵에서 찾은 인덱스)는 항상 `i`(현재)보다 작습니다. 별도의 swap 없이 `(j, i)` 순으로 반환하면 됩니다.

### 저장 시점 주의
```go
if j, ok := seen[target-v]; ok {
    return j, i
}
seen[v] = i  // 매치 검사 후에 저장
```

만약 `seen[v] = i`를 먼저 한 뒤 검사하면 `nums = [3, 3], target = 6`에서 같은 인덱스 `(0, 0)`이 매치될 수 있습니다 — 자기 자신이 짝이 됨. **검사 후 저장**이 중요.

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `if seen[target-v] != 0 { ... }` | 값이 0인 합법적 케이스를 놓침 |
| `seen[v] = i` 를 먼저 함 | nums에 같은 값이 있을 때 자기 자신 매치 |
| 두 개의 인덱스 슬라이스 `[]int{i,j}` 반환 | 시그니처 불일치 (문제는 `(int, int)`) |
