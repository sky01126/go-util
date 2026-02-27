# AI Code Generation Guidelines (Go 버전)

## 1. Language & General Rules
- **Code:** 모든 코드는 영문을 사용한다. (패키지명, 타입/함수/변수명, 상수명 등)
- **Documentation:** 모든 설명, GoDoc(주석), 커밋 메시지는 **한국어**로 작성한다.
- **Consistency:** 문서 스타일과 용어는 프로젝트 전반에서 일관성을 유지한다.
- **Variable Description:** 동일한 구조체 필드(struct field) 및 동일한 의미의 변수는 동일한 설명을 사용한다.

## 1-1. 프로젝트 레이아웃 규칙
- **프로젝트 개발:** 모듈 단위 개발이 아니라 “하나의 애플리케이션/서비스 프로젝트”를 개발하는 경우, 디렉터리 구조는 표준 Go 프로젝트 레이아웃을 따른다  
  - 참고: https://github.com/golang-standards/project-layout/blob/master/README_ko.md
- **레이아웃 원칙:**
  - 실행 진입점은 `cmd/<app-name>/main.go`에 둔다.
  - 재사용 가능한 공개 패키지는 `pkg/`에 둔다.
  - 프로젝트 내부 전용 패키지는 `internal/`에 두고 외부에서 import하지 못하도록 한다.
  - 배포 산출물 및 빌드 결과물은 저장소에 커밋하지 않고(원칙), 필요 시 `build/` 또는 `dist/`로 분리한다.
- **예외 처리:** 프로젝트 성격상 표준 레이아웃을 그대로 적용하기 어려운 경우, “왜 예외가 필요한지”와 “대안 구조”를 문서로 남긴다.

## 2. Code Style & Backend Rules (Go)
- **Readability:** 가독성을 최우선으로 하며, 유지보수가 용이한 명확한 코드를 작성한다.
- **Simplicity:** 불필요한 최신 문법이나 과한 추상화는 사용하지 않는다. 표준 라이브러리와 Go 관용구(Go idioms)를 우선한다.
- **Error Handling:** 에러 처리는 로직과 명확하게 분리하며, 필요 시 `fmt.Errorf("...: %w", err)`로 감싸 호출자에게 충분한 맥락을 제공한다.
- **Safety & Logging:**
  - nil 가능성(포인터, `interface{}`, `map`/`slice`, 채널, `error`)을 항상 고려한다.
  - 라이브러리 코드는 원칙적으로 직접 로깅하지 않고 에러에 맥락을 담아 반환한다. 애플리케이션 코드에서 운영 환경을 고려해 로그 레벨과 메시지 컨텍스트를 결정한다.
- **Naming Convention:**
  - 공개 API는 Go 네이밍 컨벤션(대문자 시작 = exported, 소문자 시작 = unexported)을 따른다.
  - 함수명은 동사를 우선하며, `Get` 접두어는 정말 필요할 때만 사용한다(보통은 `Version()`, `HomeDir()`처럼 명사형도 허용).
  - Bool 반환 함수는 `Is`, `Has`, `Exists`, `Can` 등으로 의미를 명확히 한다.
  - 에러 반환 함수는 가능하면 `(T, error)` 또는 `error` 형태로 단순하게 유지한다.

## 3. GoDoc & Documentation Rules
### 공통 규칙
- **예제 생성 금지:** 예제 코드 및 사용 예제를 생성하지 않는다.
- **금지 단어:** `주어진`, `지정된` 단어 사용을 금지한다.
- **HTML 태그 금지:** GoDoc 주석에 HTML 태그를 사용하지 않는다.
- **태그 사용 제한:** JavaDoc 스타일 태그(`@param`, `@return`, `@throws` 등)는 사용하지 않는다.
- **범위 제한:** 구조체 필드(타입 멤버 속성) 문서는 작성하지 않는다.

### GoDoc 작성 규칙
- 공개(exported) 식별자에는 GoDoc을 작성한다.
- 첫 문장은 반드시 대상 이름으로 시작한다. 예: `CopyFile ...`
- 문장은 간결한 단답형으로 작성한다.

### Deprecated 작성 규칙
- 대체 항목이 있으면 `Deprecated:`로 시작하는 문장을 사용하고, 대체 함수/타입명을 함께 명시한다.

## 4. Model & DTO Rules (Go 관점)
- **Immutability:** 불변성을 강제하기 어렵기 때문에, 외부에 노출되는 구조체는 가능한 한 읽기 전용 사용을 유도한다(필드 비노출, 생성자 함수 제공, 깊은 복사 고려 등).
- **Validation:** 입력 데이터 검증은 경계(HTTP 핸들러, CLI 입력 파서 등)에서 수행하는 것을 원칙으로 한다.
- **Separation:** 외부 노출용 구조체(응답/요청)와 내부 도메인 모델을 분리하여 의존성을 최소화한다.

## 5. Mapping Rules (Go)
- 자동 매핑 프레임워크 전제를 두지 않는다.
- 변환 로직은 명시적으로 작성하고, 네이밍은 `ConvertXToY`, `ParseX`, `NewXFromY` 등으로 의도를 드러낸다.
- 매핑 누락이 위험한 경우, “명시적 필드 할당”을 우선하여 컴파일 타임/리뷰 타임에 누락을 드러낸다.

## 6. Git 커밋 메시지 작성 규칙
### 형식: `<type>(<scope>): <description>`
- **type:** `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`, `perf`, `ci`, `build`, `revert`
- **scope (선택):** `api`, `auth`, `db`, `test`, `deps`, `config`, `cli`, `pkg` 등 프로젝트에 맞게 사용한다.
- **description 작성 규칙:**
  - 모든 커밋 메시지는 **한국어**로 작성한다.
  - 150자 이내, 명령형을 사용한다.
  - 볼드(강조) 및 마침표 사용을 금지한다.

## 7. Communication & Explanation
- **Intent:** 설계 의도("왜 이렇게 작성했는지")를 반드시 설명한다.
- **Legacy:** 기존 레거시 코드와의 호환성을 고려한다.
- **Alternatives:** 대안이 있다면 장단점을 비교 설명한다.
- **Refactoring:** 동작 변경 여부와 성능 영향도를 명시한다.
- **Assumptions:** 전제 조건(Go 버전, 사용 라이브러리, 실행 환경 등)을 명시하며, 불확실한 부분은 질문한다.