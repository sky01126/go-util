---
description: go-util에 새 유틸리티 함수/패키지를 저장소 규약대로 추가 (코드+테스트+README)
---

go-util 라이브러리에 새 유틸리티를 추가해줘: $ARGUMENTS

## 참조 스킬

작업 전 다음을 반드시 참조할 것:

- `go-util-conventions` — 이 저장소 고유 규약(평평한 패키지 레이아웃, 의존성 최소화,
  표준 testing 테스트 스타일, 에러 래핑·nil-safe 패턴, VERSION 동기화)
- `go-coding-standards` — 일반 Go 설계·에러·동시성 원칙
- `go-coding-style` — 네이밍, import 순서, GoDoc
- 저장소 루트의 `@.junie/guidelines.md` — 문서·GoDoc·커밋 규칙

## 진행 절차

1. **대상 패키지 결정**: 요청이 기존 그룹(`compress`/`files`/`logger`/`strings`)에
   맞으면 해당 디렉터리에 함수를 추가한다. 새 그룹이면 루트에 새 디렉터리 +
   같은 이름의 소문자 단수 패키지를 만든다. 모호하면 사용자에게 묻는다.
2. **함수 작성**: 공개 함수는 `(T, error)`/`error`로 단순하게. 에러는
   `fmt.Errorf("...: %w", err)`(영문 소문자) 래핑. nil·경계값을 고려. 직접 로깅 금지.
3. **GoDoc**: 첫 단어 = 식별자명, 한국어 단답형. 예제·`주어진`/`지정된`·HTML 태그 금지.
4. **테스트 작성**: 기존 `*_test.go`와 동일한 **표준 testing + t.Errorf** 스타일.
   testify·테이블 기반·mock 도입 금지. 정상·경계·에러 경로를 각각 검증.
5. **README 갱신**: "제공 유틸리티" 섹션에 import 경로와 짧은 사용 예를 추가한다.
6. **검증**: `go build ./...`, `go test ./...`, (설치돼 있으면) `golangci-lint run ./...`.

## 마무리

추가한 함수·파일과 커버한 테스트 케이스를 한국어로 간단히 요약하고,
`feat(<scope>): ...` 형식의 한국어 커밋 메시지를 제안할 것(커밋은 사용자가 지시할 때만).
