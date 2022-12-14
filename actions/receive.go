package actions

import (
	"github.com/breez/breez/bindings"
	"github.com/breez/breez/data"
	qrcode "github.com/skip2/go-qrcode"
)

func ReceiveInvoice(amount int64) (string, []byte, int64, error) {
	lsp, err := GetFirstLsp()
	if err != nil {
		return "", nil, 0, err
	}

	request := data.AddInvoiceRequest{}
	request.InvoiceDetails = &data.InvoiceMemo{}
	request.InvoiceDetails.Amount = amount
	request.LspInfo = lsp
	payreq, fee, err := bindings.GetBreezApp().AccountService.AddInvoice(&request)

	if err != nil {
		return "", nil, 0, err
	}

	png, err := qrcode.Encode(payreq, qrcode.Medium, 256)

	if err != nil {
		return "", nil, 0, err
	}

	return payreq, png, fee, nil
}

type ReceiveWithBtcResponse struct {
	QrCode  []byte
	Address string
}

func ReceiveWithBtc() (*ReceiveWithBtcResponse, error) {
	lsp, err := GetFirstLsp()
	if err != nil {
		return &ReceiveWithBtcResponse{}, err
	}

	resp, err := bindings.GetBreezApp().SwapService.AddFundsInit("", lsp.Id)
	if err != nil {
		return &ReceiveWithBtcResponse{}, err
	}

	qrcode, err := qrcode.Encode("bitcoin:"+resp.Address, qrcode.Medium, 256)

	if err != nil {
		return &ReceiveWithBtcResponse{}, err
	}

	return &ReceiveWithBtcResponse{
		QrCode:  qrcode,
		Address: resp.Address,
	}, nil
}
