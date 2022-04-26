package main

import (
	"context"
	"net/http"

	"github.com/kmhalpin/todoapp/common/env"
	errorCommon "github.com/kmhalpin/todoapp/common/error"
	passCommon "github.com/kmhalpin/todoapp/common/password"
	dbCommon "github.com/kmhalpin/todoapp/common/sqlite"
	uuidCommon "github.com/kmhalpin/todoapp/common/uuid"
	"github.com/kmhalpin/todoapp/internal/model/user"
	uRepo "github.com/kmhalpin/todoapp/internal/repository/user/sqlite"
)

func init() {
	ctx := context.Background()

	cfg := env.LoadConfig()
	uuid := uuidCommon.NewUUIDGenerator()
	db := dbCommon.NewSqlite()
	pm := passCommon.NewPasswordHashManager()
	ur := uRepo.NewSqliteUserRepository(db, uuid)

	pass, err := pm.HashPassword(cfg.RootPass)
	if err != nil {
		panic(err)
	}

	_, err = ur.GetIDByUsername(ctx, user.ROOT_USERNAME)
	if err != nil {
		if cerr, ok := err.(*errorCommon.ClientError); ok {
			if cerr.StatusCode == http.StatusNotFound {
				if _, err := ur.InsertUser(ctx, user.User{
					Username:   user.ROOT_USERNAME,
					Password:   pass,
					Permission: user.Write,
				}); err != nil {
					panic(err)
				}
			}
		} else {
			panic(err)
		}
	}
}
