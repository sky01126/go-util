# go-util

[![Go Version](https://img.shields.io/badge/Go-1.25-blue.svg)](https://golang.org)
[![Version](https://img.shields.io/badge/version-0.1.0-green.svg)](https://github.com/sky01126/go-util)

이 프로젝트는 다양한 유틸리티 함수들을 모아놓은 Go 라이브러리입니다. 지속적으로 유용한 유틸리티들을 추가하는 것을 목표로 합니다.

## 설치 방법

```bash
go get github.com/sky01126/go-util
```

## 사용 방법

현재 제공되는 유틸리티 그룹은 다음과 같습니다:

### 버전 확인

라이브러리의 버전을 확인하는 방법은 다음과 같습니다.

#### 1. 코드 내에서 확인
프로그램 내에서 라이브러리의 버전을 참조할 수 있습니다.

```go
import "github.com/sky01126/go-util"

fmt.Println(go_util.Version) // "0.1.1"
```

#### 2. VERSION 파일 확인
프로젝트 루트 디렉토리에 있는 `VERSION` 파일을 통해 확인할 수 있습니다.

```bash
cat VERSION
```

### 1. 제공 유틸리티

기본적인 파일 및 문자열 조작 유틸리티입니다.

#### Strings 유틸리티
문자열 비교, 공백 확인, 축약 등을 제공합니다.

```go
import "github.com/sky01126/go-util/strings"

// 문자열 축약
abbreviated := strings.Abbreviate("Hello World", 8) // "Hello..."

// 대소문자 무관 비교
isEqual := strings.CI.Equals("abc", "ABC") // true

// 기본값 반환
val := strings.DefaultIfEmpty("", "default") // "default"
```

#### Files 유틸리티
파일 읽기/쓰기, 존재 여부 확인, 복사 등을 제공합니다.

```go
import "github.com/sky01126/go-util/files"

// 파일 존재 확인
exists := files.Exists("test.txt")

// 파일 읽기
content, err := files.ReadString("hello.txt")

// 파일 쓰기
err = files.WriteString("hello.txt", "Hello Go!")
```

#### Compress 유틸리티
파일 및 디렉토리의 압축/해제를 제공합니다.

```go
import "github.com/sky01126/go-util/compress"

// 압축 파일 해제
err := compress.Uncompress("archive.zip", "./extracted")

// 파일 압축
err = compress.Compress([]string{"file1.txt", "dir1"}, "archive.zip")
```

#### Logger 유틸리티
`zap` 기반의 고성능 로깅 기능을 제공합니다. 전역 로깅, 컨텍스트 기반 로깅, 동적 레벨 변경 등을 지원합니다.

```go
import (
    "github.com/sky01126/go-util/logger"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// 로거 초기화 (APP_ENV=production 인 경우 파일 로깅 활성화 가능)
logger.InitLogger("app.log")
defer logger.Sync()

// 기본 로깅
logger.Info("애플리케이션 시작", zap.String("version", "1.0.0"))
logger.Error("에러 발생", zap.Error(err))

// 컨텍스트 기반 로깅
ctx := logger.WithContext(ctx, logger.With(zap.String("request_id", "req-123")))
l := logger.FromContext(ctx)
l.Info("요청 처리 완료")

// 동적 로그 레벨 변경
logger.SetLevel(zapcore.DebugLevel)
```

## 프로젝트 구조

```text
go-util/
├── compress/           # 압축 관련 유틸리티
├── files/              # 파일 관련 유틸리티
├── logger/             # 로깅 관련 유틸리티
├── strings/            # 문자열 관련 유틸리티
├── go.mod
├── README.md
└── VERSION
```

## 라이선스

MIT License
