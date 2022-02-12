package transfer

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

func Test_Transfer(t *testing.T) {
	config.Init("../../../.env")
	models.Migrate()
	gin.SetMode(gin.TestMode)
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	err := dbconfig.GetDb().Model(&user.User{}).Create(&user.User{
		UserId:   "60",
		Mnemonic: "mobile attend orange oxygen valley fan grape suit tool fancy quality disease potato bean trophy",
	}).Error
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Native token", func(t *testing.T) {
		rr := httptest.NewRecorder()

		req := TransferRequest{
			UserId:  "60",
			To:      "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
			Amount:  1,
			ChainId: 80001,
		}
		d, e := json.Marshal(req)
		if e != nil {
			t.Fatal(e)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Params = gin.Params{{Key: "network", Value: "polygon"}}
		httpReq, e := http.NewRequest("GET", "/", bytes.NewBuffer(d))
		if e != nil {
			t.Fatal(e)
		}
		c.Request = httpReq
		transfer(c)

		assert.Equal(t, 200, rr.Result().StatusCode)
	})

	t.Run("ERC20 Token", func(t *testing.T) {
		rr := httptest.NewRecorder()

		req := TransferRequest{
			UserId:  "60",
			To:      "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
			Amount:  1,
			ChainId: 80001,
		}
		d, e := json.Marshal(req)
		if e != nil {
			t.Fatal(e)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Params = gin.Params{{Key: "network", Value: "polygon"}}
		httpReq, e := http.NewRequest("GET", "/?erc20Address=0x2d7882beDcbfDDce29Ba99965dd3cdF7fcB10A1e", bytes.NewBuffer(d))
		if e != nil {
			t.Fatal(e)
		}
		c.Request = httpReq
		transfer(c)
		assert.Equal(t, 200, rr.Result().StatusCode)
	})
	t.Run("ERC721 Token", func(t *testing.T) {
		rr := httptest.NewRecorder()

		req := TransferRequest{
			UserId:  "60",
			To:      "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
			Amount:  1,
			ChainId: 80001,
		}
		d, e := json.Marshal(req)
		if e != nil {
			t.Fatal(e)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Params = gin.Params{{Key: "network", Value: "polygon"}}
		httpReq, e := http.
			NewRequest("GET", "/?erc721Address=0x975362c36b6842d48d02dbd3a077745fc1c64175&erc721TokenId=22",
				bytes.NewBuffer(d))
		if e != nil {
			t.Fatal(e)
		}
		c.Request = httpReq
		transfer(c)
		assert.Equal(t, 200, rr.Result().StatusCode)
	})

}
