package setting

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type App struct {
	Debug         bool   `yaml:"Debug"`
	LogPath       string `yaml:"LogPath"`
	LogTimeFormat string `yaml:"LogTimeFormat"`
}

type Server struct {
	HttpPort     int `yaml:"HttpPort"`
	ReadTimeout  int `yaml:"ReadTimeout"`
	WriteTimeout int `yaml:"WriteTimeout"`
}

type Mysql struct {
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	Port        string `yaml:"Port"`
	DBName      string `yaml:"DBName"`
	TablePrefix string `yaml:"TablePrefix"`
}

type Proxy struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type Config struct {
	App    App    `yaml:"App"`
	Server Server `yaml:"Server"`
	Mysql  Mysql  `yaml:"Mysql"`
	Proxy  Proxy  `yaml:"Proxy"`
}

var Setting = &Config{}

func init() {
	yamlFile, err := ioutil.ReadFile("conf/env.yaml")
	if err != nil {
		panic(fmt.Sprintf("failed to read yaml file, err: %s", err.Error()))
	}

	err = yaml.Unmarshal(yamlFile, &Setting)
	if err != nil {
		panic(fmt.Sprintf("failed to resolve yaml file, err: %s", err.Error()))
	}
}
