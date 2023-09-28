package cachebench

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// GetMainDirectory 获取进程所在目录: 末尾带反斜杠
func GetMainDirectory() string {
	path, err := filepath.Abs(os.Args[0])

	if err != nil {
		return ""
	}

	fullPath := filepath.Dir(path)
	return pathAddBackslash(fullPath)
}

// pathAddBackslash 路径最后添加反斜杠
func pathAddBackslash(path string) string {
	i := len(path) - 1

	if !os.IsPathSeparator(path[i]) {
		path += string(os.PathSeparator)
	}
	return path
}

func GetGoroutineID() int64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Failed to get Goroutine ID:", r)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idStr := string(buf[:n])

	var id int64
	fmt.Sscanf(idStr, "goroutine %d", &id)

	return id
}
