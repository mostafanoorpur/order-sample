package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"reflect"
	"strings"
)

var cnf *Config
var Default *Config

func GetConfig() *Config {
	if cnf == nil {
		cnf = Default
	}
	return cnf
}

type Config struct {
	HttpPort         string `mapstructure:"PORT"`
	PostgresHost     string `mapstructure:"POSTGRES_DB_HOST" envDefault:"postgres"`
	PostgresPort     string `mapstructure:"POSTGRES_DB_PORT" envDefault:"5432"`
	PostgresUser     string `mapstructure:"POSTGRES_DB_USER" envDefault:"postgres"`
	PostgresPassword string `mapstructure:"POSTGRES_DB_PASSWORD" envDefault:"postgres"`
	PostgresDbName   string `mapstructure:"POSTGRES_DB_NAME" envDefault:"postgres"`
}

func Init(confPath string) {
	c := Config{}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if confPath == "" {
		err := viper.ReadInConfig()
		if err != nil {
			logrus.Fatal(err, "error reading default config")
		} else {
			logrus.WithFields(logrus.Fields{"config": viper.ConfigFileUsed()}).Info("Viper using default config")
		}
	} else {
		viper.SetConfigFile(confPath)
		err := viper.ReadInConfig()
		if err != nil {
			logrus.Fatal(err, "error reading supplied config")
		} else {
			logrus.WithFields(logrus.Fields{"config": viper.ConfigFileUsed()}).Info("Viper using supplied config")
		}
	}

	bindEnvs(c)
	if err := viper.Unmarshal(&c); err != nil {
		log.Panic(err, "Error Unmarshal Viper Config File")
	}

	Default = &c
}

func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			if err := viper.BindEnv(strings.Join(append(parts, tv), ".")); err != nil {
				log.Fatal(err)
			}
		}
	}
}

// ParseConfig handles environment variable replacement.
func ParseConfig(configValue string) string {
	// Check if the configValue contains a placeholder
	if strings.Contains(configValue, "${") {
		// Extract the variable name and default value, e.g., "${REDIS_PASSWORD:-""}"
		parts := strings.Split(strings.Trim(configValue, "${}"), ":-")
		envVar := parts[0]
		// Return the environment variable value or default
		if value, exists := os.LookupEnv(envVar); exists {
			return value
		}
		// Default value is after ":-", or empty if not provided
		if len(parts) > 1 {
			return parts[1]
		}
	}
	return configValue
}
