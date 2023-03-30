// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	base "swagger/internal/handler/base"
	"swagger/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/init/database",
				Handler: base.InitDatabaseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/init/test/:name/:age",
				Handler: base.TestPathHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/init/test_query",
				Handler: base.TestQueryHandler(serverCtx),
			},
		},
	)
}
