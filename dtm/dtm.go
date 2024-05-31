package dtm

import (
	"fmt"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"log"
)

// AffairsDtm dtm分布式事务(需要先启动dtm本地部署原文件dtm-/main.go,可以访问网站后再调用方法)
// 例:Host:Port/127.0.0.1:36790
// 例:rpcHost:rpcPort/127.0.0.1:8080(rpc服务端口)
// 例:proto、protoRollback为proto两个logic文件的路径(score.score/Create&&score.score/DtmCreate)
// payload为api端调用rpc端的结构体例如:&score.CreateScoreReq
func AffairsDtm(Host string, Port int, rpcHost string, rpcPort int, proto, protoRollback string, payload proto.Message) bool {
	gid := uuid.NewString()
	saga := dtmgrpc.NewSagaGrpc(fmt.Sprintf("%s:%d", Host, Port), gid).
		Add(fmt.Sprintf("%s:%d/%s", rpcHost, rpcPort, proto), fmt.Sprintf("%s:%d/%s", rpcHost, rpcPort, protoRollback), payload)
	err := saga.Submit()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
