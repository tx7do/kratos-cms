package geoip

import (
	_ "embed"
	"errors"
	"net"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/oschwald/geoip2-golang"
)

//go:embed GeoLite2-City.mmdb
var geoipMmdbByte []byte

type Result struct {
	Country  string // 国家
	Province string // 省
	City     string // 城市
}

// GeoIp 地理位置解析结构体
type GeoIp struct {
	db             *geoip2.Reader
	outputLanguage string
}

// NewGeoIp .
func NewGeoIp() (geoip *GeoIp, err error) {
	db, err := geoip2.FromBytes(geoipMmdbByte)
	if err != nil {
		return nil, err
	}
	return &GeoIp{db: db, outputLanguage: "zh-CN"}, nil
}

func (g *GeoIp) Close() (err error) {
	return g.db.Close()
}

// SetLanguage 设置输出的语言，默认为：zh-CN
func (g *GeoIp) SetLanguage(code string) {
	g.outputLanguage = code
}

func (g *GeoIp) query(rawIP string) (city *geoip2.City, err error) {
	ip := net.ParseIP(rawIP)
	if ip == nil {
		return nil, errors.New("invalid ip")
	}

	return g.db.City(ip)
}

// GetAreaFromIP 通过IP获取地区
func (g *GeoIp) GetAreaFromIP(rawIP string) (ret Result, err error) {
	record, err := g.query(rawIP)
	if err != nil {
		log.Fatal(err)
		return ret, err
	}

	ret.Country = record.Country.Names[g.outputLanguage]
	if len(record.Subdivisions) > 0 {
		ret.Province = record.Subdivisions[0].Names[g.outputLanguage]
	}
	ret.City = record.City.Names[g.outputLanguage]

	return
}
