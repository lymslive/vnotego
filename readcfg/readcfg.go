package readcfg

import (
	"flag"
	"fmt"
	"github.com/pelletier/go-toml"
	"log"
	"os"
)

type ConfigT struct {
	BookDir string
	CfgFile string
	Server  ServerCfg
}

type ServerCfg struct {
	Host   string
	Port   int
	Static []string
}

// 定一个全局的配置变量
var Config *ConfigT = new(ConfigT)

// 读取 toml 配置文件
func loadConfig(file string) error {
	tree, err := toml.LoadFile(file)
	if err != nil {
		log.Fatalln("fail to load config file:", file, err)
		return err
	}

	err = tree.Unmarshal(Config)
	if err != nil {
		log.Fatalln("fail to unmarshal config tree:", file, err)
		return err
	}

	return nil
}

// 解析命令行及配置文件
func ParseConfig(args []string) *ConfigT {
	var host *string = flag.String("host", "localhost", "server host, ip")
	var port *int = flag.Int("port", 8000, "the default port")
	var dir *string = flag.String("dir", "", "notebook base dir")
	var conf *string = flag.String("conf", "conf.toml", "config file")
	flag.Parse()

	if dir == "" {
		dir = os.Getwd()
	}
	Config.BookDir = dir
	Config.CfgFile = conf

	path := dir + os.PathSeparator + conf
	if _, err = os.Stat(path); os.IsNotExist(err) {
		log.Fatalln("config file not exist: ", path, err)
	}

	loadConfig(path)

	if flag.Lookup("host") != nil {
		Config.Server.Host = host
	}
	if flag.Lookup("port") != nil {
		Config.Server.Port = port
	}

	return ConfigT
}
