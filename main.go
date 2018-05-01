package main

import (
	"fmt"
	"log"

	"github.com/lymslive/vnotego/notebook"
	"github.com/lymslive/vnotego/page"
	"github.com/lymslive/vnotego/readcfg"
)

func main() {

	cfg := readcfg.ParseConfig()
	notebook.SetBaseDir(cfg.BookDir)

	var address = fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	fmt.Println("will Serve on:", address)

	fmt.Printf("default log prefix[%s], flag[%d]\n", log.Prefix(), log.Flags())
	log.SetPrefix("[VNOTEGO] ")
	log.SetFlags(log.Ltime | log.Lshortfile)
	fmt.Printf("now log prefix[%s], flag[%d]\n", log.Prefix(), log.Flags())

	page.StartServe()

	log.Printf("Serve on %s success! But cannot reach here\n", address)
}
