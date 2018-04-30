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
	notebook.BookDir(cfg.BookDir)

	var address = fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Println("will Serve on:", address)

	page.StartServe()

	log.Printf("Serve on %s success! But cannot reach here\n", address)
}
