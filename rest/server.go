package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct { //all handler dependency

	cnf *config.Config

	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer( //dependency inject
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{ //pass kore dilam
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()

	manager.Use(
		middleware.CorsWithPreFlight,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.With(
		manager.GetGlobalMiddlewares()...,
	)(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort)

	fmt.Println("Server running on:", addr)

	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
