package main

import (
	"fmt"
	"net/http"

	"go-poc-split-routing/router"

	RouteInitor "go-poc-split-routing/router/init"
)

func iSayPing(c router.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}

func myCustomMdw(c router.Context) {
	fmt.Println("I println from myCustomMdw")
	c.Next()
}

func main() {
	// Can Swith Gin, Fiber or ... web framework
	myRouter := RouteInitor.Init("GIN", router.Config{Port: 3000})
	myRouter.GET("/ping", myCustomMdw, iSayPing)
	myRouter.GET("/fiber", func(c router.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Show Value from /fiber",
		})
	})

	v1 := myRouter.Group("/v1")

	v1.GET("/ping", myCustomMdw, iSayPing)

	myRouter.Start()
}
