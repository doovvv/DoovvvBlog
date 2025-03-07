package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Appmode  string `yaml:"appmode"`
	Secret   string `yaml:"secret"`
	FilePath string `yaml:"file_path"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

var AppConfig Config

func Init() {
	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		fmt.Println("读取配置文件失败")

	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		fmt.Println("解析配置文件失败")
	}
}
