package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {

	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := db.MigrateDB(dbCon); err != nil {
		fmt.Println("Database migration failed:", err)
		os.Exit(1)
	}

	fmt.Println("Database migration completed successfully")

	middlewares := middleware.NewMeddleWares(cnf)

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	productHandler := product.NewHandler(middlewares, productRepo)

	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)

	server.Start()

}
