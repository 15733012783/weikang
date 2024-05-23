package es

import (
	"fmt"
	"github.com/15733012783/weikang/nacos"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%s:%s", nacos.RpcNac.Elastic.Host, nacos.RpcNac.Elastic.Port)), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
}
