package bootstrap

import (
	"path/filepath"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	// etcd
	etcdKratos "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	etcdClient "go.etcd.io/etcd/client/v3"

	// consul
	consulKratos "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulClient "github.com/hashicorp/consul/api"

	// eureka
	eurekaKratos "github.com/go-kratos/kratos/contrib/registry/eureka/v2"

	// nacos
	nacosKratos "github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	nacosClients "github.com/nacos-group/nacos-sdk-go/clients"
	nacosConstant "github.com/nacos-group/nacos-sdk-go/common/constant"
	nacosVo "github.com/nacos-group/nacos-sdk-go/vo"

	// zookeeper
	zookeeperKratos "github.com/go-kratos/kratos/contrib/registry/zookeeper/v2"
	"github.com/go-zookeeper/zk"

	// kubernetes
	k8sRegistry "github.com/go-kratos/kratos/contrib/registry/kubernetes/v2"
	k8s "k8s.io/client-go/kubernetes"
	k8sRest "k8s.io/client-go/rest"
	k8sTools "k8s.io/client-go/tools/clientcmd"
	k8sUtil "k8s.io/client-go/util/homedir"

	// polaris
	polarisKratos "github.com/go-kratos/kratos/contrib/registry/polaris/v2"
	polarisApi "github.com/polarismesh/polaris-go/api"
	polarisModel "github.com/polarismesh/polaris-go/pkg/model"

	// servicecomb
	servicecombClient "github.com/go-chassis/sc-client"
	servicecombKratos "github.com/go-kratos/kratos/contrib/registry/servicecomb/v2"

	"kratos-blog/gen/api/go/common/conf"
)

// NewConsulRegistry 创建一个注册发现客户端 - Consul
func NewConsulRegistry(c *conf.Registry) *consulKratos.Registry {
	cfg := consulClient.DefaultConfig()
	cfg.Address = c.Consul.Address
	cfg.Scheme = c.Consul.Scheme

	var cli *consulClient.Client
	var err error
	if cli, err = consulClient.NewClient(cfg); err != nil {
		log.Fatal(err)
	}

	reg := consulKratos.New(cli, consulKratos.WithHealthCheck(c.Consul.HealthCheck))

	return reg
}

// NewEtcdRegistry 创建一个注册发现客户端 - Etcd
func NewEtcdRegistry(c *conf.Registry) *etcdKratos.Registry {
	cfg := etcdClient.Config{
		Endpoints: c.Etcd.Endpoints,
	}

	var err error
	var cli *etcdClient.Client
	if cli, err = etcdClient.New(cfg); err != nil {
		log.Fatal(err)
	}

	reg := etcdKratos.New(cli)

	return reg
}

// NewZooKeeperRegistry 创建一个注册发现客户端 - ZooKeeper
func NewZooKeeperRegistry(c *conf.Registry) *zookeeperKratos.Registry {
	conn, _, err := zk.Connect(c.Zookeeper.Endpoints, c.Zookeeper.Timeout.AsDuration())
	if err != nil {
		log.Fatal(err)
	}

	reg := zookeeperKratos.New(conn)
	if err != nil {
		log.Fatal(err)
	}

	return reg
}

// NewNacosRegistry 创建一个注册发现客户端 - Nacos
func NewNacosRegistry(c *conf.Registry) *nacosKratos.Registry {
	srvConf := []nacosConstant.ServerConfig{
		*nacosConstant.NewServerConfig(c.Nacos.Address, c.Nacos.Port),
	}

	cliConf := nacosConstant.ClientConfig{
		NamespaceId:          "public",
		TimeoutMs:            10 * 1000, // http请求超时时间，单位毫秒
		BeatInterval:         5 * 1000,  // 心跳间隔时间，单位毫秒
		UpdateThreadNum:      20,        // 更新服务的线程数
		LogLevel:             "debug",
		CacheDir:             "../../configs/cache", // 缓存目录
		LogDir:               "../../configs/log",   // 日志目录
		NotLoadCacheAtStart:  true,                  // 在启动时不读取本地缓存数据，true--不读取，false--读取
		UpdateCacheWhenEmpty: true,                  // 当服务列表为空时是否更新本地缓存，true--更新,false--不更新
	}

	cli, err := nacosClients.NewNamingClient(
		nacosVo.NacosClientParam{
			ClientConfig:  &cliConf,
			ServerConfigs: srvConf,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	reg := nacosKratos.New(cli)

	return reg
}

// NewKubernetesRegistry 创建一个注册发现客户端 - Kubernetes
func NewKubernetesRegistry(_ *conf.Registry) *k8sRegistry.Registry {
	restConfig, err := k8sRest.InClusterConfig()
	if err != nil {
		home := k8sUtil.HomeDir()
		kubeConfig := filepath.Join(home, ".kube", "config")
		restConfig, err = k8sTools.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}

	clientSet, err := k8s.NewForConfig(restConfig)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	reg := k8sRegistry.NewRegistry(clientSet)

	return reg
}

// NewEurekaRegistry 创建一个注册发现客户端 - Eureka
func NewEurekaRegistry(c *conf.Registry) *eurekaKratos.Registry {
	var opts []eurekaKratos.Option
	opts = append(opts, eurekaKratos.WithHeartbeat(time.Second))
	opts = append(opts, eurekaKratos.WithRefresh(time.Second))
	opts = append(opts, eurekaKratos.WithEurekaPath("eureka"))

	var err error
	var reg *eurekaKratos.Registry
	if reg, err = eurekaKratos.New(c.Eureka.Endpoints, opts...); err != nil {
		log.Fatal(err)
	}

	return reg
}

// NewPolarisRegistry 创建一个注册发现客户端 - Polaris
func NewPolarisRegistry(c *conf.Registry) *polarisKratos.Registry {
	var provider polarisApi.ProviderAPI
	var consumer polarisApi.ConsumerAPI

	var err error

	if consumer, err = polarisApi.NewConsumerAPI(); err != nil {
		log.Fatalf("fail to create consumerAPI, err is %v", err)
	}

	provider = polarisApi.NewProviderAPIByContext(consumer.SDKContext())

	log.Infof("start to register instances, count %d", c.Polaris.InstanceCount)

	var resp *polarisModel.InstanceRegisterResponse
	for i := 0; i < (int)(c.Polaris.InstanceCount); i++ {
		registerRequest := &polarisApi.InstanceRegisterRequest{}
		registerRequest.Service = c.Polaris.Service
		registerRequest.Namespace = c.Polaris.Namespace
		registerRequest.Host = c.Polaris.Address
		registerRequest.Port = (int)(c.Polaris.Port) + i
		registerRequest.ServiceToken = c.Polaris.Token
		registerRequest.SetHealthy(true)
		if resp, err = provider.RegisterInstance(registerRequest); err != nil {
			log.Fatalf("fail to register instance %d, err is %v", i, err)
		} else {
			log.Infof("register instance %d response: instanceId %s", i, resp.InstanceID)
		}
	}

	reg := polarisKratos.NewRegistry(provider, consumer)

	return reg
}

// NewServicecombRegistry 创建一个注册发现客户端 - Servicecomb
func NewServicecombRegistry(c *conf.Registry) *servicecombKratos.Registry {
	cfg := servicecombClient.Options{
		Endpoints: c.Servicecomb.Endpoints,
	}

	var cli *servicecombClient.Client
	var err error
	if cli, err = servicecombClient.NewClient(cfg); err != nil {
		log.Fatal(err)
	}

	reg := servicecombKratos.NewRegistry(cli)

	return reg
}
