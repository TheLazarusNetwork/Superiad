package fetchwallet

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheLazarusNetwork/mtwallet/config/dbconfig"
	"github.com/TheLazarusNetwork/mtwallet/config/envconfig"
	"github.com/TheLazarusNetwork/mtwallet/models"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/util/testingcommon"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_FetchWallet(t *testing.T) {
	envconfig.InitEnvVars()
	models.Migrate()
	gin.SetMode(gin.TestMode)
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	err := dbconfig.GetDb().Model(&user.User{}).Create(&user.User{
		UserId:   "67",
		Mnemonic: "canyon butter soup wet lemon rent mix brick wood wolf indicate fuel coral police beef goose rely frog toss roast beauty index tiger sample",
	}).Error
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	req := FetchWalletRequest{
		ChainId: 80001,
		Network: "polygon",
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
	fetchwallet(c)
	assert.Contains(t, rr.Body.String(), "0x2c08ABf0d1e7A89419bd8E44E946a2a2FA769Dd3")
	assert.Equal(t, 200, rr.Result().StatusCode)
}
