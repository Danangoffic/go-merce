package main

import "github.com/Danangoffic/go-merce/app/config"

func main() {
	db, router := config.InitAPP()
	if db.Error != nil {
		panic("Db failed to load caused : " + db.Error.Error())
	}

	router.Run(":8000")
}
