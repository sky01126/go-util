---
description: go-util 스타일 단위 테스트 작성 (표준 testing + t.Errorf, testify 미사용)
---

이 코드에 대한 단위 테스트를 **go-util 저장소 스타일**로 작성해줘: $ARGUMENTS

## 참조 스킬

- `go-util-conventions` — 특히 "테스트 스타일" 섹션
- `go-coding-style` — 테스트 함수 네이밍

## 작성 규칙 (글로벌 go-test와 다름)

이 저장소는 testify·테이블 기반 테스트를 쓰지 않는다. 기존 `strings/strings_test.go`,
`files/files_test.go`의 스타일을 그대로 따른다.

1. **표준 `testing` 패키지만 사용.** testify(assert/require/mock) 도입 금지 —
   `go.mod`에 테스트 의존성을 추가하지 않는다.
2. **함수명**: `TestXxx`(대상 함수명과 일치).
3. **검증**: 한 테스트 함수 안에서 케이스별 `if` 조건 + 실패 시 `t.Errorf`.
   메시지에 어떤 입력이 왜 실패했는지 담는다.
4. **파일 등 외부 자원**: `t.TempDir()`·`os` 표준 API로 처리. 실제 네트워크·DB 금지.
5. **결정론적**: `time.Sleep` 금지.

## 커버할 케이스

- **정상 경로**: 가장 흔한 입력
- **경계값**: 빈 문자열, nil, 0, 빈 슬라이스 등 의미 있는 경계
- **에러 경로**: 에러를 반환하는 조건마다 한 케이스
- **nil 수신자**: nil-safe 동작이 의도된 메서드는 nil 케이스도 검증

작성 후 어떤 케이스를 커버했는지 한국어로 요약하고, `go test ./...`로 통과를 확인할 것.
