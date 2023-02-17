package configs

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/viper"
)

var configs *Config

func init() {
	fmt.Println("configs init")
	configs = &Config{}
	viper := viper.New()
	currentDir, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("could not read current directory: %w", err)
		log.Fatal(err)
	}
	currentDir = strings.Split(currentDir, "backend")[0]
	path := filepath.Dir(currentDir)
	path = filepath.Join(path, "backend", "configs")
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("count not read config file %v", err))
	}
	err = viper.Unmarshal(&configs)
	if err != nil {
		panic(fmt.Sprintf("count not read config file %v", err))
	}
	configs.Initiated = true
}

func (conf *Config) Retrieve(ctx context.Context) (aws.Credentials, error) {
	if conf.Env == "dev" {
		return aws.Credentials{
			AccessKeyID:     conf.AWS.Prod.Credentials.ApiKey,
			SecretAccessKey: conf.AWS.Prod.Credentials.SecretKey,
		}, nil
	}
	return aws.Credentials{
		AccessKeyID:     conf.AWS.Prod.Credentials.ApiKey,
		SecretAccessKey: conf.AWS.Prod.Credentials.SecretKey,
	}, nil

}

func GetConfig() *Config {
	if configs != nil || !configs.Initiated {
		return configs
	}
	return configs
}

func (config *Config) GetMenuloomFontEndURL() string {
	if config.Env == "dev" {
		return config.SitesSettings.Menuloom.Dev.FrontEndUrl
	}
	return config.SitesSettings.Menuloom.Prod.FrontEndUrl
}
func (config *Config) GetFirstShipperFontEndURL() string {
	if config.Env == "dev" {
		return config.SitesSettings.FirstShipper.Dev.FrontEndUrl
	}
	return config.SitesSettings.FirstShipper.Prod.FrontEndUrl
}

func (config *Config) GetMenuloomBackEndURL() string {
	if config.Env == "dev" {
		return config.SitesSettings.Menuloom.Dev.BackEndDomain
	}
	return config.SitesSettings.Menuloom.Prod.BackEndDomain
}
func (config *Config) GetFirstShipperBackEndURL() string {
	if config.Env == "dev" {
		return config.SitesSettings.FirstShipper.Dev.BackEndDomain
	}
	return config.SitesSettings.FirstShipper.Prod.BackEndDomain
}
func (config *Config) GetFirstShipperTableName() string {
	if config.Env == "dev" {
		return config.SitesSettings.FirstShipper.DynamoDb.Devtable
	}
	return config.SitesSettings.FirstShipper.DynamoDb.ProdTable
}

func (config *Config) GetMenuloomTableName() string {
	if config.Env == "dev" {
		return config.SitesSettings.Menuloom.DynamoDb.Devtable
	}
	return config.SitesSettings.Menuloom.DynamoDb.ProdTable
}
func (config *Config) GetFirstShipperPKPrefix() string {
	return config.SitesSettings.FirstShipper.DynamoDb.PkPrefix
}

func (config *Config) GetMenuloomPKPrefix() string {
	return config.SitesSettings.Menuloom.DynamoDb.PkPrefix
}
func (config *Config) GetFirstShipperSKPrefix() string {
	return config.SitesSettings.FirstShipper.DynamoDb.SkPrefix
}
func (config *Config) GetMenuloomSKPrefix() string {
	return config.SitesSettings.Menuloom.DynamoDb.SkPrefix
}
func (config *Config) GetMenuloomServiceName() string {
	return "menuloom"
}
func (config *Config) GetFirstShipperServiceName() string {
	return "firstshipper"
}
func (config *Config) GetAwsConfig() AwsConfig {
	if config.Env == "dev" {
		return config.AWS.Dev
	}
	return config.AWS.Prod
}
