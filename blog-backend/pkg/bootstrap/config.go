package bootstrap

import (
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"

	// file
	fileKratos "github.com/go-kratos/kratos/v2/config/file"

	// etcd
	etcdKratos "github.com/go-kratos/kratos/contrib/config/etcd/v2"
	etcdClient "go.etcd.io/etcd/client/v3"

	// consul
	consulKratos "github.com/go-kratos/kratos/contrib/config/consul/v2"
	consulApi "github.com/hashicorp/consul/api"

	// nacos
	nacosKratos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	nacosClients "github.com/nacos-group/nacos-sdk-go/clients"
	nacosConstant "github.com/nacos-group/nacos-sdk-go/common/constant"
	nacosVo "github.com/nacos-group/nacos-sdk-go/vo"

	// polaris
	polarisKratos "github.com/go-kratos/kratos/contrib/config/polaris/v2"
	polarisApi "github.com/polarismesh/polaris-go"

	// apollo
	apolloKratos "github.com/go-kratos/kratos/contrib/config/apollo/v2"

	// kubernetes
	k8sKratos "github.com/go-kratos/kratos/contrib/config/kubernetes/v2"
	k8sUtil "k8s.io/client-go/util/homedir"
)

// NewConfigProvider 创建一个配置
func NewConfigProvider(configType, configHost, configPath, configKey string) config.Config {
	return config.New(
		config.WithSource(
			NewFileConfigSource(configPath),
			NewRemoteConfigSource(configType, configHost, configKey),
		),
	)
}

// NewRemoteConfigSource 创建一个远程配置源
func NewRemoteConfigSource(configType, configHost, configKey string) config.Source {
	switch configType {
	case "nacos":
		uri, _ := url.Parse(configHost)
		h := strings.Split(uri.Host, ":")
		addr := h[0]
		port, _ := strconv.Atoi(h[1])
		return NewNacosConfigSource(addr, uint64(port), configKey)
	case "consul":
		return NewConsulConfigSource(configHost, configKey)
	case "etcd":
		return NewEtcdConfigSource(configHost, configKey)
	case "apollo":
		return NewApolloConfigSource(configHost, configKey)
	}
	return nil
}

// getConfigKey 获取合法的配置名
func getConfigKey(configKey string, useBackslash bool) string {
	if useBackslash {
		return strings.Replace(configKey, `.`, `/`, -1)
	} else {
		return configKey
	}
}

// NewFileConfigSource 创建一个本地文件配置源
func NewFileConfigSource(filePath string) config.Source {
	return fileKratos.NewSource(filePath)
}

// NewNacosConfigSource 创建一个远程配置源 - Nacos
func NewNacosConfigSource(configAddr string, configPort uint64, configKey string) config.Source {
	srvConf := []nacosConstant.ServerConfig{
		*nacosConstant.NewServerConfig(configAddr, configPort),
	}

	cliConf := nacosConstant.ClientConfig{
		TimeoutMs:            10 * 1000, // http请求超时时间，单位毫秒
		BeatInterval:         5 * 1000,  // 心跳间隔时间，单位毫秒
		UpdateThreadNum:      20,        // 更新服务的线程数
		LogLevel:             "debug",
		CacheDir:             "../../configs/cache", // 缓存目录
		LogDir:               "../../configs/log",   // 日志目录
		NotLoadCacheAtStart:  true,                  // 在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: true,                  // 当服务列表为空时是否更新本地缓存，true--更新,false--不更新
	}

	nacosClient, err := nacosClients.NewConfigClient(
		nacosVo.NacosClientParam{
			ClientConfig:  &cliConf,
			ServerConfigs: srvConf,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return nacosKratos.NewConfigSource(nacosClient,
		nacosKratos.WithGroup(getConfigKey(configKey, false)),
		nacosKratos.WithDataID("bootstrap.yaml"),
	)
}

// NewEtcdConfigSource 创建一个远程配置源 - Etcd
func NewEtcdConfigSource(configHost, configKey string) config.Source {
	conf := etcdClient.Config{
		Endpoints:   []string{configHost},
		DialTimeout: time.Second, DialOptions: []grpc.DialOption{grpc.WithBlock()},
	}

	cli, err := etcdClient.New(conf)
	if err != nil {
		panic(err)
	}

	source, err := etcdKratos.New(cli, etcdKratos.WithPath(getConfigKey(configKey, true)))
	if err != nil {
		log.Fatal(err)
	}

	return source
}

// NewConsulConfigSource 创建一个远程配置源 - Consul
func NewConsulConfigSource(configHost, configKey string) config.Source {
	conf := &consulApi.Config{
		Address: configHost,
	}

	cli, err := consulApi.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	source, err := consulKratos.New(cli,
		consulKratos.WithPath(getConfigKey(configKey, true)),
	)
	if err != nil {
		log.Fatal(err)
	}

	return source
}

// NewApolloConfigSource 创建一个远程配置源 - Apollo
func NewApolloConfigSource(_, _ string) config.Source {
	source := apolloKratos.NewSource(
		apolloKratos.WithAppID("kratos"),
		apolloKratos.WithCluster("dev"),
		apolloKratos.WithEndpoint("http://localhost:8080"),
		apolloKratos.WithNamespace("application,event.yaml,demo.json"),
		apolloKratos.WithEnableBackup(),
		apolloKratos.WithSecret("ad75b33c77ae4b9c9626d969c44f41ee"),
	)
	return source
}

// NewKubernetesConfigSource 创建一个远程配置源 - Kubernetes
func NewKubernetesConfigSource(_, _ string) config.Source {
	source := k8sKratos.NewSource(
		k8sKratos.Namespace("mesh"),
		k8sKratos.LabelSelector(""),
		k8sKratos.KubeConfig(filepath.Join(k8sUtil.HomeDir(), ".kube", "config")),
	)
	return source
}

// NewPolarisConfigSource 创建一个远程配置源 - Polaris
func NewPolarisConfigSource(_, _ string) config.Source {
	configApi, err := polarisApi.NewConfigAPI()
	if err != nil {
		log.Fatal(err)
	}

	var opts []polarisKratos.Option
	opts = append(opts, polarisKratos.WithNamespace("default"))
	opts = append(opts, polarisKratos.WithFileGroup("default"))
	opts = append(opts, polarisKratos.WithFileName("default.yaml"))

	source, err := polarisKratos.New(configApi, opts...)
	if err != nil {
		log.Fatal(err)
	}

	return source
}
