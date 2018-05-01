package template

import (
	"html/template"
	"log"

	"github.com/lymslive/vnotego/notebook"
)

/* tempmgr.go
管理本服务要用到的模板，延迟并缓存编译模板
*/

const (
	TEMP_NOTE_POST string = "notepost"
	TEMP_NOTE_LIST string = "notelist"
)

var loadTemp = make(map[string]*template.Template)

// 获取加载的模板对象，或者先加载
func getTemplate(name string) (*template.Template, error) {
	pTemp, ok := loadTemp[name]
	if ok {
		return pTemp, nil
	}

	tempfile := notebook.TemplateFile(name)
	log.Printf("compile html template file: %s", name)
	// pTemp = template.Must(template.ParseFiles(tempfile))
	pTemp, err := template.ParseFiles(tempfile)
	if pTemp == nil || err != nil {
		log.Printf("fails compile template file: %s", name)
		return nil, err
	}

	loadTemp[name] = pTemp
	return pTemp, nil
}
