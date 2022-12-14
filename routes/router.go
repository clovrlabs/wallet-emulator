package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/breez/breez/bindings"
	"github.com/clovrlabs/wallet-emulator/actions"
	"github.com/gin-gonic/gin"
)

func getAccountInfo(c *gin.Context) {
	account, err := bindings.GetBreezApp().AccountService.GetAccountInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"balance":       account.Balance,
		"walletBalance": account.WalletBalance,
	})
}

func receiveWithInvoice(c *gin.Context) {
	var body struct {
		Amount int64 `binding:"required"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount is required"})
		return
	}

	lnurl, qrcode, fee, err := actions.ReceiveInvoice(body.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"qrcode":  qrcode,
		"invoice": lnurl,
		"fee":     fee,
	})
}

func receiveWithBtc(c *gin.Context) {
	resp, err := actions.ReceiveWithBtc()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"qrcode":  resp.QrCode,
		"address": resp.Address,
	})
}

func sendViaInvoiceOrId(c *gin.Context) {
	// get string
	var body struct {
		PaymentRequest string `binding:"required"`
		AmountSatoshi  int64
		Fee            int64
	}

	if err := c.ShouldBind(&body); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "PaymentRequest is required"})
		return
	}
	//check if lightning address, or nodeid (send spontaneous payment)
	lnb := body.PaymentRequest[0:4]
	var err error
	if strings.EqualFold(lnb, "lnbc") {
		err = actions.SendViaInvoice(body.PaymentRequest, body.AmountSatoshi, body.Fee)
	} else {
		// Doesn't work yet
		err = actions.SendViaNodeId(body.PaymentRequest, body.AmountSatoshi)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func sendViaBtc(c *gin.Context) {
	// get string
	var body struct {
		PaymentRequest string `binding:"required"`
		AmountSatoshi  int64
		Fee            int64
	}

	if err := c.ShouldBind(&body); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "PaymentRequest is required"})
		return
	}

	err := actions.SendViaBtc(body.PaymentRequest, body.AmountSatoshi)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func sendCommand(c *gin.Context) {
	var body struct {
		Command string `binding:"required"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Command is required"})
		return
	}

	response, err := bindings.SendCommand(body.Command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"command": response,
	})
}

func SetupRouter(router *gin.Engine, basepath string) *gin.Engine {
	router.StaticFile("/", basepath+"/public/html/index.html")
	router.StaticFile("/receive", basepath+"/public/html/receive.html")
	router.StaticFile("/send", basepath+"/public/html/send.html")
	router.StaticFile("/developers", basepath+"/public/html/developers.html")
	router.Static("/public", basepath+"/public")

	router.POST("/receive/invoice", receiveWithInvoice)
	router.GET("/receive/btc", receiveWithBtc)
	router.POST("/send/invoice", sendViaInvoiceOrId)
	router.POST("/send/btc", sendViaBtc)
	router.GET("/account/info", getAccountInfo)
	router.POST("/command", sendCommand)
	return router
}
