package tokenmiddleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheLazarusNetwork/superiad/config/envconfig"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_APIAUTH(t *testing.T) {
	envconfig.InitEnvVars()

	correctToken := envconfig.EnvVars.TOKEN
	t.Run("success if token is valid", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		req, e := http.NewRequest("GET", "", nil)
		req.Header.Add("Authorization", correctToken)
		if e != nil {
			t.Fatal(e)
		}
		c.Request = req
		ApiAuth(c)
		assert.Equal(t, 200, rr.Result().StatusCode)
	})

	t.Run("fail if token is notvalid", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		req, e := http.NewRequest("GET", "", nil)
		req.Header.Add("Authorization", correctToken+"extra data which will invalidate token")
		if e != nil {
			t.Fatal(e)
		}
		c.Request = req
		ApiAuth(c)
		assert.Equal(t, http.StatusForbidden, rr.Result().StatusCode)
	})
}
