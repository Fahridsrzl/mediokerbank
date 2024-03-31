package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type LogFileConfig struct {
	FilePath string
}

type TokenConfig struct {
	IssuerName           string
	JwtSignatureKey      []byte
	AccessTokenLifeTime  time.Duration
	RefreshTokenLifeTime time.Duration
}

type MailerConfig struct {
	MailerHost     string
	MailerPort     int
	MailerUsername string
	MailerPassword string
}

type Config struct {
	ApiConfig
	DbConfig
	LogFileConfig
	TokenConfig
	MailerConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.LogFileConfig = LogFileConfig{
		FilePath: os.Getenv("LOG_FILE"),
	}

	accessTokenLifeTime, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFE_TIME"))
	if err != nil {
		return err
	}

	refreshTokenLifeTime, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFE_TIME"))
	if err != nil {
		return err
	}

	c.TokenConfig = TokenConfig{
		IssuerName:           os.Getenv("TOKEN_ISSUE_NAME"),
		JwtSignatureKey:      []byte(os.Getenv("TOKEN_KEY")),
		AccessTokenLifeTime:  time.Duration(accessTokenLifeTime) * time.Hour,
		RefreshTokenLifeTime: time.Duration(refreshTokenLifeTime) * time.Hour,
	}

	mailerPort, err := strconv.Atoi(os.Getenv("MAILER_PORT"))
	if err != nil {
		return err
	}

	c.MailerConfig = MailerConfig{
		MailerHost:     os.Getenv("MAILER_HOST"),
		MailerPort:     mailerPort,
		MailerUsername: os.Getenv("MAILER_USERNAME"),
		MailerPassword: os.Getenv("MAILER_PASSWORD"),
	}

	if c.ApiPort == "" || c.Host == "" || c.Port == "" || c.Name == "" || c.User == "" || c.Password == "" || c.FilePath == "" || c.IssuerName == "" ||
		c.JwtSignatureKey == nil || c.AccessTokenLifeTime == 0 || c.RefreshTokenLifeTime == 0 ||
		c.MailerHost == "" || c.MailerPort == 0 || c.MailerUsername == "" || c.MailerPassword == "" {
		return errors.New("environment required")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}
