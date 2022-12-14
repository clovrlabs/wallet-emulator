package actions

import (
	"errors"

	"github.com/breez/breez/bindings"
	"github.com/breez/breez/data"
)

func GetFirstLsp() (*data.LSPInformation, error) {
	lspList, err := bindings.GetBreezApp().ServicesClient.LSPList()
	if err != nil {
		return &data.LSPInformation{}, err
	}

	for k := range lspList.Lsps {
		return lspList.Lsps[k], nil
	}
	return &data.LSPInformation{}, errors.New("Unreachable")
}
