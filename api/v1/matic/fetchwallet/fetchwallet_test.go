package fetchwallet

import (
	"net/http/httptest"
	"testing"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/app/stage/appinit"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/config/envconfig"

	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/models/user"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/store"
	"github.com/VirtuaTechnologies/VirtuaCoin_Wallet/pkg/testingcommon"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_FetchWallet(t *testing.T) {
	envconfig.InitEnvVars()

	appinit.Init()
	gin.SetMode(gin.TestMode)
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	err := store.DB.Model(&user.User{}).Create(&user.User{
		UserId:   "67",
		Mnemonic: "canyon butter soup wet lemon rent mix brick wood wolf indicate fuel coral police beef goose rely frog toss roast beauty index tiger sample",
	}).Error
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rr)
	fetchwallet(c)
	assert.Contains(t, rr.Body.String(), "0x2c08ABf0d1e7A89419bd8E44E946a2a2FA769Dd3")
	assert.Equal(t, 200, rr.Result().StatusCode)
}
