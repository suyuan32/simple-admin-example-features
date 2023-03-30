package base

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"swagger/internal/logic/base"
	"swagger/internal/svc"
	"swagger/internal/types"
)

// swagger:route get /init/test/{name}/{age} base TestPath
//
// Path Params Test | 路径参数测试
//
// Path Params Test | 路径参数测试
//
// Responses:
//  200: BaseMsgResp

func TestPathHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathType
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := base.NewTestPathLogic(r.Context(), svcCtx)
		resp, err := l.TestPath(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
