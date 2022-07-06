package go_utils

import (
	"os"
	"path/filepath"
	"strings"
)

// GetCurrentDir 返回当前运行的目录
func GetCurrentDir() string {
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); nil == err {
		return strings.Replace(dir, "\\", "/", -1) + "/"
	}
	return ""
}
