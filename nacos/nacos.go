package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
)

type ConfigNaCosRpc struct {
	Port  string `json:"port" yaml:"port"`
	Mysql struct {
		Root     string `json:"root" yaml:"root"`
		Password string `json:"password" yaml:"password"`
		Host     string `json:"host" yaml:"host"`
		Port     string `json:"port" yaml:"port"`
		Database string `json:"database" yaml:"database"`
	} `json:"mysql" yaml:"mysql"`
	Elastic struct {
		Host string `json:"host" yaml:"host"`
		Port int    `json:"port" yaml:"port"`
	} `json:"service" yaml:"elastic"`
	JwtSigningKey string `json:"jwtSigningKey" yaml:"jwtSigningKey"`
	Consul        struct {
		IpAddr string `json:"IpAddr" yaml:"ipAddr"`
	} `json:"consul" yaml:"consul"`
	Redis struct {
		Addr     string `json:"addr" yaml:"addr"`
		Password string `json:"password" yaml:"password"`
		Db       string `json:"db" yaml:"db"`
	} `json:"redis" yaml:"redis"`
	Mongodb struct {
		Host string `json:"host" yaml:"host"`
		Port string `json:"port" yaml:"port"`
	} `json:"mongodb" yaml:"mongodb"`
}

type ConfigNaCosApi struct {
	ServerName string `json:"serverName" yaml:"serverName"`
	Port       int    `json:"port" yaml:"port"`
	Jwt        struct {
		SigningKey string `json:"signingKey" yaml:"signingKey"`
	} `json:"jwt" yaml:"jwt"`
	Consul struct {
		Host string `json:"host" yaml:"host"`
		Port int    `json:"port" yaml:"port"`
	} `json:"consul" yaml:"consul"`
	Rabbitmq struct {
		Root     string `json:"root" yaml:"root"`
		Password string `json:"password" yaml:"password"`
		Host     string `json:"host" yaml:"host"`
		Port     int    `json:"port" yaml:"port"`
	} `json:"rabbitmq" yaml:"rabbitmq"`
	Crypto struct {
		Key int64 `json:"key" yaml:"key"`
	} `json:"crypto" yaml:"crypto"`
	Ali struct {
		AccessKeyID     string `json:"AccessKey_ID" yaml:"accessKeyID"`
		AccessKeySecret string `json:"AccessKey_Secret" yaml:"accessKeySecret"`
		Endpoint        string `json:"endpoint" yaml:"endpoint"`
	} `json:"ali" yaml:"ali"`
}

func NaCos(DataId, Group, ip string, NamespaceId string) (string, error) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// At least one ServerConfig>protoc --version
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      ip,
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}
	// Create naming service for service discovery
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return "", err
	}
	config, err := client.GetConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  Group,
	})
	if err != nil {
		return "", err
	}
	return config, nil
}

var RpcNac ConfigNaCosRpc

func ServiceNaCos(dataid, group, host, NamespaceId string) error {
	cos, err := NaCos(dataid, group, host, NamespaceId)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(cos), &RpcNac)
	if err != nil {
		return err
	}
	return nil
}

var ApiNac ConfigNaCosApi

func ClientNaCos(dataid, group, host, NamespaceId string) error {
	cos, err := NaCos(dataid, group, host, NamespaceId)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(cos), &ApiNac)
	if err != nil {
		return err
	}
	return nil
}
