package cmd

import (
	"ecommerce/config"
	"fmt"
)

func ServerGo() {

	cfg, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.Port)
	fmt.Println(cfg.JWTSecret)

}
