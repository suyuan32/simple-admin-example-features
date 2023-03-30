package base

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"swagger/internal/logic/base"
	"swagger/internal/svc"
	"swagger/internal/types"
)

// swagger:route get /init/test_query base TestQuery
//
// Query Params Test | 表格参数测试
//
// Query Params Test | 表格参数测试
//
// Responses:
//  200: BaseMsgResp

func TestQueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryType
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := base.NewTestQueryLogic(r.Context(), svcCtx)
		resp, err := l.TestQuery(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
