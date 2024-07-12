package util

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
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
func GetAllFiles(p string) *[]FileInfo {
	files, _ := getAllFiles(p)
	var fis []FileInfo
	for _, file := range files {
		var fi FileInfo
		fi.FullPath = file
		fi.FullName = filepath.Base(file)
		fi.PurgePath = strings.Trim(fi.FullPath, fi.FullName)
		ext := filepath.Ext(fi.FullName)
		fi.PurgeName = strings.TrimSuffix(fi.FullName, ext)
		fi.PurgeExt = strings.Replace(ext, ".", "", 1)
		fis = append(fis, fi)
	}
	return &fis
}
func getAllFiles(dirPath string) ([]string, error) {
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
