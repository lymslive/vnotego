package page

import (
	"log"
	"net/http"
	"strings"

	"github.com/lymslive/vnotego/template"
)

func responListByTag(w http.ResponseWriter, r *http.Request) bool {
	tagName := strings.TrimPrefix(r.URL.Path, SUB_TREE_TAG)
	if tagName == "" {
		return false
	}

	log.Printf("will respond by tag[%s]", tagName)
	err := template.GenNoteList(w, strings.ToLower(tagName))
	if err != nil {
		log.Printf("respond tag[%s] error: %s", tagName, err)
		return false
	}

	return true
}

func responListByDate(w http.ResponseWriter, r *http.Request) bool {
	datePath := strings.TrimPrefix(r.URL.Path, SUB_TREE_DATE)
	if datePath == "" {
		return false
	}

	log.Printf("will respond by date[%s]", datePath)
	err := template.GenNoteGlob(w, datePath)
	if err != nil {
		log.Printf("respond date[%s] error: %s", datePath, err)
		return false
	}

	return true
}

func responTagQuery(w http.ResponseWriter, r *http.Request) bool {
	return true
}
