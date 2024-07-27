package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kotopesp/sos-kotopes/internal/controller/http/model/validator"
	"github.com/kotopesp/sos-kotopes/internal/core"
)

type Router struct {
	app             *fiber.App
	entityService   core.EntityService
	authService     core.AuthService
	commentsService core.CommentsService
	formValidator   validator.FormValidatorService
}

func NewRouter(
	app *fiber.App,
	entityService core.EntityService,
	formValidator validator.FormValidatorService,
	authService core.AuthService,
	commentsService core.CommentsService,
) {
	router := &Router{
		app:             app,
		entityService:   entityService,
		formValidator:   formValidator,
		authService:     authService,
		commentsService: commentsService,
	}
	router.initRequestMiddlewares()

	router.initRoutes()

	router.initResponseMiddlewares()
}

func (r *Router) initRoutes() {
	r.app.Get("/ping", r.ping)

	v1 := r.app.Group("/api/v1")

	// entities
	v1.Get("/entities", r.getEntities)
	v1.Get("/entities/:id", r.getEntityByID)

	//comments
	v1.Get("/posts/:post_id/comments", r.getCommentsByPostID)
	v1.Post("/posts/:post_id/comments", r.createComment)
	v1.Put("/posts/:post_id/comments/:comment_id", r.updateComment)
	v1.Delete("/posts/:post_id/comments/:comment_id", r.deleteComment)
	// e.g. protected resource
	v1.Get("/protected", r.protectedMiddleware(), r.protected)

	// auth
	v1.Post("/auth/login", r.loginBasic)
	v1.Post("/auth/signup", r.signup)
	v1.Post("/auth/token/refresh", r.refreshTokenMiddleware(), r.refresh)

	// auth vk
	v1.Get("/auth/login/vk", r.loginVK)
	v1.Get("/auth/login/vk/callback", r.callback)
}

// initRequestMiddlewares initializes all middlewares for http requests
func (r *Router) initRequestMiddlewares() {
	r.app.Use(logger.New())
}

// initResponseMiddlewares initializes all middlewares for http response
func (r *Router) initResponseMiddlewares() {}
