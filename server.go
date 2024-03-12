package main

import (
	"backend-th-budget-tracker/api"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	print("hello world")
	router := gin.New()
	// router.ge
	api.RegisterEndpoints(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
