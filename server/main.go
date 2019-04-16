package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab.com/evzpav/documents-crud/server/handlers"
)

var dbSession *handlers.MongoSession
var serverUp time.Time

func init() {
	serverUp = time.Now()
	dbSession = handlers.NewMongo(serverUp)
	dbSession.CreateCollection("documents-crud", "documents")
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	log.Printf("mongo session %+v", dbSession)

	e.GET("/", Index)
	e.GET("/documents", dbSession.GetDocuments)
	e.POST("/document", dbSession.CreateDocument)
	e.PUT("/document/:id", dbSession.UpdateDocument)
	e.DELETE("/document/:id", dbSession.DeleteDocument)
	e.GET("/status", dbSession.ServerStatus)

	s := serverConfig()
	e.Logger.Fatal(e.StartServer(s))
}

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Documents CRUD!")
}

func serverConfig() *http.Server {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	return &http.Server{
		Addr:         port,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
