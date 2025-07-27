package routes

import (
	"github.com/hossainabid/go-ims/consts"
	"github.com/hossainabid/go-ims/controllers"
	m "github.com/hossainabid/go-ims/middlewares"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	echo             *echo.Echo
	productCtrl      *controllers.ProductController
	userCtrl         *controllers.UserController
	stockHistoryCtrl *controllers.StockHistoryController
	authCtrl         *controllers.AuthController
	authMiddleware   *m.AuthMiddleware
}

func New(e *echo.Echo, productCtrl *controllers.ProductController, userCtrl *controllers.UserController, stockHistoryCtrl *controllers.StockHistoryController, authCtrl *controllers.AuthController, authMiddleware *m.AuthMiddleware) *Routes {
	return &Routes{
		echo:             e,
		productCtrl:      productCtrl,
		userCtrl:         userCtrl,
		stockHistoryCtrl: stockHistoryCtrl,
		authCtrl:         authCtrl,
		authMiddleware:   authMiddleware,
	}
}

func (r *Routes) Init() {
	e := r.echo
	m.Init(e)
	// APM routes
	e.GET("/metrics", echoprometheus.NewHandler())

	g := e.Group("/v1")

	products := g.Group("/products")
	products.POST("", r.productCtrl.CreateProduct, r.authMiddleware.Authenticate(consts.PermissionProductCreate))
	products.GET("", r.productCtrl.ListProducts, r.authMiddleware.Authenticate(consts.PermissionProductList))
	products.GET("/:id", r.productCtrl.ReadProductByID, r.authMiddleware.Authenticate(consts.PermissionProductFetch))
	products.PUT("/:id", r.productCtrl.UpdateProduct, r.authMiddleware.Authenticate(consts.PermissionProductUpdate))
	products.DELETE("/:id", r.productCtrl.DeleteProduct, r.authMiddleware.Authenticate(consts.PermissionProductDelete))

	users := g.Group("/users")
	users.POST("/signup", r.userCtrl.Signup)
	users.GET("/profile", r.userCtrl.Profile, r.authMiddleware.Authenticate(""))
	users.POST("", r.userCtrl.CreateUser, r.authMiddleware.Authenticate(consts.PermissionUserCreate))
	users.GET("", r.userCtrl.ListUsers, r.authMiddleware.Authenticate(consts.PermissionUserList))
	users.GET("/:id", r.userCtrl.ReadUser, r.authMiddleware.Authenticate(consts.PermissionUserFetch))
	users.PUT("/:id", r.userCtrl.UpdateUser, r.authMiddleware.Authenticate(consts.PermissionUserUpdate))
	users.DELETE("/:id", r.userCtrl.DeleteUser, r.authMiddleware.Authenticate(consts.PermissionUserDelete))

	auth := g.Group("/auth")
	auth.POST("/login", r.authCtrl.Login)
	auth.POST("/logout", r.authCtrl.Logout, r.authMiddleware.Authenticate(""))

	stockHistories := g.Group("/stock-histories")
	stockHistories.POST("", r.stockHistoryCtrl.RecordStockHistory, r.authMiddleware.Authenticate(consts.PermissionStockHistoryRecord))
	stockHistories.GET("", r.stockHistoryCtrl.ListStockHistories, r.authMiddleware.Authenticate(consts.PermissionStockHistoryList))

}
