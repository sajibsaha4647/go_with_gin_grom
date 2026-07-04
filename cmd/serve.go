package cmd

import (
	"ecommerce/config"
	"ecommerce/infra"
	"ecommerce/rest"
	"fmt"
	"log"
)

func ServerGo() {

	cfg, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	db, err := infra.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// 3. (Optional but Recommended) Assign it to your global package variable
	// if you plan to access it via infra.DB across your app
	infra.DB = db

	log.Println("Database connection successfully established!")

	fmt.Println(cfg.Port)
	fmt.Println(cfg.JWTSecret)

	

	rest.NewServer(nil, nil).Start(cfg)

}
