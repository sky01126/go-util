# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 프로젝트 개요

`github.com/sky01126/go-util` — 재사용 가능한 Go 유틸리티 **라이브러리**(실행 가능한 앱 아님). Go 1.25.0.
패키지별로 디렉터리가 나뉜다: `compress/`(압축·해제), `files/`(파일 IO), `logger/`(zap 기반 로깅), `strings/`(문자열 유틸), 루트 `version.go`(버전).

## 코딩 가이드라인

상세한 코드 작성·문서·커밋 규칙은 아래 파일을 따른다(중복 금지):

@.junie/guidelines.md

핵심만 요약하면: 코드(식별자)는 영문, GoDoc 주석·커밋 메시지·설명은 한국어. 커밋은 `<type>(<scope>): <한국어 설명>` (Conventional Commits, 명령형, 마침표·볼드 금지).

## 빌드 / 테스트

Makefile 없음. 표준 Go 명령을 사용한다.

```bash
go build ./...
go test ./...
go test ./logger/        # 단일 패키지
go test -run TestName ./logger/   # 단일 테스트
```

## 비자명한 동작 (gotcha)

- **버전 동기화:** `version.go`가 루트 `VERSION` 파일을 `//go:embed`로 읽어 `Version` 변수에 담는다. 버전을 올릴 때는 코드가 아니라 `VERSION` 파일을 수정한다(README의 버전 배지도 함께 갱신).
- **로거 파일 로깅:** `logger.InitLogger`는 `APP_ENV` 환경 변수가 비어 있거나 `"local"`이 아닐 때만 `logPath` 파일에 로깅한다. 로컬 개발에서는 `APP_ENV=local`로 콘솔 전용 로깅이 된다.
