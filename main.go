package main

import (
	"ShopAgent/router"
)

func main() {
	r := router.Routers()
	r.Run(":8000")
}
