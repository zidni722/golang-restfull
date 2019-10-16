package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/zidni722/golang-restfull/app/web/controllers"
	"github.com/zidni722/golang-restfull/bootstrap"
	"github.com/zidni722/golang-restfull/config"
	"github.com/zidni722/pawoon-product/app/repositories/impl"
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

	userRepository := impl.NewProductRepositoryImpl()
	// example for
	v1 := b.Party("/user", r.CorsHandler).AllowMethods(iris.MethodOptions)
	{
		userController := controllers.NewUserController(r.Config.Database.DB, userRepository)
		v1.Post("/create", userController.CreateUserHandler)
		// products := v1.Party("/product")
		// {
		// 	productController := controllers.NewProductController(r.Config.Database.DB, productRepository)
		// 	products.Get("/{productUuid:string}", productController.GetProductByUuid)
		// }
	}
}
