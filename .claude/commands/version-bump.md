---
description: VERSION 파일과 README 배지를 갱신해 라이브러리 버전을 올린다
---

go-util의 버전을 올려줘: $ARGUMENTS (예: `0.1.3`, 또는 `patch`/`minor`/`major`)

## 배경

`version.go`가 루트 `VERSION` 파일을 `//go:embed`로 읽어 `go_util.Version`이 된다.
버전은 **코드가 아니라 `VERSION` 파일**에서 관리한다. 자세한 규약은 `go-util-conventions`
스킬의 "VERSION 동기화" 섹션 참조.

## 진행 절차

1. **현재 버전 확인**: `VERSION` 파일을 읽는다. 인자가 `patch`/`minor`/`major`면
   SemVer 규칙으로 다음 버전을 계산하고, 구체적 버전이면 그 값을 쓴다.
2. **VERSION 갱신**: `VERSION` 파일 내용을 새 버전 문자열로 교체한다(끝 공백·개행 주의).
3. **README 배지 갱신**: `README.md`의 `version-x.y.z-green` 배지를 새 버전으로 맞춘다.
4. **검증**: `go build ./...`로 embed가 깨지지 않았는지 확인한다. 필요하면
   `go test ./version_test.go` 또는 `go test ./...`로 버전 테스트를 돌린다.
5. **요약 및 커밋 제안**: 변경 파일을 요약하고, `chore(version): 애플리케이션 버전 x.y.z로 업데이트`
   형식의 한국어 커밋 메시지를 제안한다(커밋은 사용자가 지시할 때만 실행).

코드(`version.go`)는 수정하지 않는다 — embed 메커니즘은 그대로 둔다.
