module kratos-cms

go 1.23.0

toolchain go1.23.2

replace (
	github.com/armon/go-metrics => github.com/hashicorp/go-metrics v0.4.1
	github.com/imdario/mergo => dario.cat/mergo v0.3.16
)

require (
	entgo.io/ent v0.14.1
	github.com/go-kratos/kratos/v2 v2.8.0
	github.com/go-sql-driver/mysql v1.8.1
	github.com/google/gnostic v0.7.0
	github.com/google/uuid v1.6.0
	github.com/google/wire v0.6.0
	github.com/jackc/pgx/v4 v4.18.3
	github.com/lib/pq v1.10.9
	github.com/redis/go-redis/v9 v9.6.1
	github.com/tx7do/go-utils v1.1.11
	github.com/tx7do/go-utils/entgo v1.1.14
	github.com/tx7do/kratos-authn v1.0.0
	github.com/tx7do/kratos-authn/engine/jwt v1.0.0
	github.com/tx7do/kratos-authn/middleware v1.0.0
	github.com/tx7do/kratos-authz v1.0.0
	github.com/tx7do/kratos-authz/middleware v1.0.0
	github.com/tx7do/kratos-bootstrap v0.3.9
	github.com/tx7do/kratos-bootstrap/api v0.0.4
	github.com/tx7do/kratos-bootstrap/cache/redis v0.0.1
	github.com/tx7do/kratos-bootstrap/oss/minio v0.0.1
	github.com/tx7do/kratos-swagger-ui v0.0.0-20240514132937-bdf9b35be16c
	google.golang.org/genproto/googleapis/api v0.0.0-20241007155032-5fefd90f89a9
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.35.1
)

require (
	ariga.io/atlas v0.28.1 // indirect
	dario.cat/mergo v1.0.1 // indirect
	entgo.io/contrib v0.6.0 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/XSAM/otelsql v0.35.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.63.27 // indirect
	github.com/apolloconfig/agollo/v4 v4.4.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/armon/go-metrics v0.5.3 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/bufbuild/protocompile v0.14.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/emicklei/go-restful/v3 v3.12.1 // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fluent/fluent-logger-golang v1.9.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/go-chassis/cari v0.9.0 // indirect
	github.com/go-chassis/foundation v0.4.0 // indirect
	github.com/go-chassis/openlog v1.1.3 // indirect
	github.com/go-chassis/sc-client v0.7.0 // indirect
	github.com/go-errors/errors v1.5.1 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-kratos/kratos/contrib/config/apollo/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/config/etcd/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/config/kubernetes/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/log/fluent/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/log/logrus/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/log/tencent/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/log/zap/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/registry/eureka/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/registry/kubernetes/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/registry/servicecomb/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-kratos/kratos/contrib/registry/zookeeper/v2 v2.0.0-20241011140022-7cd86eb077dc // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/inflect v0.21.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/go-zookeeper/zk v1.0.4 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gnostic-models v0.6.9-0.20230804172637-c7be7c783f49 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/subcommands v1.2.0 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 // indirect
	github.com/hashicorp/consul/api v1.29.4 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.22.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/imdario/mergo v1.0.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgtype v1.14.3 // indirect
	github.com/jhump/protoreflect v1.17.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karlseguin/ccache/v2 v2.0.8 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/lufia/plan9stats v0.0.0-20240909124753-873cd0166683 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/minio/minio-go/v7 v7.0.70 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nacos-group/nacos-sdk-go v1.1.5 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/redis/go-redis/extra/rediscmd/v9 v9.0.5 // indirect
	github.com/redis/go-redis/extra/redisotel/v9 v9.0.5 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/sagikazarmark/locafero v0.6.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/shirou/gopsutil/v3 v3.24.5 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/sony/sonyflake v1.2.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.7.0 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.19.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/swaggest/swgui v1.8.2 // indirect
	github.com/tencentcloud/tencentcloud-cls-sdk-go v1.0.11 // indirect
	github.com/tinylib/msgp v1.2.2 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.9.0 // indirect
	github.com/tx7do/kratos-bootstrap/config/apollo v0.0.1 // indirect
	github.com/tx7do/kratos-bootstrap/config/consul v0.0.2 // indirect
	github.com/tx7do/kratos-bootstrap/config/etcd v0.0.1 // indirect
	github.com/tx7do/kratos-bootstrap/config/kubernetes v0.0.1 // indirect
	github.com/tx7do/kratos-bootstrap/config/nacos v0.0.1 // indirect
	github.com/tx7do/kratos-bootstrap/config/polaris v0.0.1 // indirect
	github.com/tx7do/kratos-bootstrap/logger/fluent v0.0.2 // indirect
	github.com/tx7do/kratos-bootstrap/logger/logrus v0.0.2 // indirect
	github.com/tx7do/kratos-bootstrap/logger/tencent v0.0.2 // indirect
	github.com/tx7do/kratos-bootstrap/logger/zap v0.0.2 // indirect
	github.com/tx7do/kratos-bootstrap/registry/consul v0.0.1 // indirect
	github.com/tx7do/kratos-bootstrap/registry/etcd v0.0.1 // indirect
	github.com/tx7do/kratos-bootstrap/registry/eureka v0.0.1 // indirect
	github.com/tx7do/kratos-bootstrap/registry/kubernetes v0.0.1 // indirect
	github.com/tx7do/kratos-bootstrap/registry/nacos v0.0.1 // indirect
	github.com/tx7do/kratos-bootstrap/registry/servicecomb v0.0.1 // indirect
	github.com/tx7do/kratos-bootstrap/registry/zookeeper v0.0.1 // indirect
	github.com/vearutop/statigz v1.4.3 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	github.com/zclconf/go-cty v1.15.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.16 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.16 // indirect
	go.etcd.io/etcd/client/v3 v3.5.16 // indirect
	go.opentelemetry.io/otel v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.31.0 // indirect
	go.opentelemetry.io/otel/metric v1.31.0 // indirect
	go.opentelemetry.io/otel/sdk v1.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.31.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/oauth2 v0.23.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/term v0.25.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	golang.org/x/tools v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241007155032-5fefd90f89a9 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.31.1 // indirect
	k8s.io/apimachinery v0.31.1 // indirect
	k8s.io/client-go v0.31.1 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20241009091222-67ed5848f094 // indirect
	k8s.io/utils v0.0.0-20240921022957-49e7df575cb6 // indirect
	sigs.k8s.io/json v0.0.0-20241010143419-9aa6b5e7a4b3 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
