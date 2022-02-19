package checkbalance

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheLazarusNetwork/mtwallet/config"
	"github.com/TheLazarusNetwork/mtwallet/config/dbconfig"
	"github.com/TheLazarusNetwork/mtwallet/models"
	"github.com/TheLazarusNetwork/mtwallet/models/user"
	"github.com/TheLazarusNetwork/mtwallet/util/testingcommon"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_CheckBalance(t *testing.T) {
	config.Init("../../../.env")
	models.Migrate()
	gin.SetMode(gin.TestMode)
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	err := dbconfig.GetDb().Model(&user.User{}).Create(&user.User{
		UserId:   "62",
		Mnemonic: "long hen advance measure donate child method aspect ceiling saddle turkey cement duck finger armor clarify hamster acid advice caution lazy deal invite remind",
	}).Error
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Fetch user balance", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		req := CheckBalanceRequest{
			UserId:  "62",
			ChainId: 80001,
		}
		body, err := json.Marshal(req)
		if err != nil {
			t.Fatal(err)
		}

		httpReq, err := http.NewRequest("POST", "/", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		c.Request = httpReq
		c.Params = gin.Params{{Key: "network", Value: "polygon"}}
		checkBalance(c)
		assert.Equal(t, 200, rr.Result().StatusCode)
	})
	t.Run("Fetch user balance for ERC20", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		req := CheckBalanceRequest{
			UserId:  "62",
			ChainId: 80001,
		}
		body, err := json.Marshal(req)
		if err != nil {
			t.Fatal(err)
		}

		httpReq, err := http.NewRequest("POST", "/?erc20address=0x2d7882bedcbfddce29ba99965dd3cdf7fcb10a1e", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		c.Request = httpReq
		c.Params = gin.Params{{Key: "network", Value: "polygon"}}
		checkBalance(c)
		assert.Equal(t, 200, rr.Result().StatusCode)
	})
	t.Run("Fetch user balance for ERC721", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)

		req := CheckBalanceRequest{
			UserId:  "62",
			ChainId: 80001,
		}
		body, err := json.Marshal(req)
		if err != nil {
			t.Fatal(err)
		}

		httpReq, err := http.NewRequest("POST", "/?erc721address=0x975362c36b6842d48d02DBD3A077745Fc1C64175", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		c.Request = httpReq
		c.Params = gin.Params{{Key: "network", Value: "polygon"}}
		checkBalance(c)
		assert.Equal(t, 200, rr.Result().StatusCode)
	})
}
