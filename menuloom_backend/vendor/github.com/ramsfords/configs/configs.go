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
	"github.com/aws/aws-sdk-go-v2/config"
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
	currentDir = strings.Split(currentDir, "auther")[0]
	fmt.Print(currentDir)
	path := filepath.Dir(currentDir)
	path = filepath.Join(path, "auther", "configs")
	viper.AddConfigPath(path)
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	fmt.Println(err)
	env := viper.Get("env").(string)
	if env == "dev" {
		viper.SetConfigName("dev")
		viper.SetConfigType("yaml")
		err = viper.ReadInConfig()
		fmt.Println(err)
		err = viper.Unmarshal(&configs)
		fmt.Println(err)
		configs.Env = env
	} else if env == "prod" {
		viper.SetConfigName("prod")
		viper.SetConfigType("yaml")
		err = viper.ReadInConfig()
		fmt.Println(err)
		err = viper.Unmarshal(&configs)
		fmt.Println(err)
		configs.Env = env
		currenDirectory, err := os.Getwd()
		if err != nil {
			log.Fatal("could not get current directory")
		}
		projectDir := strings.Split(currenDirectory, "/cmd")[0]
		configs.ProjectDir = projectDir
		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithCredentialsProvider(configs), config.WithRegion("us-west-1"))
		if err != nil {
			err = fmt.Errorf("could not load the AWS Config")
			log.Fatal(err)
		}
		configs.Initiated = true
		configs.LogPath = projectDir + "/logs/logs.txt"
		configs.AwsConfig = cfg
		configs.ZohoConfig.FrontEndUrl = configs.GetFontEndURL()
		configs.ZohoConfig.BackEndUrl = configs.GetBackEndURL()
	}
}

type NewRelic struct {
	AppName string `json:"app_name,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	License string `json:"license,omitempty"`
}
type Server struct {
	Env       string `json:"env,omitempty"`
	Host      string `json:"host,omitempty"`
	GrpcPort  string `json:"grpc_port,omitempty"`
	HttpPort  string `json:"http_port,omitempty"`
	DebugPort string `json:"debug_port,omitempty"`
	TimeOut   int    `json:"time_out,omitempty"`
}
type Cognito struct {
	KeySetUrl  string `json:"key_set_url,omitempty"`
	AppId      string `json:"app_id,omitempty"`
	Region     string `json:"region,omitempty"`
	UserPoolId string `json:"user_pool_id,omitempty"`
	PoolArn    string `json:"pool_arn,omitempty"`
	ClientId   string `json:"client_id,omitempty"`
}
type Credentials struct {
	ApiKey    string
	SecretKey string
}
type DynamoDb struct {
	Menuloom     DbSetup `json:"menuloom,omitempty"`
	FirstShipper DbSetup `json:"firstshipper,omitempty"`
}
type DbSetup struct {
	ProdTable string
	Devtable  string
	PkPrefix  string
	SkPrefix  string
}
type Redis struct {
	Password string
	Host     string
	Db       string
}
type AWS struct {
	Credentials
}
type Logger struct {
	OutputPath        string `json:"output_path,omitempty"`
	Level             string `json:"level,omitempty"`
	DisableStackTrace bool   `json:"disabled_stack_trace,omitempty"`
}
type Database struct {
	Port       string `json:"port,omitempty"`
	Host       string `json:"host,omitempty"`
	UserName   string `json:"user_name,omitempty"`
	Password   string `json:"password,omitempty"`
	DbName     string `json:"db_name,omitempty"`
	DisableSSL string `json:"disable_ssl,omitempty"`
}
type GRPC struct {
	MaxRecvMsgSizeInMB int `json:"max_recv_msg_size_in_mb,omitempty"`
	MaxSendMsgSizeInMB int `json:"max_send_msg_size_in_mb,omitempty"`
}
type SendInBlue struct {
	API string `json:"api,omitempty"`
}
type KeysFolder struct {
	Folder    string `json:"folder,omitempty"`
	ActiveKid string `json:"active_kid,omitempty"`
}
type Zipkin struct {
	ReporterURI string  `json:"reporter_uri,omitempty"`
	ServiceName string  `json:"service_name,omitempty"`
	Probability float64 `json:"probability,omitempty"`
}
type RapidShipLTL struct {
	CompanyName       string `json:"company_name,omitempty"`
	UserName          string `json:"username,omitempty"`
	Email             string `json:"email,omitempty"`
	Password          string `json:"password,omitempty"`
	AuthUrl           string `json:"auth_url,omitempty"`
	RateUrl           string `json:"rate_url,omitempty"`
	AddAddressUrl     string `json:"add_address_url,omitempty"`
	AuthTokenName     string `json:"auth_token_name,omitempty"`
	AddressesUrl      string `json:"addresses_url,omitempty"`
	QuoteHistoryUrl   string `json:"quote_history_url,omitempty"`
	DispatchUrl       string `json:"dispatch_url,omitempty"`
	CancelShipmentUrl string `json:"cancel_shipment_url,omitempty"`
	SaveQuoteUrl      string `json:"save_quote_url,omitempty"`
}
type Schneider struct {
	CompanyName     string `json:"company_name,omitempty"`
	UserName        string `json:"username,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	AuthUrl         string `json:"auth_url,omitempty"`
	RateUrl         string `json:"rate_url,omitempty"`
	AddAddressUrl   string `json:"add_address_url,omitempty"`
	AuthTokenName   string `json:"auth_token_name,omitempty"`
	AddressesUrl    string `json:"addresses_url,omitempty"`
	QuoteHistoryUrl string `json:"quote_history_url,omitempty"`
}
type CloudinaryConfig struct {
	CloudName string `json:"cloud_name,omitempty"`
	ApiKey    string `json:"api_key,omitempty"`
	ApiSecret string `json:"api_secret,omitempty"`
}

