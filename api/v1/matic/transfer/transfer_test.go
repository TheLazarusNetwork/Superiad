package transfer

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func Test_Transfer(t *testing.T) {
	envconfig.InitEnvVars()
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

	type network struct {
		name          string
		chainId       int64
		erc20Address  string
		erc721Address string
	}
	networks := []network{
		{
			name:         "polygon",
			chainId:      80001,
			erc20Address: "0x2d7882beDcbfDDce29Ba99965dd3cdF7fcB10A1e",
		}, {
			name:         "ethereum",
			chainId:      1,
			erc20Address: "0x514910771af9ca656af840dff83e8264ecf986ca",
		},
	}

	t.Run("Native token", func(t *testing.T) {
		for _, n := range networks {
			rr := httptest.NewRecorder()

			req := TransferRequest{
				UserId:  "60",
				To:      "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
				Amount:  1,
				ChainId: n.chainId,
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
			transfer(c)

			assert.Equal(t, 200, rr.Result().StatusCode)
		}

	})

	t.Run("ERC20 Token", func(t *testing.T) {
		for _, n := range networks {
			rr := httptest.NewRecorder()

			req := TransferRequest{
				UserId:  "60",
				To:      "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
				Amount:  1,
				ChainId: n.chainId,
			}
			d, e := json.Marshal(req)
			if e != nil {
				t.Fatal(e)
			}
			c, _ := gin.CreateTestContext(rr)
			reqUrl := fmt.Sprintf("/?erc20Address=%v", n.erc20Address)
			httpReq, e := http.NewRequest("GET", reqUrl, bytes.NewBuffer(d))
			if e != nil {
				t.Fatal(e)
			}
			c.Request = httpReq
			transfer(c)
			assert.Equal(t, 200, rr.Result().StatusCode)
		}

	})
	t.Run("ERC721 Token", func(t *testing.T) {
		for _, n := range networks {
			rr := httptest.NewRecorder()

			req := TransferRequest{
				UserId:  "60",
				To:      "0x876FA09c042E6CA0c2f73AAe1DD7Bf712b6BF8f0",
				Amount:  1,
				ChainId: n.chainId,
			}
			d, e := json.Marshal(req)
			if e != nil {
				t.Fatal(e)
			}
			c, _ := gin.CreateTestContext(rr)
			reqUrl := fmt.Sprintf("/?erc721Address=%v", n.erc721Address)
			httpReq, e := http.
				NewRequest("GET", reqUrl,
					bytes.NewBuffer(d))
			if e != nil {
				t.Fatal(e)
			}
			c.Request = httpReq
			transfer(c)
			assert.Equal(t, 200, rr.Result().StatusCode)
		}

	})

}
