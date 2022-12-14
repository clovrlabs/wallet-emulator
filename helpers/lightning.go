package helpers

import (
	"fmt"
	"time"

	"github.com/breez/breez/bindings"
	"github.com/breez/breez/config"
	"github.com/breez/breez/data"
	"github.com/clovrlabs/wallet-emulator/actions"
)

func SyncGraph() {
	cfg, err := config.GetConfig(bindings.GetBreezApp().GetWorkingDir())
	if err != nil || cfg.BootstrapURL == "" {
		return
	}

	url, err := bindings.GraphURL()

	if err != nil {
		return
	}
	Download(url, "./data/channel.db")
	time.Sleep(10 * time.Second)
	bindings.SyncGraphFromFile(url)
}

func SyncLSP() error {
	time.Sleep(10 * time.Second)
	lsp, err := actions.GetFirstLsp()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(lsp)
	err = bindings.ConnectToLSP(lsp.Id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	request := data.SyncLSPChannelsRequest{
		LspInfo: lsp,
	}
	_, err = bindings.GetBreezApp().SyncLSPChannels(&request)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
