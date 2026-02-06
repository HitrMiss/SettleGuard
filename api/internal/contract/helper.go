package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// EncodeDepositCall generates the calldata for the user to call SettlementEngine.deposit
func EncodeDepositCall(
	amount *big.Int,
	merchantIdentity common.Address,
	merchantPayout common.Address,
	categoryId [32]byte,
	packetId [32]byte,
	createdAt uint64,
	r *big.Int,
	s *big.Int,
) (string, error) {
	parsedABI, err := abi.JSON(strings.NewReader(SettlementEngineMetaData.ABI))
	if err != nil {
		return "", err
	}

	inputData, err := parsedABI.Pack("deposit",
		amount,
		merchantIdentity,
		merchantPayout,
		categoryId,
		packetId,
		createdAt,
		r,
		s,
	)
	if err != nil {
		return "", err
	}

	return "0x" + common.Bytes2Hex(inputData), nil
}
