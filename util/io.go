package util

import (
	"bufio"
	"os"
	"path/filepath"
)

type FileInfo struct {
	FullPath  string // 文件的绝对路径
	FullName  string // 完整的文件名
	PurgePath string // 纯路径名 带有最后一个路径分隔符
	PurgeName string // 相对路径的纯文件名
	PurgeExt  string // 不带.的扩展名
}

/*
golang实现获取指定目录下包括子文件夹下全部文件名并保存到字符串切片
*/

func GetAllFiles(dirPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

/*
按行读文件
*/
func ReadByline(fp string) ([]string, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
