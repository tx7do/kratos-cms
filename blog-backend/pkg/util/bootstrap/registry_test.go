package bootstrap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConsulRegistry(t *testing.T) {
	reg := NewConsulRegistry("host.docker.internal:8500", "http", false)
	assert.Nil(t, reg)
}

func TestNewEtcdRegistry(t *testing.T) {
	reg := NewEtcdRegistry([]string{"127.0.0.1:2379"})
	assert.Nil(t, reg)
}

func TestNewNacosRegistry(t *testing.T) {
	reg := NewNacosRegistry("127.0.0.1", 8848)
	assert.Nil(t, reg)
}

func TestNewZooKeeperRegistry(t *testing.T) {
	reg := NewZooKeeperRegistry([]string{"127.0.0.1:2181"})
	assert.Nil(t, reg)
}

func TestNewKubernetesRegistry(t *testing.T) {
	reg := NewKubernetesRegistry()
	assert.Nil(t, reg)
}

func TestNewEurekaRegistry(t *testing.T) {
	reg := NewEurekaRegistry([]string{"https://127.0.0.1:18761"})
	assert.Nil(t, reg)
}

func TestNewPolarisRegistry(t *testing.T) {
	reg := NewPolarisRegistry("127.0.0.1", 8091, 5, "default", "DiscoverEchoServer", "")
	assert.Nil(t, reg)
}

func TestNewServicecombRegistry(t *testing.T) {
	reg := NewServicecombRegistry([]string{"127.0.0.1:30100"})
	assert.Nil(t, reg)
}
