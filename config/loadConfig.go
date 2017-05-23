package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Conf 配置文件全局变量
var Conf Config

// Config 配置文件结构
type Config struct {
	Port         string
	FilePath     string
	Database     string
	Databasename string
	Usename      string
	Password     string
}

// Load 加载
func Load() {
	file, _ := ioutil.ReadFile("./config/config.yml")
	config := Config{}
	err := yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	Conf = config
}
