package main

import (
	db "Assignment-2/database"
	"Assignment-2/routers"
)

func main() {
	db.Init()
	routers.ServerOn().Run(":8080")
}
