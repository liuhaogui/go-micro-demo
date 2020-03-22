package base

import (
	"microservice/user-srv/base/config"
	"microservice/user-srv/base/db"
)

func Init()  {
	config.Init()
	db.Init()
}
