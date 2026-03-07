# 숫자 맞히기 게임: 랜덤, 시드, 시간(time) 정리

## 숫자 맞히기 게임 규칙

1. 0 ~ 99 사이의 랜덤 숫자를 1개 만든다.
2. 사용자가 숫자를 입력한다.
3. 입력값이 크면 **"큽니다"**, 작으면 **"작습니다"**를 출력하고 다시 입력 받는다.
4. 같으면 **"축하합니다"**를 출력하고 종료한다.

---

## 랜덤: math/rand

- `math/rand` 패키지를 사용한다.
- `rand.Intn(n int) int`
  - `0 ~ (n-1)` 범위의 난수를 반환한다.
  - 예) `rand.Intn(100)` → `0 ~ 99`

### `rand.Seed` 관련 경고(Go 1.20+)

IDE/정적 분석 도구에서 아래와 같은 경고가 뜰 수 있다:

> `rand.Seed is deprecated: As of Go 1.20 there is no reason to call Seed with a random value...`

이 말의 핵심은 다음과 같다.

#### 1) Go 1.20부터 전역 난수 생성기는 기본적으로 랜덤 시드로 초기화된다

- 예전에는 `rand.Seed(time.Now().UnixNano())`를 안 하면 실행할 때마다 같은 난수 순서가 시작될 수 있었다.
- Go 1.20+에서는 전역 generator가 기본적으로 랜덤하게 초기화되므로, "실행마다 랜덤"이 목적이라면 보통 **Seed를 굳이 호출할 필요가 없다.**

✅ 가장 단순한 방식(Go 1.20+ 권장):

```go
answer := rand.Intn(100)
```

#### 2) "재현 가능한 난수 시퀀스(고정 시드)"가 필요하면 로컬 RNG를 쓰는 게 권장

테스트/디버깅처럼 항상 같은 순서의 난수가 필요할 때는 전역 rand.Seed(1) 대신, 로컬 난수 생성기를 만들어 사용한다.

```go
r := rand.New(rand.NewSource(1)) // 고정 시드
fmt.Println(r.Intn(100))
```

이 방식의 장점:
- 전역 상태에 영향 없음
- 패키지/테스트 간 간섭이 줄어듦
- 동시성 환경에서도 예측 가능

---

## 무엇을 랜덤 시드로 하면 좋은가?

- 프로그램 실행 시점마다 계속 변하는 값을 시드로 쓰는 것이 일반적이다.
- 보통 현재 시각 기반 값을 시드로 사용한다.

예) 현재 시각 기반 시드:

```go
r := rand.New(rand.NewSource(time.Now().UnixNano()))
answer := r.Intn(100)
```

---

## time 패키지

time 패키지는 시간과 관련된 기능을 제공한다.

- **Time 객체**: 시각(현재 시각, 특정 시각)을 표현
  - `time.Now()`: 현재 시각(Time) 반환
  - `(t Time).UnixNano()`: Unix epoch(1970-01-01 UTC)부터 경과 시간을 나노초(int64)로 반환
- **Duration 객체**: 시간의 길이(예: 1초, 10분)를 표현
  - 예) `time.Second`, `time.Millisecond` 등
- **Location 객체**: 타임존(시간대)을 표현
  - 예) `time.LoadLocation("Asia/Seoul")` 등
