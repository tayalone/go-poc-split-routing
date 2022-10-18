package router

/*Context is Behavior of Route Context In Application*/
type Context interface {
	Next()
	JSON(int, interface{})
}

/*Route is Behavior of Route Method In Application*/
type Route interface {
	Start()
	GET(path string, handlers ...func(Context))
	Group(path string, handlers ...func(Context)) RouteGouping
}

/*RouteGouping is Behavior of Route Method In Application*/
type RouteGouping interface {
	GET(path string, handlers ...func(Context))
}

/*Config is Stric of  Configuraiont Of Http Router*/
type Config struct {
	Port int
}
