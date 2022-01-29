package middleware

import(
	"net/http"
	"prac-orm-transaction/infrastructure/table"
	"prac-orm-transaction/infrastructure"
	"fmt"
	"prac-orm-transaction/api/response"
)

func Auth() (fn func(http.Handler) http.Handler) { // 引数名を指定してるのでreturnのみでおｋ
	fn = func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Token") // Authというヘッダの値を取得する

			var dataExistsCheck table.UserInfo

			conn := infrastructure.NewMysqlRepository()
			conn.Client.First(&dataExistsCheck, "token=?", token)
			if dataExistsCheck.ID == "" {
				response.RespondError(w, http.StatusUnauthorized, fmt.Errorf("利用権限がありません"))
			}

			//if token != "admin" {         // adminという文字列か見る
			//	// エラーレスポンスを返す
			//	// この関数については後で書きます
			//	response.RespondError(w, http.StatusUnauthorized, fmt.Errorf("利用権限がありません"))
			//	return
			//}

			// 何も無ければ次のハンドラを実行する
			h.ServeHTTP(w, r)
		})
	}
	return
}