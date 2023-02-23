module github.com/ramsfords/backend

go 1.18

require (
	github.com/SebastiaanKlippert/go-wkhtmltopdf v1.8.2
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/aws/aws-sdk-go v1.44.206
	github.com/aws/aws-sdk-go-v2 v1.17.5
	github.com/aws/aws-sdk-go-v2/config v1.18.12
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.10.11
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.49
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.22.2
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.18.2
	github.com/aws/aws-sdk-go-v2/service/lambda v1.28.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.30.2
	github.com/aws/aws-sdk-go-v2/service/ses v1.15.1
	github.com/cloudflare/cloudflare-go v0.58.1
	github.com/cloudinary/cloudinary-go/v2 v2.2.0
	github.com/go-mail/mail v2.3.1+incompatible
	github.com/go-resty/resty/v2 v2.7.0
	github.com/goccy/go-json v0.10.0
	github.com/gofor-little/env v1.0.9
	github.com/gofor-little/xerror v1.0.4
	github.com/golang-jwt/jwt/v4 v4.4.3
	github.com/hashicorp/go-retryablehttp v0.7.2
	github.com/jhillyerd/enmime v0.10.1
	github.com/labstack/echo/v5 v5.0.0-20220201181537-ed2888cfa198
	github.com/labstack/gommon v0.4.0
	github.com/lestrrat-go/jwx v1.2.25
	github.com/newrelic/go-agent/v3 v3.20.3
	github.com/pkg/errors v0.9.1
	github.com/ramsfords/types_gen v0.0.0-20230219233847-c32f89b6a854
	github.com/schmorrison/go-querystring v1.1.1
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/viper v1.14.0
	github.com/valyala/fasthttp v1.44.0
	go.opentelemetry.io/otel v1.13.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.13.0
	go.opentelemetry.io/otel/sdk v1.13.0
	golang.org/x/crypto v0.5.0
	golang.org/x/time v0.3.0
	google.golang.org/appengine v1.6.7
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.13.12 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.22 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.28 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.22 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.14.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.22 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.22 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.22 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.18.3 // indirect
	github.com/aws/smithy-go v1.13.5 // indirect
	github.com/cention-sany/utf7 v0.0.0-20170124080048-26cad61bd60a // indirect
	github.com/creasty/defaults v1.5.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.0-20210816181553-5444fa50b93d // indirect
	github.com/fatih/color v1.14.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogs/chardet v0.0.0-20191104214054-4b6791f73a28 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/schema v1.2.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jaytaylor/html2text v0.0.0-20200412013138-3577fbdbcff7 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.15.15 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/lestrrat-go/backoff/v2 v2.0.8 // indirect
	github.com/lestrrat-go/blackmagic v1.0.0 // indirect
	github.com/lestrrat-go/httpcc v1.0.1 // indirect
	github.com/lestrrat-go/iter v1.0.1 // indirect
	github.com/lestrrat-go/option v1.0.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	go.opentelemetry.io/otel/trace v1.13.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230221151758-ace64dc21148 // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/mail.v2 v2.3.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
