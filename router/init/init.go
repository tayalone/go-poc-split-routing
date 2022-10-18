package init

import (
	"go-poc-split-routing/router"
	"go-poc-split-routing/router/fiber"
	"go-poc-split-routing/router/gin"
)

/*Init Reouter Router Instant */
func Init(rType string, conf router.Config) router.Route {
	switch rType {
	case "GIN":
		return gin.NewMyRouter(conf)
	case "FIBER":
		return fiber.NewFiberRouter(conf)
	default:
		return gin.NewMyRouter(conf)
	}
}
