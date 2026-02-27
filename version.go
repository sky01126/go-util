package go_util

import (
	_ "embed"
	"strings"
)

//go:embed VERSION
var versionFromFile string

// Version 프로젝트의 현재 버전을 정의한다.
var Version = strings.TrimSpace(versionFromFile)
