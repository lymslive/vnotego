package template

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lymslive/vnotego/notebook"
)

/* template/notelist.go
用模板 notelist 生成笔记列表的网页
*/

const (
	MAX_GLOB_LIST int = 100
)

type tempNoteListor struct {
	Title  string // 标题，去掉 #
	Date   string // yyyy/mm/dd
	NoteID string // yyyymmdd_n
}

type tempNoteList struct {
	SuchList []tempNoteListor
	NumList  int
	MoreList bool
	Label    string
}

func _noteOfdate(NoteID string) string {
	if len(NoteID) < 8 {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", NoteID[0:4], NoteID[4:6], NoteID[6:8])
}

func _genNoteList(w io.Writer, tempData tempNoteList) error {
	tpl, err := getTemplate(TEMP_NOTE_LIST)
	if tpl == nil || err != nil {
		return err
	}
	return tpl.Execute(w, tempData)
}

// 根据指定标签名列出笔记，直接读取 .tag 标签文件的记录
func GenNoteList(w io.Writer, tagName string) error {
	tagFile := notebook.TagFile(tagName)

	fh, err := os.Open(tagFile)
	defer fh.Close()
	if err != nil {
		log.Printf("cannot find tag file: %s", tagFile)
		return err
	}

	var tempData tempNoteList
	tempData.Label = tagName
	tempData.MoreList = false
	tempData.SuchList = make([]tempNoteListor, 0)

	fileScanner := bufio.NewScanner(fh)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		field := strings.Split(line, "\t")
		if len(field) < 2 {
			log.Printf("bad format of tagfile, line string: %s", line)
			continue
		}

		var st tempNoteListor
		st.NoteID, st.Title = field[0], field[1]
		st.Date = _noteOfdate(st.NoteID)
		tempData.SuchList = append(tempData.SuchList, st)
	}

	tempData.NumList = len(tempData.SuchList)
	log.Printf("will generate [%d] note list for tag[%s]", tempData.NumList, tagName)
	return _genNoteList(w, tempData)
}

// 根据部分日期路径来列出该日期下所有笔记，用到文件系统 glob
// datePath 来自 URL ，固定用 / 分隔子路径
func GenNoteGlob(w io.Writer, datePath string) error {
	datePath = strings.TrimSuffix(datePath, "/")
	var spDate []string = strings.Split(datePath, "/")
	var npDate int = len(spDate)

	// 补足至三段日期路径
	switch npDate {
	case 1:
		spDate = append(spDate, "*", "*")
	case 2:
		spDate = append(spDate, "*")
	case 3:
		_ = "ok"
	default:
		return errors.New("bad format of date subpath")
	}

	spDate = append(spDate, "*")
	globPath := notebook.NoteFile(filepath.Join(spDate...))

	log.Printf("will glob note with[%s]", globPath)
	matches, err := filepath.Glob(globPath)
	if err != nil {
		log.Println(err)
		return errors.New("cannot glob note")
	}

	var tempData tempNoteList
	tempData.Label = datePath
	tempData.SuchList = make([]tempNoteListor, 0)

	// 当未限定到日，只到年月时，限制最大 glob
	nMatch := len(matches)
	if nMatch > MAX_GLOB_LIST && npDate < 3 {
		nMatch = MAX_GLOB_LIST
		tempData.MoreList = true
	}

	for i, path := range matches {
		if i >= nMatch {
			break
		}

		var st tempNoteListor
		file := filepath.Base(path)
		dot := strings.LastIndex(file, notebook.NoteExt())
		st.NoteID = file[:dot]
		st.Date = _noteOfdate(st.NoteID)

		head, err := notebook.SeeNoteHead(path, false)
		if err != nil {
			log.Printf("fails to scan head of note[%s]", file)
		}
		st.Title = head.Title

		tempData.SuchList = append(tempData.SuchList, st)
	}

	tempData.NumList = len(tempData.SuchList)
	log.Printf("will generate [%d] note list for date[%s]", tempData.NumList, datePath)
	return _genNoteList(w, tempData)
}
