package configs

type Config struct {
	Env              string           `json:"env"`
	SitesSettings    SitesSettings    `json:"siteSettings"`
	NewRelic         NewRelic         `json:"newRelic"`
	Server           Server           `json:"server"`
	AWS              AWS              `json:"aws"`
	Logger           Logger           `json:"logger"`
	SendInBlue       SendInBlue       `json:"sendInBlue"`
	Zipkin           Zipkin           `json:"zipkin"`
	PdfRenderer      PdfRenderer      `json:"pdfRender"`
	Initiated        bool             `json:"initiated"`
	Zoho             Zoho             `json:"zoho"`
	CloudFlareConfig CloudFlareConfig `json:"cloudinaryConfig,omitempty"`
	FirstKey         string           `json:"-"`
	Adobe            Adobe            `json:"adobe"`
	Margin           int              `json:"margin"`
	Supabaseconfig   SupabaseConfig   `json:"supabaseconfig"`
}
type Zoho struct {
	ZohoCode         string `json:"code"`
	ZohoOrgId        string `json:"orgId"`
	ZohoClientSecret string `json:"zohoSecret"`
	ZohoClientId     string `json:"zohoClientId"`
}
type Adobe struct {
	ProjectName    string `json:"projectName"`
	ClientId       string `json:"clientId"`
	ClientSecret   string `json:"clientSecret"`
	OrganizationId string `json:"organizationId"`
	TechnicalId    string `json:"technicalId"`
	TechnicalEmail string `json:"technicalEmail"`
	PrivateKey     string `json:"privateKey"`
	JwtToken       string `json:"jwtToken"`
}
type Menuloom struct {
	Dev              SiteSetting      `json:"dev"`
	Prod             SiteSetting      `json:"prod"`
	DynamoDb         DynamoDb         `json:"dynamodb"`
	Email            Email            `json:"email"`
	CloudinaryConfig CloudinaryConfig `json:"cloudinaryConfig,omitempty"`
}
type FirstShipper struct {
	Dev              SiteSetting      `json:"dev"`
	Prod             SiteSetting      `json:"prod"`
	DynamoDb         DynamoDb         `json:"dynamodb"`
	EmailConf        EmailConf        `json:"email"`
	RapidShipLTL     RapidShipLTL     `json:"rapidShipLtl"`
	Schneider        RapidShipLTL     `json:"schneider"`
	CloudinaryConfig CloudinaryConfig `json:"cloudinaryConfig,omitempty"`
}
type EmailConf struct {
	SenderName   string
	SenderEmail  string
	RedirectLink string
}
type SiteSetting struct {
	ServiceName     string
	ServiceFullName string
	Version         string
	EnableSignUp    bool
	Domain          string
	BackEndDomain   string
	ApiDomain       string
	FrontEndUrl     string
	EmailId         string
}
type SitesSettings struct {
	Menuloom     Menuloom     `json:"menuloom"`
	FirstShipper FirstShipper `json:"firstshipper"`
}

type DynamoDb struct {
	ProdTable string
	Devtable  string
	PkPrefix  string
	SkPrefix  string
}

type Email struct {
	SmtpHost     string `json:"smptHost"`
	SmtpPort     string `json:"smtpPort"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	FromEmail    string `json:"fromEmail"`
	FromName     string `json:"fromName"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	AccountId    string `json:"accountId"`
	BaseUrl      string `json:"baseUrl"`
}

type NewRelic struct {
	AppName string `json:"app_name,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	License string `json:"license,omitempty"`
}
type Server struct {
	HostName  string
	GrpcPort  string
	HTTPPort  string
	DebugPort string
	Timeout   int
}
type Credentials struct {
	ApiKey    string
	SecretKey string
}
type AwsConfig struct {
	Credentials         Credentials
	SnsSender           string
	CognitoRegion       string
	CognitoUserPoolID   string
	CognitoClientID     string
	CognitoClientSecret string
	BaseAuthorization   string
	CognitoRedirectUri  string
	CognitoUrl          string
	JWKUrl              string
	LoginUrl            string
}
type AWS struct {
	Dev  AwsConfig
	Prod AwsConfig
}
type Logger struct {
	OutputPath        string `json:"outputPath,omitempty"`
	Level             string `json:"level,omitempty"`
	DisableStackTrace bool   `json:"disabledStackTrace,omitempty"`
}

type SendInBlue struct {
	ApiKey string `json:"apikey,omitempty"`
}
type Zipkin struct {
	ReporterURI string  `json:"reporterUri,omitempty"`
	ServiceName string  `json:"serviceName,omitempty"`
	Probability float64 `json:"probability,omitempty"`
}

type PdfRenderer struct {
	ApiKey   string
	Url      string
	FromHtml string
}
type CloudinaryConfig struct {
	CloudName string `json:"cloudName,omitempty"`
	ApiKey    string `json:"apiKey,omitempty"`
	ApiSecret string `json:"apiSecret,omitempty"`
}
type RapidShipLTL struct {
	CompanyName       string `json:"companyName,omitempty"`
	UserName          string `json:"userName,omitempty"`
	Password          string `json:"password,omitempty"`
	AuthTokenName     string `json:"authTokenName,omitempty"`
	AuthUrl           string `json:"authUrl,omitempty"`
	RateUrl           string `json:"rateUrl,omitempty"`
	AddAddressUrl     string `json:"addAddressUrl,omitempty"`
	GetAddressUrl     string `json:"getAddressUrl,omitempty"`
	QuoteHistoryUrl   string `json:"quoteHistoryUrl,omitempty"`
	DispatchUrl       string `json:"dispatchUrl,omitempty"`
	SaveQuoteUrl      string `json:"saveQuoteUrl,omitempty"`
	CancelShipmentUrl string `json:"cancelShipmentUrl,omitempty"`
}
type CloudFlareConfig struct {
	ApiKey      string `json:"apiKey,omitempty"`
	Email       string `json:"email,omitempty"`
	AccountId   string `json:"accountId,omitempty"`
	NamespaceId string `json:"namespaceId,omitempty"`
	ZoneId      string `json:"zoneId,omitempty"`
}
type SupaConfig struct {
	Url         string `json:"url"`
	AnonKey     string `json:"anonkey"`
	ServiceRole string `json:"servicerole"`
	DbUrl       string `json:"dburl"`
	DbPassword  string `json:"dbpassword"`
	JwtKey      string `json:"jwtkey"`
}
type SupabaseConfig struct {
	Dev  SupaConfig `json:"dev"`
	Prod SupaConfig `json:"prod"`
}
