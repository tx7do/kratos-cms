module github.com/tx7do/kratos-blog/blog-backend

go 1.18

require (
	entgo.io/contrib v0.3.4
	entgo.io/ent v0.11.6
	github.com/go-kratos/kratos/contrib/config/consul/v2 v2.0.0-20221130041748-2cf82fa4a75c
	github.com/go-kratos/kratos/contrib/config/etcd/v2 v2.0.0-20221130041748-2cf82fa4a75c
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20221130041748-2cf82fa4a75c
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20221130041748-2cf82fa4a75c
	github.com/go-kratos/kratos/v2 v2.5.4
	github.com/go-kratos/swagger-api v1.0.1
	github.com/go-redis/redis/extra/redisotel/v8 v8.11.5
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.7.0
	github.com/golang-jwt/jwt/v4 v4.4.3
	github.com/google/gnostic v0.6.9
	github.com/google/subcommands v1.2.0
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/gosimple/slug v1.13.1
	github.com/hashicorp/consul/api v1.17.0
	github.com/jackc/pgx/v4 v4.17.2
	github.com/lib/pq v1.10.7
	github.com/nacos-group/nacos-sdk-go v1.1.3
	github.com/olekukonko/tablewriter v0.0.5
	github.com/spf13/cobra v1.6.1
	github.com/stretchr/testify v1.8.1
	github.com/tx7do/kratos-authn/authn v0.0.0-20221130123508-fed84d7952e9
	github.com/tx7do/kratos-authn/engine/jwt v0.0.0-20221220153430-9d1496472c00
	github.com/tx7do/kratos-authn/middleware v0.0.0-20221220161625-08fd42ab33cc
	github.com/tx7do/kratos-authz v0.0.4
	github.com/tx7do/kratos-authz/middleware v0.0.0-20221220154012-95ab9f84cac0
	go.etcd.io/etcd/client/v3 v3.5.6
	go.opentelemetry.io/otel v1.11.2
	go.opentelemetry.io/otel/exporters/jaeger v1.11.1
	go.opentelemetry.io/otel/sdk v1.11.2
	golang.org/x/crypto v0.3.0
	golang.org/x/tools v0.3.1-0.20221202221704-aa9f4b2f3d57
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37
	google.golang.org/grpc v1.52.0
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/clickhouse v0.5.0
	gorm.io/driver/mysql v1.4.4
	gorm.io/driver/postgres v1.4.5
	gorm.io/driver/sqlite v1.4.3
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.24.2
)

require (
	ariga.io/atlas v0.9.1-0.20230119145809-92243f7c55cb // indirect
	github.com/ClickHouse/ch-go v0.51.1 // indirect
	github.com/ClickHouse/clickhouse-go/v2 v2.5.1 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.18 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-faster/city v1.0.1 // indirect
	github.com/go-faster/errors v0.6.1 // indirect
	github.com/go-kratos/grpc-gateway/v2 v2.5.1-0.20210811062259-c92d36e434b1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/go-redis/redis/extra/rediscmd/v8 v8.11.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.14.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-hclog v0.14.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jhump/protoreflect v1.10.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.14 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/microsoft/go-mssqldb v0.17.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/paulmach/orb v0.8.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/rivo/uniseg v0.4.3 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tx7do/kratos-authn v0.0.6 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.6 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.6 // indirect
	go.opentelemetry.io/otel/trace v1.11.2 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	gopkg.in/ini.v1 v1.42.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
