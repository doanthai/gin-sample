package main

import (
	"fmt"
	"gin-sample/config"
	"gin-sample/src/router"
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"time"
)

func main() {
	cfg := config.GetConfig()

	r := initRouterDefault()

	router.InitRouter(r)

	s := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Listening and serving HTTP on :", cfg.App.Port)
	s.ListenAndServe()
}

func init() {
	config.LoadConfig("local", "./config")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())
	cfg := config.GetConfig()

	// Show info service
	fmt.Println("CONFIG: ")
	fmt.Println("- Port: " + cfg.App.Port)
	fmt.Println("- Num CPU: " + fmt.Sprint(runtime.NumCPU()))

}

func initRouterDefault() *gin.Engine {

	router := gin.Default()
	router.Use(sentry.Recovery(raven.DefaultClient, false))

	return router
}
