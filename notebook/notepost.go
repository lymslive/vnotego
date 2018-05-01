package notebook

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

// 笔记基础头信息，前两行的标题与标签
type NoteHead struct {
	Title string
	Tags  []string
}

// 笔记类型
// Body 其实包含文件所有内容，包括从中解析的标题行与标签
type NotePost struct {
	NoteID
	NoteHead
	Body []byte
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

// 将笔记内容转换为 HTML
func (we *NotePost) HTML() []byte {
	return markdown.ByteToHTML(we.Body)
}

// 根据文件名，对应笔记名，寻找笔记，返回笔记对象指针
// 如果不能读取笔记文件，返回 nil
// file 参数不需要后缀
// seeHead 参数表示是否要解析标题与标签行
func FindNote(file string, seeHead bool) (*NotePost, error) {
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

	if seeHead {
		np.parseNote()
	}

	return np, err
}

// 解析标题与标签
func (we *NotePost) parseNote() error {
	byteReader := bytes.NewReader(we.Body)
	byteScanner := bufio.NewScanner(byteReader)
	head, err := _parseNote(byteScanner, true)
	we.NoteHead = head

	log.Printf("note<%s> parsed title: %+v", we.NoteID, we.Title)
	log.Printf("note<%s> parsed tag: %+v", we.NoteID, we.Tags)
	if err != nil {
		log.Println("fails to parse note[%s]: %s", we.NoteID, err)
	}

	return err
}

// 获取笔记文件基本信息，标题与标签列表
func SeeNoteHead(path string, includeTag bool) (head NoteHead, err error) {
	fh, err := os.Open(path)
	defer fh.Close()
	if err != nil {
		log.Printf("cannot open file[%s]: %s", path, err)
		return
	}

	fileScanner := bufio.NewScanner(fh)
	head, err = _parseNote(fileScanner, includeTag)
	return
}

func _parseNote(scanner *bufio.Scanner, scanTag bool) (head NoteHead, err error) {
	if scanner.Scan() {
		line := scanner.Text()
		head.Title = preg_title.ReplaceAllString(line, "")
	} else {
		err = scanner.Err()
	}

	if !scanTag {
		return
	}

	if scanner.Scan() {
		line := scanner.Text()
		lsTag := preg_label.FindAllString(line, -1)
		head.Tags = make([]string, 0, len(lsTag))
		for _, tag := range lsTag {
			// 去除两端的 `` 反引号
			tag = tag[1 : len(tag)-1]
			if tag != PUBLIC_LABEL && tag != PRIVATE_LABEL {
				head.Tags = append(head.Tags, tag)
			}
		}
	} else {
		err = scanner.Err()
	}

	return
}
