package main

import (
	mylog "FastRename/log"
	"FastRename/rename"
	"FastRename/sql"
	"FastRename/util"
	"log"
	"os"
	"strings"
)

func init() {
	mylog.SetLog()
	sql.SetLevelDB()
}
func main() {
	files := util.GetAllFiles("/mnt/d/tars")
	for _, file := range *files {
		newPugreName := rename.Replace(file.PurgeName)
		log.Printf("替换前:%v 替换后:%v\n", file.PurgeName, newPugreName)
		newFullName := strings.Join([]string{file.PurgePath, newPugreName}, "")
		newFullName = strings.Join([]string{newFullName, file.PurgeExt}, ".")
		log.Printf("替换前绝对路径:%v 替换后绝对路径:%v\n", file, newFullName)
		sql.GetLevelDB().Put([]byte(file.FullName), []byte(newFullName), nil)
		os.Rename(file.FullPath, newFullName)
	}
}
