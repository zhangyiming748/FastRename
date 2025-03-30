package main

import (
	mylog "FastRename/log"
	"FastRename/sql"
)

func init() {
	mylog.SetLog()
	sql.SetLevelDB()
}
func main() {

}
