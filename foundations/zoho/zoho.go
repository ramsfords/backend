package zoho

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/ramsfords/backend/configs"
)

//go:embed tokens.txt
var tokenstr string

type tokens struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ApiDomain    string    `json:"api_domain"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
	Error        string    `json:"error"`
}

// New initializes a Zoho structure
func New(conf *configs.Config) (*Zoho, error) {
	retryClient := retryablehttp.NewClient()
	retryClient.Logger = nil
	retryClient.RetryMax = 1
	retryClient.HTTPClient = &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
	token := &tokens{}
	err := json.Unmarshal([]byte(tokenstr), token)
	if err != nil {
		return nil, err
	}
	z := Zoho{
		Client:     retryClient.StandardClient(),
		ZohoTLD:    "com",
		TokensFile: "./tokens.txt",
		Oauth: OAuth{
			BaseURL:      "https://accounts.zoho.com/oauth/v2/",
			ClientID:     conf.Zoho.ZohoClientId,
			ClientSecret: conf.Zoho.ZohoClientSecret,
			RedirectURI:  "https://www.ramsfords.com",
			Scopes: []string{
				"ZohoBooks.fullaccess.all",
				"ZohoCRM.modules.ALL",
			},
			Token: *token,
		},
	}
	z.RefreshTokenRequest()
	return &z, nil

}

// SetZohoTLD can be used to set the TLD extension for API calls for example for Zoho in EU and China.
// by default this is set to "com", other options are "eu" and "ch"
func (z *Zoho) SetZohoTLD(s string) {
	z.ZohoTLD = s
	z.Oauth.BaseURL = fmt.Sprintf("https://accounts.zoho.%s/oauth/v2/", s)
}

// CustomHTTPClient can be used to provide a custom HTTP Client that replaces the once instantiated
// when executing New()
//
// A notable use case is AppEngine where a user must use the appengine/urlfetch packages provided http client
// when performing outbound http requests.
func (z *Zoho) CustomHTTPClient(c *http.Client) {
	z.Client = c
}

// SetOrganizationID can be used to add organization id in zoho struct
// which is needed for expense apis
func (z *Zoho) SetOrganizationID(orgID string) {
	z.OrganizationID = orgID
}

// Zoho is for accessing all APIs. It is used by subpackages to simplify passing authentication
// values between API subpackages.
type Zoho struct {
	Oauth          OAuth
	Client         *http.Client
	TokensFile     string
	OrganizationID string
	ZohoTLD        string
}

// OAuth is the OAuth part of the Zoho struct
type OAuth struct {
	Scopes       []string
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Token        tokens
	BaseURL      string
}
