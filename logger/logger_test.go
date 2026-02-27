package logger

import (
	"context"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestInitLogger(t *testing.T) {
	InitLogger("")
	if L() == nil {
		t.Error("로거가 초기화되지 않았습니다")
	}
}

func TestSetLevel(t *testing.T) {
	InitLogger("")
	SetLevel(zapcore.DebugLevel)
	if atom.Level() != zapcore.DebugLevel {
		t.Errorf("로깅 레벨이 설정되지 않았습니다. 기대값: %v, 실제값: %v", zapcore.DebugLevel, atom.Level())
	}

	SetLevel(zapcore.ErrorLevel)
	if atom.Level() != zapcore.ErrorLevel {
		t.Errorf("로깅 레벨이 변경되지 않았습니다. 기대값: %v, 실제값: %v", zapcore.ErrorLevel, atom.Level())
	}
}

func TestContext(t *testing.T) {
	InitLogger("")
	ctx := context.Background()

	// 기본 로거 반환 확인
	l1 := FromContext(ctx)
	if l1 != L() {
		t.Error("컨텍스트에 로거가 없을 때 기본 로거를 반환해야 합니다")
	}

	// 로거 주입 및 추출 확인
	namedLogger := L().Named("test")
	ctx = WithContext(ctx, namedLogger)
	l2 := FromContext(ctx)
	if l2 != namedLogger {
		t.Error("컨텍스트에서 주입된 로거를 추출하지 못했습니다")
	}
}

func TestConvenienceMethods(t *testing.T) {
	InitLogger("")
	// 실행 시 패닉이 발생하지 않는지 확인
	Debug("디버그 메시지", zap.String("key", "value"))
	Info("정보 메시지")
	Warn("경고 메시지")
	Error("에러 메시지")
	Sync()
}
