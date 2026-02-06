### if 초기문; 조건문 정리

```
if filename, success := UploadFile(); success {
    fmt.Println("Upload success", filename)
} else {
    fmt.Println("Failed to upload")
}
```

이 문법은 if문 안에서 변수 선언(초기문) + 조건 검사를 한 번에 처리하는 방식이다.

같은 로직의 일반 형태

```
filename, success := UploadFile()

if success {
    fmt.Println("Upload success", filename)
} else {
    fmt.Println("Failed to upload")
}
```

두 코드는 동작 로직이 같다.
차이는 변수의 범위(scope) 와 가독성/관리 방식이다.

---

### 차이점 핵심

1. 범위(scope)
- if 초기문; 조건문에서 선언한 filename, success는
if-else 블록 안에서만 유효하다.
	•	바깥에서 재사용할 필요가 없으면 이 방식이 더 안전하고 깔끔하다.

2.  코드 의도 표현
- “이 변수는 이 조건 판단을 위해서만 사용한다”는 의도가 분명해진다.
- 불필요하게 바깥 스코프를 오염시키지 않는다.

3. 언제 분리해서 쓰나?
- filename, success를 if문 이후에도 계속 써야 한다면 선언을 바깥으로 빼는 방식이 맞다.

---

### 한 줄 요약

if 초기문; 조건문은 같은 로직을 더 좁은 범위에서 안전하게 처리하는 Go 스타일 문법이다.