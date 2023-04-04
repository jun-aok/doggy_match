package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	MySqlConnectionString   string
	DistinationSlackChannel string
}

var singletonInstance *Config = newConfig()

func newConfig() *Config {
	env := "development"
	arg_env := os.Getenv("ENV")
	if arg_env != "" {
		env = arg_env
	}
	// main.goから見た相対パス
	f, e1 := os.Open(fmt.Sprintf("./config/%s.json", env))
	if e1 != nil {
		panic(fmt.Sprintf("設定ファイルを読み込めませんでした %s", e1.Error()))
	}
	var cfg Config
	e2 := json.NewDecoder(f).Decode(&cfg)
	if e2 != nil {
		panic(fmt.Sprintf("設定ファイルを読み込めませんでした %s", e2.Error()))
	}
	return &cfg
}

// メモ Configはシングルトンでグローバル
func NewConfig() *Config {
	return singletonInstance
}