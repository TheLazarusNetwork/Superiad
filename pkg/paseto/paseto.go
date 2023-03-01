package paseto

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	gopaseto "aidanwoods.dev/go-paseto"
	"github.com/MyriadFlow/superiad/pkg/auth"
	"github.com/MyriadFlow/superiad/pkg/claims"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	ErrAuthHeaderMissing = errors.New("authorization header is required")
)

func PASETO(c *gin.Context) {
	var headers GenericAuthHeaders
	err := c.BindHeader(&headers)
	if err != nil {

		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to bind")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if headers.Authorization == "" {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Autherisation header is missing")
		c.Abort()
		return
	}
	token := headers.Authorization
	splitToken := strings.Split(token, "Bearer ")
	pasetoToken := splitToken[1]
	parser := gopaseto.NewParser()
	parser.AddRule(gopaseto.NotExpired())
	publickey := auth.Getpublickey()
	parsedToken, err := parser.ParseV4Public(publickey, pasetoToken, nil)
	if err != nil {

		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to bindfailed to scan claims for paseto token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		jsonvalue := parsedToken.ClaimsJSON()
		ClaimsValue := claims.CustomClaims{}
		json.Unmarshal(jsonvalue, &ClaimsValue)
		c.Set("Email", ClaimsValue.Email)
		c.Set("UUID", ClaimsValue.UUID)
		c.Next()
	}

}
