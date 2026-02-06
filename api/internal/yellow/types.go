package yellow

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Network Endpoints
const (
	ProductionEndpoint = "wss://clearnet.yellow.com/ws"
	SandboxEndpoint    = "wss://clearnet-sandbox.yellow.com/ws"
)

// StateIntent represents the channel status intent
type StateIntent int

const (
	Operate StateIntent = iota
	Initialize
	Resize
	Finalize
)

type AppDefinition struct {
	Protocol     string           `json:"protocol"`
	Participants []common.Address `json:"participants"`
	Weights      []int            `json:"weights"`
	Quorum       int              `json:"quorum"`
	Challenge    int              `json:"challenge"`
	Nonce        int              `json:"nonce"`
}

type AppAllocation struct {
	Participant common.Address `json:"participant"`
	Asset       common.Address `json:"asset"`
	Amount      *big.Int       `json:"amount"`
}

type AppSession struct {
	Definition  AppDefinition   `json:"definition"`
	Allocations []AppAllocation `json:"allocations"`
}

type RPCRequest struct {
	ID        int           `json:"id"`
	Method    string        `json:"method"`
	Params    []interface{} `json:"params"`
	Timestamp int64         `json:"timestamp,omitempty"`
}

type RPCResponse struct {
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
