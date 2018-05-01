package template

import (
	"html/template"
	"io"
	"log"

	"github.com/lymslive/vnotego/notebook"
)

/* template/notepost.go
用模板 notepost 生成单篇笔记网页
*/

type tempNotePost struct {
	Title  string
	Tags   []string
	Date   string
	NoteID string
	Body   template.HTML
}

// 向写入流输出页面内容
func GenNotePost(w io.Writer, pNote *notebook.NotePost) error {
	var tempData tempNotePost
	tempData.Title = pNote.Title
	tempData.Tags = pNote.Tags
	tempData.Date = pNote.Date.String()
	tempData.NoteID = pNote.NoteID.String()
	tempData.Body = template.HTML(pNote.HTML())

	tpl, err := getTemplate(TEMP_NOTE_POST)
	if tpl == nil || err != nil {
		return err
	}

	log.Printf("will generate html from markdown note: %s", pNote.NoteID)
	return tpl.Execute(w, tempData)
}

/* 注意：
Execute 与 ExecuteTemplate 区分，多命名。
*/
