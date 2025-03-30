package test

import (
	"FastRename/rename"
	"testing"
)
// go test -v -timeout 20h -run TestNetEase
func TestNetEase(t *testing.T) {
	fp := "/Users/zen/Downloads/Music"
	_ = rename.RenameNeteaseMusic(fp)

}
