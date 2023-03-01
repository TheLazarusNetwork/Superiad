package auth

import (
	"context"
	"net/http"

	"github.com/MyriadFlow/superiad/config/envconfig"
	"github.com/MyriadFlow/superiad/models/user"
	"github.com/MyriadFlow/superiad/pkg/auth"
	"github.com/MyriadFlow/superiad/pkg/claims"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
	log "github.com/sirupsen/logrus"
)

var supabaseClient *supa.Client

type GenericAuthHeaders struct {
	Authorization string
}

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/auth")
	{
		g.POST("/", authReq)
	}
}
func Init() {
	supabaseUrl := envconfig.EnvVars.SUPABASEURL
	supabaseKey := envconfig.EnvVars.SUPABASEKEY
	supabaseClient = supa.CreateClient(supabaseUrl, supabaseKey)
}

func authReq(c *gin.Context) {
	var headers GenericAuthHeaders
	var claimsvalue claims.CustomClaims
	var payload AuthenticatePayload
	err := c.BindHeader(&headers)
	if err != nil {

		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to bind")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var reqBody AuthenticateReq
	err = c.ShouldBindJSON(&reqBody)
	if err != nil || reqBody.Type != "Supabase" {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Unable to Decode Request Body")
		errResponse := ErrAuthenticate("Unable to Decode Request Body")
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	supabase_jwt_token := c.Request.Header.Get("TOKEN")
	ctx := context.Background()
	supabaseUser, err := supabaseClient.Auth.User(ctx, supabase_jwt_token)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Unable to Authenticate Token")

		c.JSON(http.StatusInternalServerError, "unable to authenticate ")
		return
	}
	claimsvalue.Email = supabaseUser.Email

	userDetails, err := user.GetUser(claimsvalue.Email)
	if err != nil {
		claimsvalue.UUID = uuid.NewString()
		err = user.AddUser(claimsvalue.UUID, claimsvalue.Email)
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("Unable to Add User ")
			errResponse := ErrAuthenticate(err.Error())
			c.JSON(http.StatusInternalServerError, errResponse)
			return
		}
	}
	claimsvalue.UUID = userDetails.UserId
	customClaims := claims.New(claimsvalue)
	pasetoToken, err := auth.GenerateTokenPaseto(customClaims)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to generate token")
		errResponse := ErrAuthenticate(err.Error())
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
	payload.Token = pasetoToken
	payload.Status = 200
	payload.Message = "success"
	payload.Success = true
	c.JSON(http.StatusOK, payload)
}

func ErrAuthenticate(errvalue string) AuthenticatePayload {
	var payload AuthenticatePayload
	payload.Success = false
	payload.Status = 401
	payload.Message = errvalue
	return payload
}
