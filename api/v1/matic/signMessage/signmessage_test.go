package signmessage

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MyriadFlow/superiad/app/stage/appinit"
	"github.com/MyriadFlow/superiad/config/envconfig"

	"github.com/MyriadFlow/superiad/models/user"
	"github.com/MyriadFlow/superiad/pkg/store"
	"github.com/MyriadFlow/superiad/pkg/testingcommon"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_SignMessage(t *testing.T) {
	envconfig.InitEnvVars()

	appinit.Init()
	gin.SetMode(gin.TestMode)
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	err := store.DB.Model(&user.User{}).Create(&user.User{
		UserId:   "60",
		Mnemonic: "mobile attend orange oxygen valley fan grape suit tool fancy quality disease potato bean trophy",
	}).Error
	if err != nil {
		t.Fatal(err)
	}

	_ = "0x844507295543d7dcb9a6bf5fa436437eaec309aa2edcfc694a407b8a30e84b464d02934432478d1e201440ec5712c6d7e15e66e1db672cf01d0cf9a1003926881c"

	t.Run("Native token", func(t *testing.T) {
		message := "test message to sign"
		rr := httptest.NewRecorder()

		req := SignMessageRequest{
			UserId:  "60",
			Message: message,
		}
		d, e := json.Marshal(req)
		if e != nil {
			t.Fatal(e)
		}
		c, _ := gin.CreateTestContext(rr)
		httpReq, e := http.NewRequest("GET", "/", bytes.NewBuffer(d))
		if e != nil {
			t.Fatal(e)
		}
		c.Request = httpReq
		signMessage(c)

		assert.Equal(t, 200, rr.Result().StatusCode)
		t.Fatal(rr.Body.String())

	})
}
