package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	middleware2 "prac-orm-transaction/api/middleware"
	"prac-orm-transaction/api/controller"
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
func (s *Server) Routing(uh controller.UserHandler) {
	s.Router.Use(middleware.Timeout(60 * time.Second))
	s.Router.Use(middleware.Logger)
	s.Router.Route("/sign", func(api chi.Router) {
		api.Use(middleware2.Auth("db connection")) // /api/*で必ず通るミドルウェア
		api.Route("/up", func(signup chi.Router) {
			signup.Post("/",uh.CreateUserAccount())
		})
		api.Route("/in", func(signin chi.Router) {
			signin.Post("/",func(w http.ResponseWriter, r *http.Request) {
						w.Write([]byte("hello world"))
					})
		})
	})
	//s.Router.Route("/api", func(api chi.Router) {
	//	api.Use(Auth("db connection"))                 // /api/*で必ず通るミドルウェア
	//	api.Route("/users", func(members chi.Router) { // /api/members/* でグループ化
	//		members.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//			w.Write([]byte("hello world"))
	//		}) // /api/members で受け取るハンドラ
	//	})
	//})
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	// Authする何かのエンドポイントという想定
	//s.router.Route("/api/auth", func(auth chi.Router) {
	//	auth.Get("/login")
	//})
}
