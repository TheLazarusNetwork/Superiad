package tokenmiddleware

import (
	"errors"
	"net/http"

	"github.com/TheLazarusNetwork/mtwallet/config/envconfig"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/mtwallet/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
)

var (
	ErrAuthHeaderMissing = errors.New("authorization header is required")
)

func ApiAuth(c *gin.Context) {
	var headers GenericAuthHeaders
	err := c.BindHeader(&headers)
	token := headers.Authorization

	if err != nil {
		logValidationFailed(token, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if headers.Authorization == "" {
		logValidationFailed(token, err)
		httphelper.ErrResponse(c, http.StatusBadRequest, ErrAuthHeaderMissing.Error())
		c.Abort()
		return
	}

	if token != envconfig.EnvVars.TOKEN {
		logValidationFailed(token, err)
		httphelper.ErrResponse(c, http.StatusForbidden, "token is invalid")
		c.Abort()
		return
	}
}

func logValidationFailed(token string, err error) {
	logwrapper.Warnf("validation failed with token %v, error: %v", token, err)
}
