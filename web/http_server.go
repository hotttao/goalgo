package web

import (
	"github.com/gin-gonic/gin"
)

func AuthHandler(word string) func(c *gin.Context) {
	gin.Recovery()
	return func(c *gin.Context) {
		c.Next()
		c.JSON(200, gin.H{"visit": word})
	}
}

func StartWebServer() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	// authorized.Use(gin.AuthRequired())
	{
		authorized.GET("/login", AuthHandler("logini"))
		authorized.GET("/submit", AuthHandler("submit"))
		authorized.GET("/read", AuthHandler("read"))

		// nested group
		testing := authorized.Group("testing")
		// visit 0.0.0.0:8080/testing/analytics
		testing.GET("/analytics", AuthHandler("testing/analytics"))
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
