package routers

import (
	"github.com/gin-gonic/gin"
	"phone_email/actions"
)

var r *gin.Engine

func Load() *gin.Engine {
	r.PUT("number/:call_id/:number/:type", actions.NumberPut)

	return r
}

func init() {
	r = gin.Default()
}
