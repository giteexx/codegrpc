package main

import (
	"context"
	"log"
	"testing"

	"github.com/giteexx/codegrpc/api"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func TestClient(t *testing.T) {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts:              []string{"127.0.0.1:2379"},
			Key:                "code.rpc",//通道名称
			User:               "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			Pass:               "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			CertFile:           "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			CertKeyFile:        "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			CACertFile:         "",    // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
			InsecureSkipVerify: false, // 当 etcd 开启 acl 时才填写，这里为了展示所以没有删除，实际使用如果没有开启 acl 可忽略
		},
	})
	client := api.NewGatewayClient(conn.Conn())
	in := &api.WayRequest{}
	in.ApiField = &api.ApiField{}
	in.ChannelInfo = &api.ChannelInfo{}
	
	in.OrderField = &api.OrderField{}
	resp, err := client.GetCode(context.Background(), in)
	if err != nil {
		log.Println("链接异常", err)
		return
	}
	log.Println("返回信d息", resp)
}
