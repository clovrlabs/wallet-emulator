package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/breez/breez/bindings"
	"github.com/clovrlabs/wallet-emulator/client"
	"github.com/clovrlabs/wallet-emulator/helpers"
	"github.com/clovrlabs/wallet-emulator/routes"
	"github.com/gin-gonic/gin"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	helpers.CopyConfigFiles()

	bindings.Init("./temp", "./data", client.AppServices{})
	bindings.Start()

	bindings.NewSyncJob("./data")
	go helpers.SyncGraph()
	go helpers.SyncLSP()

	router := gin.New()
	gin := routes.SetupRouter(router, basepath)

	port := "3000"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	gin.Run(fmt.Sprintf(":%s", port))
}
