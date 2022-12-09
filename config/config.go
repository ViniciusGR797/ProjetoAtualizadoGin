package config

import (
	// Package "os" - usado para manipulação de arquivos
	// Package "strconv" - implementa conversões de tipos primitivos, ex: String para int
	"os"
	"strconv"
)

const (
	// DEVELOPER - Modo de desenvolvimento - Ambiente programável
	// HOMOLOGATION - Modo de homologação - Fase de testes
	// PRODUCTION - Modo de produção - Usuário final
	DEVELOPER    = "developer"
	HOMOLOGATION = "homologation"
	PRODUCTION   = "production"
)

// Config - Configurações do servidor
type Config struct {
	SRV_PORT    string `json:"srv_port"`
	WEB_UI      bool   `json:"web_ui"`
	Mode        string `json:"mode"`
	OpenBrowser bool   `json:"open_browser"`
	DBConfig    `json:"dbconfig"`
}

// DBConfig - Configurações do banco de dados
type DBConfig struct {
	DB_DRIVE string `json:"db_drive"`
	DB_HOST  string `json:"db_host"`
	DB_PORT  string `json:"db_port"`
	DB_USER  string `json:"db_user"`
	DB_PASS  string `json:"db_pass"`
	DB_NAME  string `json:"db_name"`
	DB_DSN   string `json:"-"`
}

// NewConfig - Cria uma nova configuração
func NewConfig(config *Config) *Config {

	var conf *Config

	if config == nil || config.SRV_PORT == "" {
		conf = DefaultConfig()
	} else {
		conf = config
	}

	SRV_PORT := os.Getenv("SRV_PORT")
	if SRV_PORT != "" {
		conf.SRV_PORT = SRV_PORT
	}

	SRV_MODE := os.Getenv("SRV_MODE")
	if SRV_MODE != "" {
		conf.Mode = SRV_MODE
	}

	SRV_WEB_UI := os.Getenv("SRV_WEB_UI")
	if SRV_WEB_UI != "" {
		conf.WEB_UI, _ = strconv.ParseBool(SRV_WEB_UI)
	}

	SRV_DB_DRIVE := os.Getenv("SRV_DB_DRIVE")
	if SRV_DB_DRIVE != "" {
		conf.DBConfig.DB_DRIVE = SRV_DB_DRIVE
	}

	SRV_DB_HOST := os.Getenv("SRV_DB_HOST")
	if SRV_DB_HOST != "" {
		conf.DBConfig.DB_HOST = SRV_DB_HOST
	}

	SRV_DB_PORT := os.Getenv("SRV_DB_PORT")
	if SRV_DB_PORT != "" {
		conf.DBConfig.DB_PORT = SRV_DB_PORT
	}

	SRV_DB_USER := os.Getenv("SRV_DB_USER")
	if SRV_DB_USER != "" {
		conf.DBConfig.DB_USER = SRV_DB_USER
	}

	SRV_DB_PASS := os.Getenv("SRV_DB_PASS")
	if SRV_DB_PASS != "" {
		conf.DBConfig.DB_PASS = SRV_DB_PASS
	}

	SRV_DB_NAME := os.Getenv("SRV_DB_NAME")
	if SRV_DB_NAME != "" {
		conf.DBConfig.DB_NAME = SRV_DB_NAME
	}

	return config
}

// defaultConf - Configurações padrão
func DefaultConfig() *Config {
	default_config := Config{
		SRV_PORT:    "9000",
		WEB_UI:      true,
		OpenBrowser: true,
		DBConfig: DBConfig{
			DB_DRIVE: "mysql",
			DB_HOST:  "localhost",
			DB_PORT:  "3306",
			DB_USER:  "root",
			DB_PASS:  "",
			DB_NAME:  "group_d",
		},
		Mode: PRODUCTION,
	}

	return &default_config
}
