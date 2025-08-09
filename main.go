package main

import "dragonball-api/routers"

func main() {
	router := routers.SetupRouter()
	router.Run()
}