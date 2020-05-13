package setting

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Mysql struct {
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	Port        string `yaml:"Port"`
	DBName      string `yaml:"DBName"`
	TablePrefix string `yaml:"TablePrefix"`
}

type Config struct {
	Mysql Mysql `yaml:"Mysql"`
}

var Setting = &Config{}

func init() {
	yamlFile, err := ioutil.ReadFile("conf/app.yaml")
	if err != nil {
		panic(fmt.Sprintf("failed to read yaml file, err: %s", err.Error()))
	}

	err = yaml.Unmarshal(yamlFile, &Setting)
	if err != nil {
		panic(fmt.Sprintf("failed to resolve yaml file, err: %s", err.Error()))
	}
}
