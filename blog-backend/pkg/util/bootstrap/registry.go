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
)

// NewConsulRegistry 创建一个注册发现客户端 - Consul
func NewConsulRegistry(address, scheme string, healthCheck bool) *consulKratos.Registry {
	conf := consulClient.DefaultConfig()
	conf.Address = address
	conf.Scheme = scheme

	var cli *consulClient.Client
	var err error
	if cli, err = consulClient.NewClient(conf); err != nil {
		log.Fatal(err)
	}

	reg := consulKratos.New(cli, consulKratos.WithHealthCheck(healthCheck))

	return reg
}

// NewEtcdRegistry 创建一个注册发现客户端 - Etcd
func NewEtcdRegistry(address []string) *etcdKratos.Registry {
	conf := etcdClient.Config{
		Endpoints: address,
	}

	var err error
	var cli *etcdClient.Client
	if cli, err = etcdClient.New(conf); err != nil {
		log.Fatal(err)
	}

	reg := etcdKratos.New(cli)

	return reg
}

// NewZooKeeperRegistry 创建一个注册发现客户端 - ZooKeeper
func NewZooKeeperRegistry(address []string) *zookeeperKratos.Registry {
	conn, _, err := zk.Connect(address, time.Second*15)
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
func NewNacosRegistry(address string, port uint64) *nacosKratos.Registry {
	srvConf := []nacosConstant.ServerConfig{
		*nacosConstant.NewServerConfig(address, port),
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
func NewKubernetesRegistry() *k8sRegistry.Registry {
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
func NewEurekaRegistry(address []string) *eurekaKratos.Registry {
	var opts []eurekaKratos.Option
	opts = append(opts, eurekaKratos.WithHeartbeat(time.Second))
	opts = append(opts, eurekaKratos.WithRefresh(time.Second))
	opts = append(opts, eurekaKratos.WithEurekaPath("eureka"))

	var err error
	var reg *eurekaKratos.Registry
	if reg, err = eurekaKratos.New(address, opts...); err != nil {
		log.Fatal(err)
	}

	return reg
}

// NewPolarisRegistry 创建一个注册发现客户端 - Polaris
func NewPolarisRegistry(host string, startPort int, instanceCount int, namespace, service, token string) *polarisKratos.Registry {
	var provider polarisApi.ProviderAPI
	var consumer polarisApi.ConsumerAPI

	var err error

	if consumer, err = polarisApi.NewConsumerAPI(); err != nil {
		log.Fatalf("fail to create consumerAPI, err is %v", err)
	}

	provider = polarisApi.NewProviderAPIByContext(consumer.SDKContext())

	log.Infof("start to register instances, count %d", instanceCount)

	var resp *polarisModel.InstanceRegisterResponse
	for i := 0; i < instanceCount; i++ {
		registerRequest := &polarisApi.InstanceRegisterRequest{}
		registerRequest.Service = service
		registerRequest.Namespace = namespace
		registerRequest.Host = host
		registerRequest.Port = startPort + i
		registerRequest.ServiceToken = token
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
func NewServicecombRegistry(address []string) *servicecombKratos.Registry {
	conf := servicecombClient.Options{
		Endpoints: address,
	}

	var cli *servicecombClient.Client
	var err error
	if cli, err = servicecombClient.NewClient(conf); err != nil {
		log.Fatal(err)
	}

	r := servicecombKratos.NewRegistry(cli)

	return r
}
