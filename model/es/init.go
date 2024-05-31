package es

import (
	"fmt"
	"github.com/15733012783/weikang/nacos"
	"github.com/olivere/elastic/v7"
	"log"
)

func InitEs() (*elastic.Client, error) {
	es, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%s:%s", nacos.RpcNac.Elastic.Host, nacos.RpcNac.Elastic.Port)), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	log.Println("es连接成功")
	return es, nil
}
