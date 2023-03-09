module github.com/ramsfords/backend

go 1.20

require (
	cloud.google.com/go/vision v1.2.0
	github.com/SebastiaanKlippert/go-wkhtmltopdf v1.9.0
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/aws/aws-sdk-go v1.44.210
	github.com/aws/aws-sdk-go-v2 v1.17.5
	github.com/aws/aws-sdk-go-v2/config v1.18.15
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.10.15
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.55
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.22.4
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.18.5
	github.com/aws/aws-sdk-go-v2/service/lambda v1.30.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.30.5
	github.com/aws/aws-sdk-go-v2/service/ses v1.15.3
	github.com/cloudflare/cloudflare-go v0.61.0
	github.com/cloudinary/cloudinary-go/v2 v2.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-mail/mail v2.3.1+incompatible
	github.com/go-resty/resty/v2 v2.7.0
	github.com/goccy/go-json v0.10.0
	github.com/gofor-little/env v1.0.11
	github.com/gofor-little/xerror v1.0.5
	github.com/golang-jwt/jwt/v4 v4.2.0
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/go-retryablehttp v0.7.2
	github.com/jhillyerd/enmime v0.11.0
	github.com/labstack/echo/v5 v5.0.0-20220717203827-74022662be4a
	github.com/labstack/gommon v0.4.0
	github.com/lestrrat-go/jwx v1.2.25
	github.com/newrelic/go-agent/v3 v3.20.3
	github.com/pkg/errors v0.9.1
	github.com/ramsfords/types_gen v0.0.0-20230227204152-1ccdfb287a0e
	github.com/schmorrison/go-querystring v1.1.1
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/viper v1.15.0
	github.com/surendrakandel/supa-go v0.0.0-20230227205758-d0a4713bc4f3
	github.com/valyala/fasthttp v1.44.0
	go.opentelemetry.io/otel v1.13.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.13.0
	go.opentelemetry.io/otel/sdk v1.13.0
	golang.org/x/crypto v0.6.0
	golang.org/x/time v0.1.0
	google.golang.org/genproto v0.0.0-20221227171554-f9683d7f8bef
)

require (
	cloud.google.com/go v0.105.0 // indirect
	cloud.google.com/go/compute v1.14.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/longrunning v0.3.0 // indirect
	cloud.google.com/go/vision/v2 v2.5.0 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.13.15 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.23 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.29 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.23 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.30 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.21 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.14.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.18.5 // indirect
	github.com/aws/smithy-go v1.13.5 // indirect
	github.com/cention-sany/utf7 v0.0.0-20170124080048-26cad61bd60a // indirect
	github.com/creasty/defaults v1.5.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.0-20210816181553-5444fa50b93d // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogs/chardet v0.0.0-20191104214054-4b6791f73a28 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.1 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/gorilla/schema v1.2.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jaytaylor/html2text v0.0.0-20200412013138-3577fbdbcff7 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/lestrrat-go/backoff/v2 v2.0.8 // indirect
	github.com/lestrrat-go/blackmagic v1.0.0 // indirect
	github.com/lestrrat-go/httpcc v1.0.1 // indirect
	github.com/lestrrat-go/iter v1.0.1 // indirect
	github.com/lestrrat-go/option v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nedpals/postgrest-go v0.1.2 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel/trace v1.13.0 // indirect
	golang.org/x/net v0.6.0 // indirect
	golang.org/x/oauth2 v0.0.0-20221014153046-6fdb5e3db783 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/api v0.107.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.52.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/mail.v2 v2.3.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
