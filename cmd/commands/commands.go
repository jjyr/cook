package commands

import (
	"github.com/urfave/cli"
	"github.com/jjyr/cook/config"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"github.com/jjyr/cook/controller"
	"gopkg.in/yaml.v2"
	"github.com/jjyr/cook/common"
	"fmt"
)

func getConfigPath(context *cli.Context) (configPath string) {
	configPath = context.String("config")
	if configPath != "" {
		return
	}
	workDir, err := os.Getwd()
	if err != nil {
		logrus.Fatal(err)
	}
	configPath = path.Join(workDir, "./cook.yml")
	return
}

func initConfig(context *cli.Context) (c config.Config) {
	configPath := getConfigPath(context)
	c, err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Fatalf("error: %s\nLoad config file failed: %s, type 'cook help' to see usage\n", err, configPath)
	}
	return
}

func Main(c *cli.Context) (err error) {
	ctl := controller.NewController(initConfig(c))
	if err = ctl.Build(); err != nil {
		logrus.Fatal(err)
	}
	if err = ctl.Prepare(); err != nil {
		logrus.Fatal(err)
	}
	if err = ctl.Deploy(); err != nil {
		logrus.Fatal(err)
	}
	return
}

func Build(c *cli.Context) (err error) {
	ctl := controller.NewController(initConfig(c))
	if err = ctl.Build(); err != nil {
		logrus.Fatal(err)
	}
	return
}

func Prepare(c *cli.Context) (err error) {
	ctl := controller.NewController(initConfig(c))
	if err = ctl.Prepare(); err != nil {
		logrus.Fatal(err)
	}
	return
}

func Deploy(c *cli.Context) (err error) {
	ctl := controller.NewController(initConfig(c))
	if err = ctl.Deploy(); err != nil {
		logrus.Fatal(err)
	}
	return
}

func Config(c *cli.Context) (err error) {
	var configure config.Config
	if c.Bool("sample") {
		configure = config.Config{
			Target: []common.Server{{Host: "my_web_server", User: "web"}},
			Deploy: []common.DeployDesc{{Path: "docker-compose.yml"}},
		}
	} else {
		configure = initConfig(c)
	}
	out, err := yaml.Marshal(configure)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	return
}
