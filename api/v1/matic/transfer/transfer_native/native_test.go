package transfer_native

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

func Test_Transfer(t *testing.T) {
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

	t.Run("Native token", func(t *testing.T) {
		rr := httptest.NewRecorder()

		req := TransferRequest{
			UserId: "60",
			To:     "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
			Amount: 1,
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
		nativeTransfer(c)

		assert.Equal(t, 200, rr.Result().StatusCode)

	})

}
