package readcfg

import (
	"fmt"
	"github.com/lymslive/gotoo/assert"
	"testing"
)

func TestToml(t *testing.T) {
	loadConfig("conf.toml")
	fmt.Printf("%T : %+v\n", config, config)
	fmt.Printf("%T : %#v\n", *config, *config)

	assert.BeginTest(t)
	defer assert.EndTest()

	assert.Equal(config.Server.Host, "localhost", "toml parse error")
	assert.Equal(config.Server.Port, 8000, "toml parse error")
	assert.Equal(len(config.Server.Static), 3, "toml parse error")
	assert.Equal(config.Server.Static[0], "css", "toml parse error")
	assert.Equal(config.Server.Static[1], "js", "toml parse error")
	assert.Equal(config.Server.Static[2], "img", "toml parse error")
}

func TestMain(m *testing.M) {
	ParseConfig()
	fmt.Printf("%T : %#v\n", *config, *config)
}