type Scylladb struct {
	Name     string
	Public   []string
	Username string
	Password string
	Region   string
	Keyspace string
}

type Config struct {
	Domain               string `json:"domain,omitempty"`
	Env                  string `json:"env,omitempty"`
	DevFrontEndURL       string `json:"dev_frontend_url"`
	ProdFrontEndURL      string `json:"prod_frontend_url"`
	DevBackEndURL        string `json:"dev_backend_url"`
	ProdBackendURL       string `json:"prod_backend_url"`
	Testport             string `json:"test_port,omitempty"`
	ServiceName          string `json:"service_name,omitempty"`
	ServiceFullName      string `json:"service_full_name,omitempty"`
	Version              string `json:"version,omitempty"`
	LogPath              string `json:"log_path,omitempty"`
	UserTestLog          string `json:"user_test_log,omitempty"`
	EnableSignUp         bool   `json:"enable_sign_up,omitempty"`
	SessionKey           string `json:"session_key,omitempty"`
	LambdaGateway        string `json:"lambda_gateway,omitempty"`
	AwsDev               bool   `json:"aws_dev,omitempty"`
	AwsHost              string `json:"aws_host,omitempty"`
	Docker               bool   `json:"docker,omitempty"`
	ForgotPasswordExpiry int64  `json:"forgot_password_expiry,omitempty"`
	FrontEndUrl          string `json:"fontend_url"`
	ConfirmEmailLink     string `json:"confirm_email_link"`
	ResetPasswordLink    string `json:"reset_password_link"`
	WelcomeEmailLink     string `json:"welcome_email_link"`
	ProjectDir           string `json:"project_dir"`
	AwsConfig            aws.Config
	Initiated            bool
	DynamoDb             `json:"dynamodb,omitempty"`
	Cognito              `json:"cognito,omitempty"`
	NewRelic             `json:"new_relic,omitempty"`
	Server               `json:"server,omitempty"`
	AWS                  `json:"aws,omitempty"`
	Logger               `json:"logger,omitempty"`
	Redis                `json:"redis,omitempty"`
	Database             `json:"database,omitempty"`
	GRPC                 `json:"grpc,omitempty"`
	SendInBlue           `json:"send_in_blue,omitempty"`
	KeysFolder           `json:"keys_folder,omitempty"`
	Zipkin               `json:"zipkin,omitempty"`
	RapidShipLTL         `json:"rapid_ship_ltl,omitempty"`
	Schneider            `json:"schneider"`
	LastPickupTime       int `dynamodbav:"last_pickup_time,omitempty" json:"last_pickup_time,omitempty"`
	QuoteCountIncreaseBy int
	PdfRender            `json:"pdf_render,omitempty"`
	ZohoConfig           `json:"zoho_config,omitempty"`
	CloudinaryConfig     `json:"cloudinary_config,omitempty"`
	ScyllaDb             `json:"scylla_db,omitempty"`
}
type ZohoConfig struct {
	BaseUrl      string
	ClientId     string
	ClientSecret string
	AccountId    string
	SmtpHost     string
	SmtpPort     string
	UserName     string
	Password     string
	FromAddress  string
	SenderName   string
}

type PdfRender struct {
	ApiKey   string
	Url      string
	FromHtml string
}
type Scylladb struct {
	Name     string
	Public   []string
	Username string
	Password string
	Region   string
	Keyspace string
}

func (conf *Config) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     conf.AWS.Credentials.ApiKey,
		SecretAccessKey: conf.AWS.Credentials.SecretKey,
	}, nil
}

func GetConfig() *Config {
	if configs != nil || !configs.Initiated {
		return configs
	}
	return configs
}

func (config *Config) GetFontEndURL() string {
	if config.Env == "dev" {
		url := config.DevFrontEndURL
		return url
	}
	url := config.ProdFrontEndURL
	return url
}

func (config *Config) GetBackEndURL() string {
	if config.Env == "dev" {
		url := config.DevBackEndURL
		return url
	}
	url := config.ProdBackendURL
	return url
}
func (config *Config) GetFirstShipperTableName() string {
	if config.Env == "dev" {
		return config.FirstShipper.Devtable
	}
	return config.FirstShipper.ProdTable
}

func (config *Config) GetMenuloomTableName() string {
	if configs.Env == "dev" {
		return config.Menuloom.Devtable
	}
	return config.Menuloom.ProdTable
}
func (config *DynamoDb) GetFirstShipperPKPrefix() string {
	return config.FirstShipper.PkPrefix
}

func (config *DynamoDb) GetMenuloomPKPrefix() string {
	return config.Menuloom.PkPrefix
}
func (config *DynamoDb) GetFirstShipperSKPrefix() string {
	return config.FirstShipper.SkPrefix
}

func (config *DynamoDb) GetMenuloomSKPrefix() string {
	return config.Menuloom.SkPrefix
}
