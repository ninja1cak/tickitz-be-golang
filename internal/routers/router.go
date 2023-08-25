package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	// router.Use(cors.New(config.CorsConfig))
	// router.Use(cors.Default())

	user(router, db)
	movie(router, db)
	auth(router, db)
	schedule(router, db)
	booking(router, db)
	return router
}
