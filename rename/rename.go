package rename

import (
	"strings"
)

var Keywords = []string{
	"4K",
	"4K",
}

func Replace(oldName string) string {
	oldName = strings.ReplaceAll(oldName, "4K超高清HiRes", "")
	oldName = strings.ReplaceAll(oldName, "顶级修复HiRes", "")
	oldName = strings.ReplaceAll(oldName, "4K顶级修复", "")
	oldName = strings.ReplaceAll(oldName, "无损封装", "")
	oldName = strings.ReplaceAll(oldName, "8K顶级画质", "")
	oldName = strings.ReplaceAll(oldName, "4K超清", "")
	oldName = strings.ReplaceAll(oldName, "4K重制", "")
	oldName = strings.ReplaceAll(oldName, "HiRes", "")
	oldName = strings.ReplaceAll(oldName, "修复版", "")
	oldName = strings.ReplaceAll(oldName, "修复", "")
	oldName = strings.ReplaceAll(oldName, "4K", "")
	oldName = strings.ReplaceAll(oldName, "全308集禁止自学走弯路清华大佬耗费196小时录制的golang零基础入门教程Java程序员转行golang开发必看go教程go项目go实战", "")
	oldName = strings.TrimSpace(oldName)
	return oldName
}
