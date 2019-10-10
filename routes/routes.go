package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/context"
	"github.com/zidni722/golang-restfull/bootstrap"
	"github.com/zidni722/golang-restfull/config"
	"github.com/zidni722/golang-restfull/web/controllers"
)

type Route struct {
	Config      *config.Configuration
	CorsHandler context.Handler
}

func NewRoute(config *config.Configuration) *Route {
	return &Route{
		Config: config,
		CorsHandler: cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			AllowedHeaders:   []string{"*"},
		}),
	}
}

func (r *Route) Configure(b *bootstrap.Bootstrapper) {
	b.Get("/", controllers.GetHomeHandler)
}
