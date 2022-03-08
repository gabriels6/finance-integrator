package main

func main() {
	// Routes() - imported from routes.go
	router := Routes()
	router.Run(":"+GetEnv("PORT"))
}

