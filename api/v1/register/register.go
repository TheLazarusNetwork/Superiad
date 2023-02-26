package register

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TheLazarusNetwork/superiad/config/envconfig"
	"github.com/TheLazarusNetwork/superiad/models/user"
	"github.com/TheLazarusNetwork/superiad/pkg/auth"
	"github.com/TheLazarusNetwork/superiad/pkg/claims"
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
	g := r.Group("/register")
	{
		g.GET("/", register)
	}
}
func Init() {
	supabaseUrl := envconfig.EnvVars.SUPABASEURL
	supabaseKey := envconfig.EnvVars.SUPABASEKEY
	supabaseClient = supa.CreateClient(supabaseUrl, supabaseKey)
}

func register(c *gin.Context) {
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
	if headers.Authorization == "" {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Autherisation header is missing")
		c.Abort()
		return
	}
	supabase_jwt_token := c.Request.Header.Get("SUPAJWT_TOKEN")
	ctx := context.Background()
	supabaseUser, err := supabaseClient.Auth.User(ctx, supabase_jwt_token)
	if err != nil {
		fmt.Println("err value", err.Error())
	}
	claimsvalue.Email = supabaseUser.Email
	// insert into db or check for uuid if present return the paseto token
	userDetails, err := user.GetUser(claimsvalue.Email)
	if err != nil {
		claimsvalue.UUID = uuid.NewString()
		err = user.AddUser(claimsvalue.UUID, claimsvalue.Email)
		if err != nil {
			//log unable to add user
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
