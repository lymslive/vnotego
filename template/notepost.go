package template

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/lymslive/vnotego/notebook"
)

var pTempNote *template.Template = nil

func getTempNote() (*template.Template, error) {
	if pTempNote == nil {
		tempfile := notebook.TemplateFile("notepost")
		// tempfile := notebook.BookDir("") + "/htpl/notepost.html"
		log.Printf("compile html template file: %s", tempfile)
		pTempNote = template.Must(template.ParseFiles(tempfile))
		if pTempNote == nil {
			return nil, errors.New("fails to compile template file: " + tempfile)
		}
	}
	return pTempNote, nil
}

type tempNotePost struct {
	Title string
	Tags  []string
	Date  string
	Body  template.HTML
}

func GenNotePost(w http.ResponseWriter, pNote *notebook.NotePost) error {
	var tempData tempNotePost
	tempData.Title = pNote.Title
	tempData.Tags = pNote.Tags
	tempData.Date = pNote.Date.String()
	tempData.Body = template.HTML(pNote.HTML())

	tpl, err := getTempNote()
	if err != nil {
		return err
	}

	log.Printf("will generate html from markdown note: %s", pNote.NoteID)
	return tpl.Execute(w, tempData)
}

/* 注意：
Execute 与 ExecuteTemplate 区分，多命名。
*/
