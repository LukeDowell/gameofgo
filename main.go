package main

import (
	"dowell.dev/gameofgo/api"
	"dowell.dev/gameofgo/repo"
)

func main() {
	r := api.SetupRouter(repo.New())
	_ = r.Run(":8080")
}
