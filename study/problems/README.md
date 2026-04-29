# Go 코딩테스트 연습 문제집

Go의 주요 패턴과 문법을 단계별로 학습하기 위한 문제 모음입니다.

## 학습 순서

| # | 폴더 | 주제 | 핵심 학습 |
|---|------|------|-----------|
| 01 | `01_fizzbuzz` | 반복문/조건문 | `for`, `switch`, `fmt.Println` |
| 02 | `02_slice_rotate` | 슬라이스 조작 | `append`, `copy`, slicing |
| 03 | `03_map_wordcount` | 맵 | `map[K]V`, `strings.Fields` |
| 04 | `04_string_anagram` | 문자열 | `rune`, `sort`, 문자 빈도 |
| 05 | `05_two_sum` | 알고리즘 + 맵 | 해시맵 활용 패턴 |
| 06 | `06_struct_sort` | 구조체 + 정렬 | `sort.Slice`, struct field 정렬 |
| 07 | `07_interface_shape` | 인터페이스 | 다형성, 메서드 셋 |
| 08 | `08_pointer_linkedlist` | 포인터 | `*T`, 자기참조 구조체 |
| 09 | `09_error_custom` | 에러 핸들링 | `errors.As`, `fmt.Errorf("%w", ...)` |
| 10 | `10_goroutine_sum` | 고루틴 + WaitGroup | `go`, `sync.WaitGroup` |
| 11 | `11_channel_workerpool` | 채널 + 워커 풀 | `chan`, `select`, `close` |
| 12 | `12_generic_filter` | 제네릭 (Go 1.18+) | type parameter, constraint |

## 사용 방법

각 폴더의 `main.go`를 열어 `// TODO:` 부분을 구현하세요.

```bash
# 기본 문제
go run ./problems/01_fizzbuzz/

# 같은 패턴의 추가 문제
go run ./problems/01_fizzbuzz/extra1/
go run ./problems/01_fizzbuzz/extra2/
```

각 문제 파일에는 `main()` 안에 테스트 케이스가 들어 있어, 구현 후 실행하면 기대값과 비교한 결과가 출력됩니다.

## 폴더 구조

각 항목별로 **3개 문제** (기본 1 + 같은 패턴 응용 2):

```
01_fizzbuzz/
  main.go           ← 기본 (FizzBuzz)
  SOLUTION.md
  extra1/main.go    ← 같은 패턴 응용 1 (Collatz)
  extra1/SOLUTION.md
  extra2/main.go    ← 같은 패턴 응용 2 (자릿수 합 소수)
  extra2/SOLUTION.md
```

**전체 문제 = 12 항목 × 3 = 36 개**

## 항목별 추가 문제 요약

| # | extra1 | extra2 |
|---|--------|--------|
| 01 | Collatz 수열 길이 | 자릿수 합이 소수인 수 카운트 |
| 02 | 2D 슬라이스 Flatten | 슬라이스 Chunk 분할 |
| 03 | 첫 등장 인덱스 매핑 | 첫 글자별 그룹화 |
| 04 | 회문 판별 (양 끝 투포인터) | Run-Length Encoding |
| 05 | 가장 긴 중복 없는 부분문자열 | Group Anagrams |
| 06 | 학생 점수 랭킹 | 이벤트 시간순 정렬 (3-key) |
| 07 | Animal 인터페이스 (Sound, Name) | Storage 인터페이스 (Memory/Null 모킹) |
| 08 | 두 정렬 리스트 합치기 (dummy) | 중간 노드 (slow/fast) |
| 09 | 다중 에러 타입 분기 (errors.As) | 에러 체인 깊이 (errors.Unwrap) |
| 10 | 병렬 Max 찾기 | sync.Mutex 안전 카운터 |
| 11 | 3-stage Pipeline | Fan-In (멀티 producer) |
| 12 | 제네릭 Set[T comparable] | 제네릭 Min/Max (cmp.Ordered) |

## 정답 및 해설

각 문제 폴더에 `SOLUTION.md` 가 있습니다. **먼저 본인이 풀어본 뒤** 열어보는 걸 권장합니다.

해설에는 다음이 포함되어 있어요:
- 정답 코드 전체
- Go 관용구/패턴 설명 (왜 이렇게 쓰는지)
- 다른 풀이 방법
- 자주 하는 실수와 증상

## 학습 팁

- 먼저 **시그니처와 자료형**을 보고 어떤 자료구조를 쓸지 결정
- 표준 라이브러리(`strings`, `sort`, `slices`, `errors`, `sync`)를 적극 활용
- `go vet`, `gofmt`로 코드 검증
- 동시성 문제는 `go run -race` 로 race condition 점검
