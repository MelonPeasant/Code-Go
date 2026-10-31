package main

import (
	"fmt"
	"io"
	"os"
)

// 日志文件名称
const logFile = "iplog.txt"

// 读取日志最后一行，拿到上一次记录的IP
func getLastIP() (string, error) {
	file, err := os.Open(logFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil // 文件不存在，说明没有历史记录
		}
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	lines := splitLines(string(content))
	if len(lines) == 0 {
		return "", nil
	}
	lastLine := lines[len(lines)-1]
	// 格式示例：2026-10-30 14:30:00 | 123.45.67.89
	var tStr, ip string
	_, err = fmt.Sscanf(lastLine, "%s %s | %s", &tStr, &tStr, &ip)
	if err != nil {
		return "", nil
	}
	return ip, nil
}

func splitLines(s string) []string {
	var res []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			res = append(res, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		res = append(res, s[start:])
	}
	return res
}

func main() {
	lastIP, err := getLastIP()
	if err != nil {
		fmt.Println("读取历史记录失败：", err)
		return
	} else {
		fmt.Println(lastIP)
	}

}
