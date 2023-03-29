package base

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-example-mongo/internal/logic/base"
	"github.com/suyuan32/simple-admin-example-mongo/internal/svc"
)

// swagger:route get /init/create base CreateExample
//
// Create example | 创建的例子
//
// Create example | 创建的例子
//
// Responses:
//  200: BaseMsgResp

func CreateExampleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewCreateExampleLogic(r.Context(), svcCtx)
		resp, err := l.CreateExample()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
