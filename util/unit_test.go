package util

import "testing"

func TestGetAllFiles(t *testing.T) {
	fp := "C:\\Users\\zen\\Github\\FastRename\\util"
	files := GetAllFiles(fp)
	for _, file := range *files {
		t.Logf("%+v\n", file)
	}
}
