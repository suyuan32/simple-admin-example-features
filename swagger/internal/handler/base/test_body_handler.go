package base

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"swagger/internal/logic/base"
	"swagger/internal/svc"
	"swagger/internal/types"
)

// swagger:route post /init/test_body base TestBody
//
// Body Params Test | 表格参数测试
//
// Body Params Test | 表格参数测试
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: BodyType
//
// Responses:
//  200: BaseMsgResp

func TestBodyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BodyType
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := base.NewTestBodyLogic(r.Context(), svcCtx)
		resp, err := l.TestBody(&req)
		if err != nil {

			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
