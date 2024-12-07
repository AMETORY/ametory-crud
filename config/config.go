package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// ServerConfiguration holds the server-related config values
type ServerConfiguration struct {
	AppName    string `mapstructure:"app_name"`
	AppDesc    string `mapstructure:"app_desc"`
	Version    string `mapstructure:"version"`
	ApiURL     string `mapstructure:"api_url"`
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	SecretKey  string `mapstructure:"secret_key"`
	ExpiredJWT int    `mapstructure:"expired_jwt"`
}

// DatabaseConfiguration holds the database-related config values
type DatabaseConfiguration struct {
	Type      string `mapstructure:"type"`
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Name      string `mapstructure:"name"`
	AuthTable string `mapstructure:"auth_table"`
}

// SchedulerConfiguration holds the scheduler-related config values
type SchedulerConfiguration struct {
	Enabled  bool `mapstructure:"enabled"`
	Interval int  `mapstructure:"interval"`
}

// MailerConfiguration holds the mailer-related config values
type MailerConfiguration struct {
	SMTPHost string `mapstructure:"smtp_host"`
	SMTPPort int    `mapstructure:"smtp_port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// S3Configuration holds the S3-related config values
type S3Configuration struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
}

// ElasticSearchConfiguration holds the ES-related config values
type ElasticSearchConfiguration struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// RedisConfiguration holds the Redis-related config values
type RedisConfiguration struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// TwilioSMSGateway holds the Twilio SMS config values
type TwilioSMSGateway struct {
	AccountSID  string `mapstructure:"account_sid"`
	AuthToken   string `mapstructure:"auth_token"`
	PhoneNumber string `mapstructure:"phone_number"`
}

// Configuration holds all configurations from YAML
type Configuration struct {
	Server    ServerConfiguration        `mapstructure:"server"`
	Database  DatabaseConfiguration      `mapstructure:"database"`
	Scheduler SchedulerConfiguration     `mapstructure:"scheduler"`
	Mailer    MailerConfiguration        `mapstructure:"mailer"`
	S3        S3Configuration            `mapstructure:"s3"`
	ES        ElasticSearchConfiguration `mapstructure:"es"`
	Redis     RedisConfiguration         `mapstructure:"redis"`
	SMS       TwilioSMSGateway           `mapstructure:"sms"`
}

var App = &Configuration{}

// InitConfig initializes the configuration by reading from the YAML file
func InitConfig() (*Configuration, error) {

	// Set the config file path and type (yaml)
	viper.SetConfigName(".env") // config.yaml
	viper.AddConfigPath(".")    // Look for the config file in the current directory
	viper.SetConfigType("yaml")

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config file, %s", err)
	}

	// Initialize Configuration struct
	var config Configuration

	// Unmarshal the config into the struct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("Error unmarshalling config into struct: %s", err)
	}
	App = &config
	return &config, nil
}
