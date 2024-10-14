package user

import (
	"net/http"

	"github.com/MrLeeang/my-zero/api/internal/logic/user"
	"github.com/MrLeeang/my-zero/api/internal/svc"
	"github.com/MrLeeang/my-zero/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req types.CreateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := user.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
