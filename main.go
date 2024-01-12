package main

import (
	"github.com/gabriels6/finance-integrator/utils"
)

func main() {
	// Routes() - imported from routes.go
	router := Routes()
	router.Run(":"+utils.GetEnv("PORT"))
}

