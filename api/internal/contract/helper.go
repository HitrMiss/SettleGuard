package contract

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func EncodePaymentCall(
	amount *big.Int,
	merchantIdentity common.Address,
	merchantPayout common.Address,
	categoryId [32]byte,
	packetId [32]byte,
) (string, error) {
	parsedABI, err := abi.JSON(strings.NewReader(SettlementEngineMetaData.ABI))
	if err != nil {
		return "", fmt.Errorf("failed to parse Settlement Engine ABI: %w", err)
	}

	inputData, err := parsedABI.Pack("processPayment",
		amount,
		merchantIdentity,
		merchantPayout,
		categoryId,
		packetId,
	)
	if err != nil {
		return "", fmt.Errorf("failed to pack arguments: %w", err)
	}

	return "0x" + common.Bytes2Hex(inputData), nil
}
