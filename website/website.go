package website

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rapde/rap/website/zipfs"
)

// RAR with website static files data
var RAR = []byte{}

// Engine web server engine
type Engine struct {
	router *gin.Engine
}

// Run start web server
func (engine *Engine) Run(addr ...string) {
	engine.router.Run(addr...)
}

// New return web server engine
func New() *Engine {

	router := gin.Default()

	// ---------------------- Serving static files ----------------------

	fs, err := zipfs.New(RAR, &zipfs.Options{Prefix: "build/"})
	if err != nil {
		log.Fatal(err)
	}

	router.GET("/", func(c *gin.Context) {
		c.FileFromFS("./", fs)
	})

	router.NoRoute(func(c *gin.Context) {
		var (
			path   = c.Request.URL.Path
			_, err = fs.Stat("build" + path)
		)

		if err != os.ErrNotExist {
			c.FileFromFS(path, fs)
		} else {
			c.FileFromFS("./", fs)
		}
	})

	// ------------------------------------------------------------------

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	engine := &Engine{
		router: router,
	}

	return engine
}
