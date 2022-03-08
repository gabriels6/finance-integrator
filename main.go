package main

func main() {
	// Routes() - imported from routes.go
	router := Routes()
	router.Run("localhost:3000")
}

