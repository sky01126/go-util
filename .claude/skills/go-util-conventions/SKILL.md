---
name: go-util-conventions
description: >
    go-util 라이브러리 고유 규약 — 평평한 패키지 레이아웃, 의존성 최소화 원칙,
    표준 testing 기반 테스트 스타일(testify 미사용), 에러 래핑·nil-safe 리시버
    패턴, 한국어 GoDoc 규칙, VERSION go:embed 동기화를 다룬다. go-util 저장소에서
    코드를 추가·리뷰·테스트할 때 글로벌 go-coding-standards 위에 얹어 참조한다.
---

# go-util 고유 규약

이 문서는 `github.com/sky01126/go-util` 저장소에만 적용되는 **delta**다. 일반 Go
규칙은 글로벌 `go-coding-standards`·`go-coding-style`을 따르고, 문서·GoDoc·커밋
규칙은 저장소 루트의 `@.junie/guidelines.md`를 따른다. 충돌 시 이 문서가 우선한다.

## 성격: 재사용 라이브러리

- 실행 파일이 아닌 **라이브러리**다. `cmd/`·`internal/`·`pkg/` 레이아웃을 쓰지
  않고, 기능 그룹별로 **루트 직속 평평한 패키지 디렉터리**를 둔다
  (`compress/`, `files/`, `logger/`, `strings/`).
- 새 유틸리티 그룹은 새 디렉터리 + 같은 이름의 패키지로 추가한다. 패키지명은
  디렉터리명과 일치시키고 소문자 단수형으로 짓는다.

## 의존성 최소화

- 외부 의존성 추가는 신중히 한다. 현재 직접 의존성은 `zap`(logger)과
  `archiver`(compress)뿐이다.
- **테스트에 testify를 도입하지 않는다.** 표준 `testing` 패키지만 사용한다
  (아래 테스트 스타일 참조). 글로벌 `go-test` 커맨드는 testify를 전제하지만,
  이 저장소에서는 `/util-test`를 쓴다.

## 테스트 스타일 (표준 testing)

실제 저장소 컨벤션은 테이블 기반·testify가 아니라 **표준 `testing` + 수동 if 검증
+ `t.Errorf`**다. `strings/strings_test.go`를 기준 예시로 삼는다.

- 함수명: `TestXxx` (대상 함수명과 일치)
- 한 테스트 함수 안에서 여러 케이스를 연속 if 문으로 검증
- 실패 시 `t.Errorf`로 어떤 입력이 왜 실패했는지 메시지에 담는다
- mock·assert 라이브러리를 쓰지 않는다
- 외부 자원(파일 등)이 필요하면 `t.TempDir()`·`os` 표준 API로 처리

```go
func TestIsEmpty(t *testing.T) {
	if !IsEmpty("") {
		t.Errorf("IsEmpty(\"\") should be true")
	}
	if IsEmpty("abc") {
		t.Errorf("IsEmpty(\"abc\") should be false")
	}
}
```

## 함수·에러 패턴

- 공개 함수는 `(T, error)` 또는 `error`로 단순하게 유지한다.
- 에러는 `fmt.Errorf("동작 설명 %s: %w", arg, err)`로 래핑한다. **에러 문자열은
  영문 소문자**로 시작하고 마침표를 붙이지 않는다(`"read file %s: %w"`).
- 리소스 정리 실패도 에러로 전파한다: named return + `defer`에서 `Close()` 에러를
  병합하는 패턴(`files.Copy` 참조).
- nil 수신자에서도 안전하게 동작시킬 수 있으면 그렇게 한다
  (`CasePolicy.Equals`는 nil이면 대소문자 구분 동작으로 폴백).
- 라이브러리 코드는 직접 로깅하지 않는다. 맥락을 에러에 담아 호출자에게 넘긴다.

## GoDoc (요약, 상세는 @.junie/guidelines.md)

- 공개 식별자에 GoDoc 작성. **첫 단어는 식별자명**으로 시작: `Exists 파일 또는 ...`
- 한국어 단답형 한 문장. 예제 코드 금지, HTML·JavaDoc 태그 금지.
- 금지 단어: `주어진`, `지정된`. 구조체 필드 문서는 작성하지 않는다.

## VERSION 동기화

- 루트 `VERSION` 파일이 `version.go`에서 `//go:embed`로 읽혀 `go_util.Version`이
  된다. 버전 변경은 코드가 아니라 **`VERSION` 파일**을 수정한다.
- 버전을 올릴 때 README의 버전 배지도 함께 갱신한다. 커밋 타입은 `chore(version)`.

## 새 기능 추가 시 산출물 순서

1. 패키지 코드 (`<pkg>/<pkg>.go`) — 공개 함수 + 한국어 GoDoc
2. 테스트 (`<pkg>/<pkg>_test.go`) — 표준 testing 스타일
3. README 사용법 섹션 갱신 (해당 패키지 import 경로·짧은 코드 예시)
