package notebook

import (
	"os"
	"path/filepath"
)

// 笔记本全局变量
var pathsep string
var bookdir = struct {
	base string
	date string
	tag  string
	blog string
	temp string
}{}

var bookext = struct {
	note string
	temp string
}{}

func init() {
	pathsep = string(os.PathSeparator)
	bookdir.base = os.Getenv("HOME") + pathsep + "notebook"

	bookdir.date = "d"
	bookdir.tag = "t"
	bookdir.blog = "b"
	bookdir.temp = "htpl"

	bookext.note = ".md"
	bookext.temp = ".html"
}

// 设置基础目录
func BookDir(dir string) (old string) {
	old = bookdir.base
	if dir != "" {
		bookdir.base = dir
	}
	return old
}

// 返回可开放浏览的静态目录
func StaticDir() []string {
	var dir = []string{bookdir.date, bookdir.tag, bookdir.blog}
	return dir
}

func PathSep() string {
	return pathsep
}

// 获取一个模板文件的全路径
func TemplateFile(name string) string {
	file := name + bookext.temp
	return filepath.Join(bookdir.base, bookdir.temp, file)
}
