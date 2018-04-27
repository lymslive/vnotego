package notebook

import (
	"fmt"
	"github.com/lymslive/vnotego/date"
	"strings"
)

// 笔记本全局变量
var bookdir string = "~/notebook"
var pathsep string = "/"
var extension string = ".md"

// 设置基础目录
func BookDir(dir string) (old string) {
	old = bookdir
	if dir != "" {
		bookdir = dir
	}
	return old
}

// 笔记类型
type NotePost struct {
	NoteID
	Title string
	Tags  []string
	Body  string
}

func (self *NotePost) String() string {
	var str = fmt.Sprintf("%s\t%s", self.NoteID, self.Title)
	if 0 == len(self.Tags) {
		return str
	}

	var tags = strings.Join(self.Tags, "|")
	str = fmt.Sprintf("%s\t|%s|", str, tags)
	return tags
}

// 获取笔记对应的文件名，dir 指定日记的基准目录
func (self *NotePost) File(dir string) string {
	oldSep := date.SepField(date.SEP_PATH)
	defer date.SepField(oldSep)

	file := fmt.SPrintf("%s%s%s%s", self.Date, pathsep, self.NoteID, extension)
	if dir != "" {
		file = dir + pathsep + file
	}

	return file
}
