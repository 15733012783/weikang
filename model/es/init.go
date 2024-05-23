package es

import (
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetURL("http://120.27.208.86:9200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
}
