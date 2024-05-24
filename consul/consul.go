package consul

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

func RegisterConSul(Ghost string, Gport int, Host string, Port int, Name string) {
	var err error
	sprintf := fmt.Sprintf("%v:%v", Host, Port)
	ConsulCli, err := api.NewClient(&api.Config{
		Address: sprintf,
	})
	if err != nil {
		log.Println(err, "服务注册失败")
		return
	}
	Srvid := uuid.New().String()
	check := &api.AgentServiceCheck{
		Interval:                       "5s",
		Timeout:                        "5s",
		GRPC:                           fmt.Sprintf("%s:%d", Ghost, Gport),
		DeregisterCriticalServiceAfter: "30s",
	}
	err = ConsulCli.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      Srvid,
		Name:    Name,
		Tags:    []string{"GRPC"},
		Port:    Gport,
		Address: Ghost,
		Check:   check,
	})
	if err != nil {
		log.Println(err, "服务注册失败")
		return
	}
	return
}

var currentIndex int

func GetConSul(serverName, Address string) (string, error) {
	cc, err := api.NewClient(&api.Config{
		Address: Address,
	})
	if err != nil {
		fmt.Printf("api.NewClient failed, err:%v\n", err)
		return "", err
	}
	serviceMap, date, err := cc.Agent().AgentHealthServiceByName(serverName)
	if serviceMap != "passing" {
		log.Println("获取consul服务发现失败***！", err)
		return "", err
	}
	// 选一个服务机（这里选最后一个）
	if len(date) == 0 {
		return "", errors.New("没有可用的服务")
	}
	// 获取当前要访问的服务的索引
	currentIndex = (currentIndex + 1) % len(date)

	// 获取当前要访问的服务地址
	selectedService := date[currentIndex]
	addr := selectedService.Service.Address + ":" + strconv.Itoa(selectedService.Service.Port)
	fmt.Println(addr)
	return addr, nil
}

// 服务过滤
func RegisterConsul(Ghost string, Gport int, Host string, Port int, Name string) {
	sprintf := fmt.Sprintf("%v:%v", Host, Port)
	client, err := api.NewClient(&api.Config{
		Address: sprintf,
	})
	registration := api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    Name,
		Tags:    []string{"GRPC"},
		Port:    Port,
		Address: Host,
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "5s",
			GRPC:                           fmt.Sprintf("%s:%d", Ghost, Gport),
			DeregisterCriticalServiceAfter: "30s",
		},
	}
	result, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, Name))
	if err != nil {
		log.Panic("consul过滤服务失败", err)
		return
	}
	var BaseSrvAddr string
	for _, val := range result {
		if val.Address == fmt.Sprintf("%s:%d", Ghost, Gport) {
			BaseSrvAddr = val.Address
			log.Println("consul服务已存在")
			break
		}
	}
	if BaseSrvAddr == "" {
		err = client.Agent().ServiceRegister(&registration)
		if err != nil {
			log.Fatal("consul注册服务失败", err)
			return
		}
		log.Println("consul服务注册完成")
	}
}
