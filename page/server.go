package page

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lymslive/vnotego/notebook"
	"github.com/lymslive/vnotego/readcfg"
)

// 启动服务器
func StartServe() {
	cfg := readcfg.GetConfig()
	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	setStaticFile()
	setHandler()

	log.Fatal(http.ListenAndServe(address, nil))
}

// 设置各子树、子路径的响应回调函数
const (
	SUB_TREE_TAG  string = "/tag/"
	SUB_TREE_DATE string = "/date/"
	SUB_TREE_RAW  string = "/raw/"

	SUB_PATH_TAG  string = "/tag.xgo"
	SUB_PATH_DATE string = "/date.xgo"
)

func setHandler() {
	http.HandleFunc("/test/", handleTest)
	http.HandleFunc("/home", handleWelcome)

	http.HandleFunc("/", handleRoot)

	http.HandleFunc(SUB_TREE_TAG, handleTreeTag)
	http.HandleFunc(SUB_TREE_DATE, handleTreeDate)
	http.HandleFunc(SUB_TREE_RAW, handleTreeRaw)
	http.HandleFunc(SUB_PATH_TAG, handleQueryTag)
	http.HandleFunc(SUB_PATH_DATE, handleQueryDate)
}

// 静态文件服务
func setStaticFile() {
	cfg := readcfg.GetConfig()

	var dirs = make([]string, 0, 8)
	dirs = append(dirs, notebook.StaticDir()...)
	dirs = append(dirs, cfg.Server.Static...)

	root := cfg.BookDir
	fileServer := http.FileServer(http.Dir(root))

	for _, dir := range dirs {
		pattern := "/" + dir + "/"
		path := cfg.BookDir + notebook.PathSep() + dir
		log.Printf("static add: %s, from root: %s", pattern, path)
		http.Handle(pattern, fileServer)
	}
}

// 根回调函数入口
func handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("req from: %q, res by: %s", r.URL.Path, "handleRoot")

	switch {
	case responNote(w, r):
		log.Printf("success to responNote()")
	default:
		log.Printf("fails to respond to: %q", r.URL.Path)
		http.NotFound(w, r)
	}
}

// 各注册分配回调
func handleWelcome(w http.ResponseWriter, r *http.Request) {
	log.Printf("req from: %q, res by: %s", r.URL.Path, "welcome")

	fmt.Fprintln(w, "Hello, wrold!")
	fmt.Fprintln(w, "Hello, golang!")
	fmt.Fprintln(w, "Hello, vim note!")
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	log.Printf("req from: %q, res by: %s", r.URL.Path, "handler")

	fmt.Fprintf(w, "RUL.Path = %q\n", r.URL.Path)
}

func handleTreeTag(w http.ResponseWriter, r *http.Request) {
	log.Printf("req from: %q, res by: %s", r.URL.Path, "handleTreeTag")
	switch {
	case responListByTag(w, r):
		log.Printf("success to responListByTag()")
	default:
		log.Printf("fails to respond to: %q", r.URL.Path)
		http.NotFound(w, r)
	}
}

func handleTreeDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("req from: %q, res by: %s", r.URL.Path, "handleTreeDate")
	switch {
	case responListByDate(w, r):
		log.Printf("success to responListByDate()")
	default:
		log.Printf("fails to respond to: %q", r.URL.Path)
		http.NotFound(w, r)
	}
}

func handleTreeRaw(w http.ResponseWriter, r *http.Request) {
	log.Printf("req from: %q, res by: %s", r.URL.Path, "handleTreeRaw")
	switch {
	case responNoteRaw(w, r):
		log.Printf("success to responNoteRaw()")
	default:
		log.Printf("fails to respond to: %q", r.URL.Path)
		http.NotFound(w, r)
	}
}

func handleQueryTag(w http.ResponseWriter, r *http.Request) {
	log.Printf("req from: %q, res by: %s", r.URL.Path, "handleQueryTag")
	fmt.Fprintf(w, "RUL.Path = %q\n", r.URL.Path)
}

func handleQueryDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("req from: %q, res by: %s", r.URL.Path, "handleQueryDate")
	fmt.Fprintf(w, "RUL.Path = %q\n", r.URL.Path)
}
