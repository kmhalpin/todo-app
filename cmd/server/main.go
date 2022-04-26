package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/kmhalpin/todoapp/common/env"
	httpCommon "github.com/kmhalpin/todoapp/common/http"
	jwtCommon "github.com/kmhalpin/todoapp/common/jwt"
	passCommon "github.com/kmhalpin/todoapp/common/password"
	dbCommon "github.com/kmhalpin/todoapp/common/sqlite"
	uuidCommon "github.com/kmhalpin/todoapp/common/uuid"
	aDelivery "github.com/kmhalpin/todoapp/internal/delivery/auth/http"
	tDelivery "github.com/kmhalpin/todoapp/internal/delivery/todo/http"
	uDelivery "github.com/kmhalpin/todoapp/internal/delivery/user/http"
	tRepo "github.com/kmhalpin/todoapp/internal/repository/todo/sqlite"
	uRepo "github.com/kmhalpin/todoapp/internal/repository/user/sqlite"
	aUCase "github.com/kmhalpin/todoapp/internal/usecase/auth"
	tUCase "github.com/kmhalpin/todoapp/internal/usecase/todo"
	uUCase "github.com/kmhalpin/todoapp/internal/usecase/user"
)

func main() {
	cfg := env.LoadConfig()
	uuid := uuidCommon.NewUUIDGenerator()
	db := dbCommon.NewSqlite()
	h := httpCommon.NewHTTPServer()
	jwt := jwtCommon.NewJWTManager(cfg.AccessTokenKey)
	pm := passCommon.NewPasswordHashManager()

	staticServer := static.Serve("/", static.LocalFile("./web/build", true))
	h.Router.Use(staticServer)
	h.Router.Use(httpCommon.MiddlewareErrorHandler())
	h.Router.RedirectTrailingSlash = true
	root := h.Router.Group("/api")

	ur := uRepo.NewSqliteUserRepository(db, uuid)
	uc := uUCase.NewUserUsecase(ur, pm)
	uDelivery.NewHTTPUserDelivery(root.Group("/user"), uc, jwt)

	tr := tRepo.NewSqliteTodoRepository(db, uuid)
	tc := tUCase.NewTodoUsecase(tr, uc)
	tDelivery.NewHTTPTodoDelivery(root.Group("/todo"), tc, jwt)

	ac := aUCase.NewAuthUsecase(ur, pm, jwt)
	aDelivery.NewHTTPAuthDelivery(root.Group("/auth"), ac)

	h.Router.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet &&
			!strings.ContainsRune(c.Request.URL.Path, '.') &&
			!strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Request.URL.Path = "/"
			staticServer(c)
		}
	})

	log.Fatal(h.Router.Run(fmt.Sprintf(":%d", cfg.Port)))
}
