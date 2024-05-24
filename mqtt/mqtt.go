package Mqtt

import (
	"fmt"
	"github.com/15733012783/weikang/nacos"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

// 定义消息处理函数，用于处理接收到的消息
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// 定义连接成功处理函数，打印连接成功信息
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

// 定义连接丢失处理函数，打印连接丢失信息
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func Mqtt(text interface{}) {
	// MQTT代理地址和端口
	var broker = nacos.ApiNac.Mqtt.Broker
	var port = nacos.ApiNac.Mqtt.Port

	// 创建MQTT客户端选项
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", broker, port))
	opts.SetClientID(nacos.ApiNac.Mqtt.ClientID)
	opts.SetUsername(nacos.ApiNac.Mqtt.Username)
	opts.SetPassword(nacos.ApiNac.Mqtt.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	// 创建MQTT客户端
	client := mqtt.NewClient(opts)

	// 连接到MQTT代理
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题并发布消息
	sub(client)
	publish(client, text)

	// 断开与MQTT代理的连接
	client.Disconnect(250)
}

// 发布消息函数
func publish(client mqtt.Client, text interface{}) {
	token := client.Publish("topic/test", 0, false, text)
	token.Wait()
	time.Sleep(time.Second)
}

// 订阅主题函数
func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}
