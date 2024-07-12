package sql

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os"
	"strings"
)

var levelDB *leveldb.DB

func SetLevelDB() {
	location := strings.Join([]string{"leveldb"}, string(os.PathSeparator))
	db, err := leveldb.OpenFile(location, nil)
	if err != nil {
		log.Fatalf("leveldb数据库创建失败:%v\n", err)
	}
	levelDB = db
}
func GetLevelDB() *leveldb.DB {
	return levelDB
}
