package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"prac-orm-transaction/api/response"
	"time"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

// Router ルーティング設定
func (s *Server) Routing() {
	s.Router.Use(middleware.Timeout(60 * time.Second))

	s.Router.Route("/api", func(api chi.Router) {
		api.Use(Auth("db connection"))                 // /api/*で必ず通るミドルウェア
		api.Route("/users", func(members chi.Router) { // /api/members/* でグループ化
			//members.Get("/{id}") // /api/members/1  などで受け取るハンドラ
			members.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("hello world"))
			}) // /api/members で受け取るハンドラ
		})
	})
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	// Authする何かのエンドポイントという想定
	//s.router.Route("/api/auth", func(auth chi.Router) {
	//	auth.Get("/login")
	//})
}

func Auth(db string) (fn func(http.Handler) http.Handler) { // 引数名を指定してるのでreturnのみでおｋ
	fn = func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Auth") // Authというヘッダの値を取得する
			if token != "admin" {         // adminという文字列か見る
				// エラーレスポンスを返す
				// この関数については後で書きます
				response.RespondError(w, http.StatusUnauthorized, fmt.Errorf("利用権限がありません"))
				return
			}
			// 何も無ければ次のハンドラを実行する
			h.ServeHTTP(w, r)
		})
	}
	return
}
