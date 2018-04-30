package notebook

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/lymslive/vnotego/markdown"
)

// 笔记标题与标签的正则表达式
const (
	REGEXP_TITLE string = "^[#\\s]+"
	REGEXP_LABEL string = "`[^`]+`"
)

var (
	preg_title *regexp.Regexp = nil
	preg_label *regexp.Regexp = nil
)

func init() {
	preg_title = regexp.MustCompile(REGEXP_TITLE)
	preg_label = regexp.MustCompile(REGEXP_LABEL)
}

// 笔记类型
// Body 其实包含文件所有内容，包括从中解析的标题行与标签
type NotePost struct {
	NoteID
	Title string
	Tags  []string
	Body  []byte
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
func (we *NotePost) File() string {
	file := we.NoteID.String() + bookext.note
	return filepath.Join(bookdir.base, bookdir.date, we.Date.String(), file)
}

// 根据文件名，对应笔记名，寻找笔记，返回笔记对象指针
// 如果没有笔记文件，返回 nil
func FindNote(file string) (*NotePost, error) {
	if dot := strings.LastIndex(file, "."); dot > 0 {
		file = file[:dot]
	}

	var nid *NoteID
	nid, err := Endcode(file)
	if err != nil {
		return nil, err
	}

	var np = new(NotePost)
	np.NoteID = *nid

	path := np.File()
	np.Body, err = ioutil.ReadFile(path)
	if err != nil {
		log.Println("fails to read note file:", path)
		return nil, err
	}

	np.parseNote()
	return np, err
}

// 解析标题与标签
func (we *NotePost) parseNote() error {
	byteReader := bytes.NewReader(we.Body)
	byteSanner := bufio.NewScanner(byteReader)

	var line string
	if byteSanner.Scan() {
		line = byteSanner.Text()
		we.Title = preg_title.ReplaceAllString(line, "")
		log.Printf("note<%s> parsed title: %+v", we.NoteID, we.Title)
	}

	if byteSanner.Scan() {
		line = byteSanner.Text()
		lsTag := preg_label.FindAllString(line, -1)
		we.Tags = make([]string, 0, len(lsTag))
		for _, tag := range lsTag {
			// 去除两端的 `` 反引号
			tag = tag[1 : len(tag)-1]
			if tag != PUBLIC_LABEL && tag != PRIVATE_LABEL {
				we.Tags = append(we.Tags, tag)
			}
		}
		log.Printf("note<%s> parsed tag: %+v", we.NoteID, we.Tags)
	} else {
		log.Println("fails to parse note:", we.NoteID)
	}

	return nil
}

// 将笔记内容转换为 HTML
func (we *NotePost) HTML() []byte {
	return markdown.ByteToHTML(we.Body)
}
