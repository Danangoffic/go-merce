package main

import (
	"github.com/Danangoffic/go-merce/app/config"
	"github.com/Danangoffic/go-merce/app/routes"
)

func main() {
	db := config.InitDB()
	if db.Error != nil {
		panic("Db failed to load caused : " + db.Error.Error())
	}

	router := config.SetupRouter()
	routes.LoadRouters(router, db)

	router.Run(":8000")
}
