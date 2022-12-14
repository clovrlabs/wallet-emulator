package actions

import (
	"fmt"
	"math"

	"github.com/breez/breez/bindings"
	"github.com/breez/breez/boltz"
)

func SendViaBtc(paymentRequest string, amountSatoshi int64) error {
	err := bindings.GetBreezApp().AccountService.ValidateAddress(paymentRequest)

	if err != nil {
		return err
	}

	fees, err := bindings.GetBreezApp().SwapService.ClaimFeeEstimates(paymentRequest)

	if err != nil {
		return err
	}

	info, err := boltz.GetReverseSwapInfo()

	if err != nil {
		return err
	}

	toSend := amountSatoshi
	var received int64 = 0
	for _, fee := range fees {
		boltzFees := int64(math.Ceil(float64(toSend)*info.Fees.Percentage/100)) + info.Fees.Lockup
		received = amountSatoshi - boltzFees - fee
	}

	hash, err := bindings.GetBreezApp().SwapService.NewReverseSwap(amountSatoshi, info.FeesHash, paymentRequest)

	if err != nil {
		return err
	}

	err = bindings.GetBreezApp().SwapService.SetReverseSwapClaimFee(hash, amountSatoshi-received)

	if err != nil {
		return err
	}

	swap, err := bindings.GetBreezApp().SwapService.FetchReverseSwap(hash)

	if err != nil {
		return err
	}

	err = bindings.GetBreezApp().SwapService.PayReverseSwap(hash, "", "", "", swap.ClaimFee)

	if err != nil {
		return err
	}

	return nil
}

func SendViaInvoice(paymentRequest string, amountSatoshi int64, fee int64) error {
	resp, err := bindings.GetBreezApp().AccountService.SendPaymentForRequest(paymentRequest, amountSatoshi, fee)

	fmt.Println(resp)

	if err != nil {
		return err
	}
	return nil
}

func SendViaNodeId(paymentRequest string, amountSatoshi int64) error {
	resp, err := bindings.GetBreezApp().AccountService.SendSpontaneousPayment(
		paymentRequest, "", amountSatoshi, 0, "", "", nil)

	fmt.Println(resp)

	if err != nil {
		return err
	}
	return nil
}
