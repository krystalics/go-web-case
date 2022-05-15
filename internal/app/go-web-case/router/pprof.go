package router

import (
	"github.com/gin-gonic/gin"
	"net/http/pprof"
)

func RegisterPprof(e *gin.Engine) {
	//maybe use some special middleware
	routers := []URLHandlerMap{
		{"GET", "/debug/pprof/", GinHandler(pprof.Index)},
		{"GET", "/debug/pprof/heap", GinHandler(pprof.Handler("heap").ServeHTTP)},
		{"GET", "/debug/pprof/goroutine", GinHandler(pprof.Handler("goroutine").ServeHTTP)},
		{"GET", "/debug/pprof/allocs", GinHandler(pprof.Handler("allocs").ServeHTTP)},
		{"GET", "/debug/pprof/block", GinHandler(pprof.Handler("block").ServeHTTP)},
		{"GET", "/debug/pprof/threadcreate", GinHandler(pprof.Handler("threadcreate").ServeHTTP)},
		{"GET", "/debug/pprof/cmdline", GinHandler(pprof.Cmdline)},
		{"GET", "/debug/pprof/profile", GinHandler(pprof.Profile)},
		{"GET", "/debug/pprof/symbol", GinHandler(pprof.Symbol)},
		{"POST", "/debug/pprof/symbol", GinHandler(pprof.Symbol)},
		{"GET", "/debug/pprof/trace", GinHandler(pprof.Trace)},
		{"GET", "/debug/pprof/mutex", GinHandler(pprof.Handler("mutex").ServeHTTP)},
	}

	RouterFill(&e.RouterGroup, &routers)
}
