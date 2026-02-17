# 구조체 포인터 초기화 & 인스턴스 라이프사이클 통합 정리

## 1) 구조체 포인터 초기화

### (1) 이미 만든 구조체 변수의 주소 받기

```go
var data Data
var p *Data = &data
```

- data: Data 타입 값 변수
- p: data의 주소를 담는 *Data 포인터

### (2) 리터럴로 만들면서 바로 포인터 받기

```go
var p *Data = &Data{}
```

- Data{} 인스턴스를 만들고 그 주소를 p에 저장
- 필드는 제로값으로 초기화 (int=0, string="", bool=false)

> 둘 다 p의 타입은 *Data 차이는 "값 변수(data)를 먼저 만들었는지" 여부

---

## 2) 인스턴스(Instance)란?

- 인스턴스는 메모리에 실제로 생성된 데이터 실체
- `var data Data`를 실행하면 Data 인스턴스 생성
- `var p *Data = &data`는 그 인스턴스를 p가 참조하게 만듦
- p를 통해 인스턴스 필드 값을 읽고 수정할 수 있음

---

## 3) 인스턴스 개수 예시

### (1) 인스턴스 1개 + 포인터 3개

```go
var p1 *Data = &Data{}
var p2 *Data = p1
var p3 *Data = p1
```

- Data 인스턴스는 1개
- p1, p2, p3는 같은 주소(같은 인스턴스)를 가리킴
- p2/p3로 수정한 값이 p1에서도 동일하게 보임

### (2) 인스턴스 3개 (값 복사)

```go
var data1 Data
var data2 Data = data1
var data3 Data = data1
```

- data2, data3는 data1의 값 복사본(별도 인스턴스)
- 값이 처음엔 같아도 메모리는 다름
- data2 수정이 data1, data3에 영향 없음

---

## 4) new() 내장 함수

```go
p1 := &Data{}       // 리터럴 + 주소 연산자(&)
var p2 = new(Data)  // new()로 *Data 생성
p3 := new(Data)     // := 사용 가능 (타입 추론)
```

- new(Data)는 제로값 Data 인스턴스의 주소(*Data)를 반환
- p2, p3는 기본값 초기화 방식

### 초기값 지정 차이

- `&Data{...}`: 생성 시점 초기값 지정 가능
  ```go
  p1 := &Data{value: 10}
  ```
- `new(Data)`: 생성 시점 필드값 직접 지정 문법 없음
  ```go
  p2 := new(Data) // 기본값 생성
  p2.value = 10   // 생성 후 대입
  ```

---

## 5) 인스턴스는 언제 사라질까?

- 함수 종료 시 지역 변수(예: u)는 스코프를 벗어나 소멸
- 인스턴스는 즉시 삭제되지 않음
- 더 이상 도달 가능한 참조가 없으면 "도달 불가능(unreachable) 객체"가 됨
- unreachable 객체는 이후 GC 실행 시 메모리 회수 대상

### 예시

```go
func TestFunc() {
    u := &User{} // User 인스턴스 생성, u가 참조
    u.Age = 30
    fmt.Println(u)
} // 함수 종료 -> u 소멸
  // 다른 참조 경로가 없으면 unreachable
  // 다음 GC 사이클에서 회수 대상
```

---

## 핵심 요약

- 포인터는 인스턴스의 주소를 참조한다.
- 값 복사는 새 인스턴스를 만든다.
- `&Data{}`와 `new(Data)`는 둘 다 *Data를 만들지만, 초기값 지정 가능 여부가 다르다.
- 인스턴스는 참조를 잃어 unreachable이 되면 GC 시점에 메모리에서 회수된다.