package es

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/15733012783/weikang/model/mysql"
	"github.com/olivere/elastic/v7"
	"sort"
)

// DeviceIk todo:添加分词器,kibana也可以直接添加
func DeviceIk() error {
	es, _ := InitEs()
	_, err := es.CreateIndex("car").BodyString(`
			{
		  "settings": {
			"analysis": {
			  "analyzer": {
				"my_analyzer": { 
				  "tokenizer": "ik_smart"
				}
			  }
			}
		  },
		  "mappings": {
			"properties": {
			  "Name": {
				"type": "text",
				"analyzer": "my_analyzer"
			  },
				"Model": {
				"type": "text",
				"analyzer": "my_analyzer"
			  }
			}
		  }
		}
	`).Do(context.Background())
	if err != nil {
		return errors.New("添加分词器失败")
	}
	return nil
}

// DeviceIndex todo:生成设备索引
func DeviceIndex(id string, device *mysql.Device) (err error) {
	es, _ := InitEs()
	_, err = es.Index().Index("weikang_device").Id(id).BodyJson(device).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// DeviceIkEs  todo:分词搜索并且高亮显示
func DeviceIkEs(name string) []map[string]interface{} {
	es, _ := InitEs()
	do, _ := es.Search().Index("weikang_device").Query(elastic.NewMatchQuery("Name", name).Analyzer("my_analyzer")).Highlight(elastic.NewHighlight().Field("Stores").PreTags("<span color ='red'>").PostTags("</span>")).Do(context.Background())
	var data []map[string]interface{}
	for _, v := range do.Hits.Hits {
		var d map[string]interface{}
		json.Unmarshal(v.Source, &d)
		d["Name"] = v.Highlight
		data = append(data, d)
	}
	return data

}

// DeviceFullSearch todo:关键词全文模糊搜索
func DeviceFullSearch(text string) []map[string]interface{} {
	es, _ := InitEs()
	do, _ := es.Search().Index("weikang_device").Query(elastic.NewQueryStringQuery(text)).Do(context.Background())
	var data []map[string]interface{}
	for _, v := range do.Hits.Hits {
		var d map[string]interface{}
		json.Unmarshal(v.Source, &d)
		d["CarSale"] = v.Highlight
		data = append(data, d)
	}
	return data
}

// DeviceGroupCont todo:聚合查询,按照品牌分组,每组数量
func DeviceGroupCont() {
	es, _ := InitEs()
	do, _ := es.Search().Index("weikang_device").Query(elastic.NewMatchAllQuery()).Aggregation("Brand", elastic.NewTermsAggregation().Field("Brand").SubAggregation("Stock", elastic.NewValueCountAggregation())).Do(context.Background())
	var kvs []DeviceKeyValue
	ter, _ := do.Aggregations.Terms("Brand")
	for _, v := range ter.Buckets {
		k := v.Key
		s, _ := v.Aggregations.ValueCount("Brand")
		kv := DeviceKeyValue{
			Key:   int(k.(float64)),
			Value: int(*s.Value),
		}
		kvs = append(kvs, kv)
	}
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value > kvs[j].Value
	})
	fmt.Println(kvs)
	var a []int
	for _, v := range kvs {
		a = append(a, v.Key)
	}
	fmt.Println(a)
}

type DeviceKeyValue struct {
	Key   int
	Value int
}
