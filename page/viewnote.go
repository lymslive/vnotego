package page

import (
	"log"
	"net/http"
	"path"

	"github.com/lymslive/vnotego/notebook"
	"github.com/lymslive/vnotego/template"
)

func responNote(w http.ResponseWriter, r *http.Request) bool {
	file := path.Base(r.URL.Path)
	pNote, err := notebook.FindNote(file, true)
	if pNote == nil || err != nil {
		log.Printf("cannot find note file: %s", file)
		return false
	}

	// n, err := w.Write(pNote.HTML())

	log.Printf("will respond note <%s>", file)
	err = template.GenNotePost(w, pNote)
	if err != nil {
		log.Printf("respond note <%s> error: %s", file, err)
		return false
	}

	return true
}

func responNoteRaw(w http.ResponseWriter, r *http.Request) bool {
	file := path.Base(r.URL.Path)
	pNote, err := notebook.FindNote(file, false)
	if pNote == nil || err != nil {
		log.Printf("cannot find note file: %s", file)
		return false
	}

	n, err := w.Write(pNote.Body)
	log.Printf("responNoteRaw: write byte: %d", n)
	if err != nil {
		log.Printf("responNoteRaw: write err: %s", err)
		return false
	}

	return true
}
