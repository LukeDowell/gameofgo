package main

import (
	"dowell.dev/gameofgo/api"
)

func main() {
	r := api.SetupRouter()
	_ = r.Run(":8080")
}
