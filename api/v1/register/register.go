package register

import (
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/mtwallet/api/middleware/auth/tokenmiddleware"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/register")
	{
		g.Use(tokenmiddleware.ApiAuth)
		g.GET("", register)
	}
}

func register(c *gin.Context) {
	uid, err := user.AddUser()
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to add user").SendD(c)
		logwrapper.Errorf("failed to add user, error: %s", err)

	} else {
		payload := RegisterPayload{
			Uid: uid,
		}
		httpo.NewSuccessResponse(200, "user registration successfull", payload).SendD(c)
	}
}
