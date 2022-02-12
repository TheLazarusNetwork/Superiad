package register

import (
	"github.com/TheLazarusNetwork/mtwallet/api/middleware/auth/tokenmiddleware"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/httphelper"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/register")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.GET("/", register)
	}
}

func register(c *gin.Context) {
	uid, err := user.AddUser()
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to add user, error: %v", err.Error())
	} else {
		payload := RegisterPayload{
			Uid: uid,
		}
		httphelper.SuccessResponse(c, "user registration successfull", payload)
	}
}
