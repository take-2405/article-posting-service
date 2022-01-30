package middleware

import (
	"fmt"
	"net/http"
	"prac-orm-transaction/infrastructure"
	"prac-orm-transaction/infrastructure/table"
	"prac-orm-transaction/presentation/response"
)

func Auth() (fn func(http.Handler) http.Handler) { // 引数名を指定してるのでreturnのみでおｋ
	fn = func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Token")

			var dataExistsCheck table.UserInfo

			conn := infrastructure.NewMysqlRepository()
			conn.Client.First(&dataExistsCheck, "token=?", token)
			if dataExistsCheck.ID == "" {
				response.RespondError(w, http.StatusUnauthorized, fmt.Errorf("利用権限がありません"))
			}

			r.Header.Set("userID", dataExistsCheck.ID)

			// 何も無ければ次のハンドラを実行
			h.ServeHTTP(w, r)
		})
	}
	return
}
