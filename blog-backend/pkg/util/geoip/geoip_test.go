package geoip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeoIp(t *testing.T) {
	g, err := NewGeoIp()
	assert.Nil(t, err)

	ret, err := g.GetAreaFromIP("47.108.149.89")
	assert.Nil(t, err)
	assert.Equal(t, ret.Country, "中国")
	assert.Equal(t, ret.Province, "四川省")
	assert.Equal(t, ret.City, "成都")
}
