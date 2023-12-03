package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suraj1294/go-fiber-pg-auth/cmd/api/handler"
	"github.com/suraj1294/go-fiber-pg-auth/cmd/api/response"
)

func (app *Application) RegisterRoutes() *fiber.App {

	router := fiber.New()

	api := fiber.New()

	router.Mount("/api", api)

	// Or extend your config for customization
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*",
	// 	AllowHeaders:     "Accept, Content-Type, X-CSRF-Token, Authorization",
	// 	AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
	// 	AllowCredentials: true,
	// }))

	router.Static("/", "./client/dist")

	appStore := handler.GetAppHandler(*app.Auth, *app.AppDB)

	authHandler := handler.AuthHandler{AppStore: *appStore}
	moviesHandler := handler.MoviesHandler{AppStore: *appStore}
	usersHandler := handler.UsersHandler{AppStore: *appStore}
	middlewareHandler := handler.MiddlewareHandler{AppStore: *appStore}

	// router.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(response.NewSuccessResponse("ok"))
	// })
	//root := router.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(response.NewSuccessResponse("ok"))
	})

	movies := api.Group("/movies")
	// api/movies
	movies.Get("/", middlewareHandler.CustomMiddleware, moviesHandler.AllMovies)
	// api/movies/:id
	movies.Get("/:id", moviesHandler.Movie)
	// api/users
	users := api.Group("/users")
	users.Get("/", usersHandler.MockUsers)
	auth := api.Group("/auth")
	// api/auth/login
	auth.Post("/login", authHandler.Authenticate)
	auth.Get("/refresh", authHandler.RefreshToken)
	auth.Get("/logout", authHandler.Logout)
	auth.Get("/me", authHandler.AuthProfile)

	api.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "page not found"})
	})

	// router.Use(func(c *fiber.Ctx) error {
	// 	return c.Redirect("/")
	// })

	return router

}
