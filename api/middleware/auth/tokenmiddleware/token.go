package tokenmiddleware

import (
	"errors"
	"net/http"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/TheLazarusNetwork/superiad/config/envconfig"

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
		c.Abort()
		return
	}
	if headers.Authorization == "" {
		logValidationFailed(token, err)
		httpo.NewErrorResponse(http.StatusBadRequest, ErrAuthHeaderMissing.Error()).SendD(c)
		c.Abort()
		return
	}

	if token != envconfig.EnvVars.TOKEN {
		logValidationFailed(token, err)
		httpo.NewErrorResponse(http.StatusForbidden, "token is invalid").SendD(c)
		c.Abort()
		return
	}
}

func logValidationFailed(token string, err error) {
	logo.Warnf("validation failed with token %v, error: %s", token, err)
}
