package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
	"os"
	app2 "ppo/internal/app"
	"ppo/internal/config"
	iLogger "ppo/pkg/logger"
	"ppo/pkg/storage/postgres"
	"ppo/web"
	"ppo/web/v1"
	v2 "ppo/web/v2"
)

var tokenAuth *jwtauth.JWTAuth

// @title Сервис поиска рецептов салатов
// @version 1.0
// @description Сервис призван помочь с поиском рецептов.

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8081
// @BasePath /api/v2
// @query.collection.format multi

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	logFile, err := os.OpenFile(cfg.Logger.File, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(logFile)
	logger := iLogger.NewLogger(cfg.Logger.Level, logFile)

	tokenAuth = jwtauth.New("HS256", []byte(cfg.Jwt.Key), nil)

	fmt.Println("connecting to database...")
	pool, err := postgres.NewConn(context.Background(), cfg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("successfully connected to database!")

	app := app2.NewApp(pool, cfg, logger)
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mux.Use(middleware.Logger)

	mux.Route("/api", func(r chi.Router) {
		r.Route("/v2", func(r chi.Router) {
			r.Post("/login", v2.LoginHandler(app))
			r.Post("/register", v2.RegisterHandler(app))

			r.Route("/users", func(r chi.Router) {
				r.Use(jwtauth.Verifier(tokenAuth))
				r.Use(jwtauth.Authenticator(tokenAuth))
				r.Use(web.ValidateUserRoleJWT)

				r.Get("/{id}", v2.GetUser(app))

				r.Route("/me", func(r chi.Router) {
					r.Get("/salads", v2.GetUserSaladsByFlag(app))
				})
			})

			r.Route("/salads", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))

					r.Post("/", v2.CreateSalad(app))
				})

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(web.Authenticator(tokenAuth))

					r.Get("/", v2.GetSaladsFunc(app))
				})

				r.Route("/{id}", func(r chi.Router) {
					// salad info
					r.Get("/", v2.GetSaladById(app))

					// types
					r.Get("/types", v2.GetSaladTypes(app))

					// comments
					r.Get("/comments", v2.GetCommentsFunc(app))

					r.Group(func(r chi.Router) {
						r.Use(jwtauth.Verifier(tokenAuth))
						r.Use(jwtauth.Authenticator(tokenAuth))

						// salads
						r.Patch("/", v2.UpdateSalad(app))
						r.Delete("/", v2.DeleteSalad(app))

						// ingredients
						r.Post("/ingredients", v2.LinkIngredientToSalad(app))
						r.Delete("/ingredients", v2.UnlinkIngredientFromSalad(app))

						// types
						r.Post("/types", v2.LinkTypeToSalad(app))
						r.Delete("/types", v2.UnlinkTypeFromSalad(app))
					})
				})
			})

			r.Route("/recipes", func(r chi.Router) {
				r.Get("/{id}", v2.GetSaladRecipe(app))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/", v2.CreateRecipe(app))
					r.Patch("/{id}", v2.UpdateRecipe(app))
				})
			})

			r.Route("/recipeSteps", func(r chi.Router) {
				r.Get("/{id}", v2.GetRecipeSteps(app))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/", v2.CreateRecipeStep(app))
					r.Patch("/{id}", v2.UpdateRecipeStep(app))
					r.Delete("/{id}", v2.DeleteRecipeStep(app))
				})
			})

			r.Route("/ingredients", func(r chi.Router) {
				r.Get("/{id}", v2.GetRecipeIngredients(app))
				r.Get("/", v2.GetIngredientsByPage(app))
			})

			r.Route("/types", func(r chi.Router) {
				r.Get("/", v2.GetSaladTypesByPage(app))
			})

			r.Route("/measurements", func(r chi.Router) {
				r.Get("/{id}", v2.GetMeasurementByRecipe(app))
				r.Get("/", v2.GetAllMeasurements(app))
			})

			r.Route("/comments", func(r chi.Router) {
				//r.Get("/{id}", v2.GetCommentById(app))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/", v2.CreateComment(app))
					r.Patch("/{id}", v2.UpdateComment(app))
					r.Delete("/{id}", v2.DeleteComment(app))
				})
			})

		})

		r.Route("/v1", func(r chi.Router) {
			mux.Post("/login", v1.LoginHandler(app))
			mux.Post("/signup", v1.RegisterHandler(app))

			mux.Route("/users", func(r chi.Router) {
				r.Get("/{id}", v1.GetUser(app))

			})

			mux.Route("/salads", func(r chi.Router) {
				r.Get("/", v1.GetSalads(app))
				r.Get("/{id}/rating", v1.GetSaladRating(app))
				r.Get("/{id}", v1.GetSaladById(app))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/create", v1.CreateSalad(app))
					r.Patch("/{id}/update", v1.UpdateSalad(app))
					r.Get("/user", v1.GetUserSalads(app))
					r.Get("/userRated", v1.GetUserRatedSalads(app))
					r.Delete("/{id}/delete", v1.DeleteSalad(app))
				})

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateAdminRoleJWT)

					r.Get("/status", v1.GetSaladsWithStatus(app))
				})
			})

			mux.Route("/recipe", func(r chi.Router) {
				r.Get("/{id}", v1.GetSaladRecipe(app))
				r.Get("/{id}/rating", v1.GetSaladRating(app))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/create", v1.CreateRecipe(app))
					r.Patch("/{id}/update", v1.UpdateRecipe(app))
				})
			})

			mux.Route("/recipeSteps", func(r chi.Router) {
				r.Get("/{id}", v1.GetRecipeSteps(app))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/create", v1.CreateRecipeStep(app))
					r.Patch("/{id}/update", v1.UpdateRecipeStep(app))
					r.Delete("/{id}/delete", v1.DeleteRecipeStep(app))
				})
			})

			mux.Route("/ingredients", func(r chi.Router) {
				r.Get("/{id}", v1.GetRecipeIngredients(app))
				r.Get("/", v1.GetIngredientsByPage(app))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/link", v1.LinkIngredientToSalad(app))
					r.Patch("/unlink", v1.UnlinkIngredientFromSalad(app))
				})
			})

			mux.Route("/types", func(r chi.Router) {
				r.Get("/{id}", v1.GetSaladTypes(app))
				r.Get("/", v1.GetSaladTypesByPage(app))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/link", v1.LinkTypeToSalad(app))
					r.Patch("/unlink", v1.UnlinkTypeFromSalad(app))
				})
			})

			mux.Route("/ingredientTypes", func(r chi.Router) {
				r.Get("/{id}", v1.GetIngredientType(app))
			})

			mux.Route("/measurements", func(r chi.Router) {
				r.Get("/", v1.GetMeasurementByRecipe(app))
				r.Get("/all", v1.GetAllMeasurements(app))
			})

			mux.Route("/comments", func(r chi.Router) {
				r.Get("/", v1.GetCommentsBySalad(app))
				r.Get("/userSalad", v1.GetUserComment(app))
				r.Get("/{id}", v1.GetCommentById(app))

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(tokenAuth))
					r.Use(jwtauth.Authenticator(tokenAuth))
					r.Use(web.ValidateUserRoleJWT)

					r.Post("/create", v1.CreateComment(app))
					r.Patch("/{id}/update", v1.UpdateComment(app))
					r.Delete("/{id}/delete", v1.DeleteComment(app))
				})
			})
		})
	})

	fmt.Println("server was started")
	http.ListenAndServe(":8081", mux)
}
