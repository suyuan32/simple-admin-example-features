// Code generated by goctl. DO NOT EDIT.
// Source: example.proto

package exampleclient

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp   = example.BaseIDResp
	BaseResp     = example.BaseResp
	BaseUUIDResp = example.BaseUUIDResp
	Empty        = example.Empty
	IDReq        = example.IDReq
	IDsReq       = example.IDsReq
	PageInfoReq  = example.PageInfoReq
	UUIDReq      = example.UUIDReq
	UUIDsReq     = example.UUIDsReq

	Example interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultExample struct {
		cli zrpc.Client
	}
)

func NewExample(cli zrpc.Client) Example {
	return &defaultExample{
		cli: cli,
	}
}

func (m *defaultExample) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := example.NewExampleClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}