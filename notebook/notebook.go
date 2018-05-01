package notebook

import (
	"os"
	"path/filepath"
)

// 笔记本全局变量
var pathsep string

type stBookDir struct {
	base string
	date string
	tag  string
	blog string
	temp string
}
type stBookExt struct {
	note string
	temp string
	tag  string
}

var bookdir stBookDir
var bookext stBookExt

// 初始化各种目录后缀名
func init() {
	pathsep = string(os.PathSeparator)
	bookdir.base = os.Getenv("HOME") + pathsep + "notebook"

	bookdir.date = "d"
	bookdir.tag = "t"
	bookdir.blog = "b"
	bookdir.temp = "htpl"

	bookext.note = ".md"
	bookext.temp = ".html"
	bookext.tag = ".tag"
}

// 设置基础目录
func SetBaseDir(dir string) (old string) {
	old = bookdir.base
	if dir != "" {
		bookdir.base = dir
	}
	return old
}

func PathSep() string {
	return pathsep
}

func NoteExt() string {
	return bookext.note
}

// 返回可开放浏览的静态目录
func StaticDir() []string {
	var dir = []string{bookdir.date, bookdir.tag, bookdir.blog}
	return dir
}

// 获取一个模板文件的全路径
func TemplateFile(name string) string {
	file := name + bookext.temp
	return filepath.Join(bookdir.base, bookdir.temp, file)
}

func TagFile(name string) string {
	file := name + bookext.tag
	return filepath.Join(bookdir.base, bookdir.tag, file)
}

func NoteFile(name string) string {
	file := name + bookext.note
	return filepath.Join(bookdir.base, bookdir.date, file)
}
