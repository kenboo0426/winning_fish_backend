package config

import (
	"github.com/joho/godotenv"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
}

var Config ConfigList

func init() {
	LoadConfig()
	LoadEnv()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		panic(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("3000"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
	}
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
