//+build wireinject

package main

import (
	"database/sql"
	"go_playground/Go-000/Week04/myApp/internal/app/myApp/biz"
	"go_playground/Go-000/Week04/myApp/internal/app/myApp/service"

	"github.com/google/wire"
)

func initializeMyApp(db *sql.DB) service.Service {
	wire.Build(
		service.NewService,
		biz.NewBiz)
	return service.Service{}
}
