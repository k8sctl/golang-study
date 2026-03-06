package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Seoul")
	const longForm = "Jan 2, 2006 at 3:04pm"
	t1, _ := time.ParseInLocation(longForm, "Jun 10, 2021 at 10:00pm", loc)
	fmt.Println(t1, t1.Location(), t1.UTC())

	const shortForm = "2006-Jan-02"
	t2, _ := time.Parse(shortForm, "2021-Jun-14") // UTC 타임존을 기준으로
	fmt.Println(t2, t2.Location())

	t3, err := time.Parse("2021-06-01 15:20:21", "2021-06-14 20:04:05")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3, t3.Location())

	d := t2.Sub(t1)
	fmt.Println(d)
}

/*
[코드 설명]
이 코드는 문자열로 된 날짜/시간을 time.Time으로 파싱하고,
타임존(Location)과 UTC 변환, 그리고 두 시각의 차이(Sub)를 확인하는 예제다.

1) Asia/Seoul 타임존으로 파싱하기 (ParseInLocation)
- time.LoadLocation("Asia/Seoul")로 서울 타임존 정보를 로드한다.
- longForm은 Go의 레이아웃(layout) 규칙(기준 시간: "Jan 2, 2006 at 3:04pm")을 따른 포맷 문자열이다.
- time.ParseInLocation(layout, value, loc)은
  "입력 문자열에 타임존 정보가 없을 때" loc(여기서는 Asia/Seoul) 기준으로 해석한다.
- t1은 "Jun 10, 2021 at 10:00pm"을 서울 시간(KST, +0900) 기준으로 파싱한 결과다.
- t1.UTC()는 같은 순간을 UTC로 변환한 값이다.
  (KST는 UTC보다 9시간 빠르므로, 22:00 KST -> 13:00 UTC로 변환된다)

2) time.Parse는 기본적으로 UTC로 파싱한다
- shortForm = "2006-Jan-02" 레이아웃으로 "2021-Jun-14"를 파싱한다.
- 입력 문자열에 타임존 정보가 없으므로, time.Parse는 기본적으로 UTC로 해석한다.
- 그래서 t2는 "2021-06-14 00:00:00 +0000 UTC"가 된다.

3) t3 파싱이 실패하는 이유 (layout을 잘못 지정)
- time.Parse(layout, value)에서 layout은 "실제 날짜값"을 쓰는 게 아니라,
  반드시 Go의 기준 시간(2006-01-02 15:04:05)을 이용해 포맷을 표현해야 한다.
- 그런데 아래 코드는 layout에 "2021-06-01 15:20:21"처럼 실제 날짜를 넣었기 때문에 파싱이 실패한다.
  t3, err := time.Parse("2021-06-01 15:20:21", "2021-06-14 20:04:05")
- err가 출력되고, t3는 파싱 실패 시의 제로값(time.Time의 기본값)이 된다:
  0001-01-01 00:00:00 +0000 UTC

  (참고) 의도대로 파싱하려면:
  t3, err := time.Parse("2006-01-02 15:04:05", "2021-06-14 20:04:05")

4) 두 시각의 차이 계산(Sub)
- d := t2.Sub(t1)은 "t2 - t1"의 시간 차이(Duration)를 계산한다.
- 내부 계산은 절대 시각 기준(UTC 기준)으로 맞춰서 수행된다.

  여기서:
  - t1 = 2021-06-10 22:00 KST = 2021-06-10 13:00 UTC
  - t2 = 2021-06-14 00:00 UTC
  따라서 d = (2021-06-14 00:00) - (2021-06-10 13:00) = 83h0m0s

[출력에서 보게 되는 것]
- 첫 줄: t1(서울 기준) / t1의 Location / t1의 UTC 변환 값
- 둘째 줄: t2(UTC 기준) / t2의 Location(UTC)
- 셋째 줄: t3 파싱 에러 메시지
- 넷째 줄: t3 제로값과 Location(UTC)
- 마지막 줄: t2와 t1의 시간 차이(Duration)
*/

/*
[출력 결과]
2021-06-10 22:00:00 +0900 KST Asia/Seoul 2021-06-10 13:00:00 +0000 UTC
2021-06-14 00:00:00 +0000 UTC UTC
parsing time "2021-06-14 20:04:05" as "2021-06-01 15:20:21": cannot parse "-06-14 20:04:05" as "1"
0001-01-01 00:00:00 +0000 UTC UTC
83h0m0s
*/
