package rename

import (
	"FastRename/util"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func RenameNeteaseMusic(root string) []string {
	files, err := util.GetAllFiles(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		log.Printf("file: %v\n", file)
		dir := filepath.Dir(file)
		base := filepath.Base(file)
		log.Printf("dir: %v\n", dir)
		log.Printf("base: %v\n", base)
		if strings.Contains(base, "-") {
			newName := strings.Split(base, "-")[1]
			newName = strings.Replace(newName, " ", "", 1)
			log.Printf("newName: %v\n", newName)
			newPath := filepath.Join(dir, newName)
			log.Printf("newPath: %v\n", newPath)
			log.Printf("最终重命名之前的报告: 原文件路径 = %v, 新文件路径 = %v\n", file, newPath)
			os.Rename(file, newPath)
		}
	}
	return files
}
