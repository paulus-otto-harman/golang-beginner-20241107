package router

import (
	"20241107/class/1/handler"
	"20241107/class/1/repository"
	"20241107/class/1/service"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
)

func initTemplate() (*repository.WebPageData, *template.Template) {
	tmpl, err := template.ParseGlob("view/*.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
		return nil, nil
	}

	return &repository.WebPageData{}, tmpl
}

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	handleOrder := handler.InitOrderHandler(*service.InitOrderService(*repository.InitOrderRepo(db)))
	handleWebTemplate := handler.InitWebPageHandler(*service.InitWebPageService(*repository.InitWebPageRepo(initTemplate())))

	//r.Use(middleware.Logger)
	//r.Use(middleware.BasicAuth)

	r.Get("/login", handleWebTemplate.Login)
	r.Get("/logout", handleWebTemplate.Logout)

	r.Get("/dashboard", handleWebTemplate.Dashboard)

	r.Route("/books", func(r chi.Router) {
		r.Get("/", handleWebTemplate.BookIndex)
		r.Get("/{id}", handleWebTemplate.BookShow)
		r.Get("/create", handleWebTemplate.BookCreate)
		r.Get("/{id}/edit", handleWebTemplate.BookEdit)

		r.Get("/{id}/discount/create", handleWebTemplate.BookDiscountCreate)
	})

	r.Route("/api", func(r chi.Router) {
		r.Post("/orders", handleOrder.Create)
	})

	r.Get("/style.css", handleWebTemplate.Static)

	return r
}
