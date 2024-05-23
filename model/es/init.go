package es

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"weikang_database/nacos"
)

var client *elastic.Client

func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%s:%s", nacos.RpcNac.Elastic.Host, nacos.RpcNac.Elastic.Port)), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
}
