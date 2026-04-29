# 06-extra1. Rank Students — 정답 및 해설

## 정답 코드

```go
func RankStudents(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		if students[i].Score != students[j].Score {
			return students[i].Score > students[j].Score   // 내림차순
		}
		return students[i].Name < students[j].Name         // 오름차순
	})
}
```

## 해설

### 같은 패턴 — 다중 키 정렬
06번 SortUsers 와 골격이 동일. 차이는 1차 키가 **내림차순** 이라 비교 부등호 방향이 반대.

### 오름차순 vs 내림차순
```go
return a < b   // 오름차순 (작은 게 앞)
return a > b   // 내림차순 (큰 게 앞)
```

less 함수의 의미는 "i가 j보다 앞에 와야 하는가". 점수 내림차순이라면 점수가 큰 i가 앞에 와야 하므로 `>`.

### 음수 트릭은 피할 것
```go
// ❌ "내림차순은 -1 곱하면 됨" — 정수 오버플로우 가능
return -students[i].Score < -students[j].Score

// ✅ 부등호 방향을 직접 뒤집기
return students[i].Score > students[j].Score
```

## 자주 하는 실수

| 실수 | 결과 |
|------|------|
| `Score < Score` 그대로 사용 | 점수 낮은 학생이 앞 — 랭킹 의미 반대 |
| 1차 키만 처리, tie-break 없음 | 같은 점수 학생들의 순서가 비결정적 |
| `sort.SliceStable` 만 쓰고 less 안에서 1차만 비교 | 동작은 하지만 less가 다중 키를 갖는 게 명시적 |
