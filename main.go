package main

import (
	"doovvvblog/model"
	"doovvvblog/routes"
	"doovvvblog/utils"
)

func main() {
	utils.Init()
	model.InitDB()
	routes.InitRouter()
}
