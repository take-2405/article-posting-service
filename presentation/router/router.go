package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"prac-orm-transaction/presentation/controller"
	middleware2 "prac-orm-transaction/presentation/middleware"
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
func (s *Server) Routing(uh controller.UserHandler, ah controller.ArticleHandler) {
	s.Router.Use(middleware.Timeout(60 * time.Second))
	s.Router.Use(middleware.Logger)
	s.Router.Route("/sign", func(api chi.Router) {
		api.Route("/up", func(signup chi.Router) {
			signup.Post("/", uh.CreateUserAccount())
		})
		api.Route("/in", func(signin chi.Router) {
			signin.Post("/", uh.SignIn())
		})
	})

	s.Router.Route("/article", func(api chi.Router) {
		api.Use(middleware2.Auth()) // /api/*で必ず通るミドルウェア
		api.Route("/create", func(create chi.Router) {
			create.Post("/", ah.CreateArticle())
		})
		api.Route("/fix", func(fix chi.Router) {
			fix.Put("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("fix"))
			})
		})
		api.Route("/search", func(search chi.Router) {
			search.Put("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("search"))
			})
		})
	})

	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
}
