package http

import (
	"gitflic.ru/spbu-se/sos-kotopes/internal/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	app             *fiber.App
	entityService   core.EntityService
	authService     interface{}
	commentsService core.CommentsService
}

func NewRouter(
	app *fiber.App,
	entityService core.EntityService,
	authService interface{},
	commentsService core.CommentsService,
) {
	router := &Router{
		app:             app,
		entityService:   entityService,
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
	v1.Post("/comments", r.createComment)
	v1.Put("/posts/:post_id/comments/:comment_id", r.updateComment)
	v1.Delete("/comments/:comment_id", r.deleteComment)
}

// initRequestMiddlewares initializes all middlewares for http requests
func (r *Router) initRequestMiddlewares() {
	r.app.Use(logger.New())
}

// initResponseMiddlewares initializes all middlewares for http response
func (r *Router) initResponseMiddlewares() {}
