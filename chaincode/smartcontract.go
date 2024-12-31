package chaincode

import (
	"fmt"
	"log"
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
	// pb"github.com/mohammedtpal/fabric-samples/asset-transfer-basic/chaincode-go3/chaincode/protoF"
	"github.com/golang/protobuf/proto"
	// "chaincode/protoF"
)
// import "github.com/mohammedtpal/fabric-samples/asset-transfer-basic/chaincode-go3/chaincode/protoF"


// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// HandleGreeting function simulates greeting functionality
func (s *SmartContract) HandleGreeting(ctx contractapi.TransactionContextInterface, name string) (string, error) {
	// Simulate a response based on the input
	response := fmt.Sprintf("Hello, %s", name)
	return response, nil
}

// InitLedger initializes the ledger with some initial assets
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	return nil
}

// InvokeGreeting simulates an HTTP handler that receives a Protobuf message, processes it, and returns a response
func (s *SmartContract) InvokeGreeting(ctx contractapi.TransactionContextInterface, data []byte) ([]byte, error) {
	// Unmarshal the incoming Protobuf data into the GreetingRequest message
	var req GreetingRequest
	if err := proto.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Protobuf data: %v", err)
	}

	// Simulate chaincode logic to handle the greeting
	message, err := s.HandleGreeting(ctx, req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to handle greeting: %v", err)
	}

	// Create a response message
	resp := &GreetingResponse{
		Message: message,
	}

	// Marshal the response to Protobuf and return it
	respData, err := proto.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Protobuf response: %v", err)
	}

	return respData, nil
}

func main() {
	// Create a new SmartContract instance and start it
	smartContract := new(SmartContract)

	// Create the chaincode and start it
	cc, err := contractapi.NewChaincode(smartContract)
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	// Start the chaincode
	if err := cc.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
