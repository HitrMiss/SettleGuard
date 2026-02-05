package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// internal/contract/helper.go

func EncodePaymentCall(
	amount *big.Int,
	merchantIdentity common.Address,
	merchantPayout common.Address,
	categoryId [32]byte,
	packetId [32]byte,
) (string, error) {
	parsedABI, err := abi.JSON(strings.NewReader(PaymentVaultMetaData.ABI))
	if err != nil {
		return "", err
	}

	inputData, err := parsedABI.Pack("lockFunds",
		common.Address{},
		amount,
		merchantIdentity,
		merchantPayout,
		categoryId,
		packetId,
	)
	return "0x" + common.Bytes2Hex(inputData), err
}
