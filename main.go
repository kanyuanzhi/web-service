package main

import (
	"fmt"
	"github.com/kanyuanzhi/web-service/global"
	"github.com/kanyuanzhi/web-service/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(fmt.Sprintf("%s:%d", global.Object.Host, global.Object.Port))
}
