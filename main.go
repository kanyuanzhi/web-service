package main

import "github.com/kanyuanzhi/web-service/internal/routers"

func main() {
	r := routers.NewRouter()
	r.Run()
}
