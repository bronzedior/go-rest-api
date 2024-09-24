package main

import (
	"errors"
	"fmt"
	"go-rest-api/handler"
	"go-rest-api/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r, err := routes(model.Database)
	if err != nil {
		fmt.Printf("Failed to make router: %v\n", err)
		os.Exit(1)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err = server.ListenAndServe(); err != nil {
		fmt.Printf("Failed to run server: %v\n", err)
		os.Exit(1)
	}
}

func routes(db map[string]model.Car) (http.Handler, error) {
	if db == nil {
		return nil, errors.New("routes: no connection to database")
	}

	r := gin.Default()

	public := r.Group("/api/v1")
	{
		public.GET("/cars", handler.ListProduct(db))
		public.GET("/cars/:id", handler.GetProduct(db))
		public.POST("/cars", handler.CreateProduct(db))
		public.PUT("/cars/:id", handler.UpdateProduct(db))
		public.DELETE("/cars/:id", handler.DeleteProduct(db))
	}

	return r, nil
}
