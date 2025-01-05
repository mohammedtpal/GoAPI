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
func (s *SmartContract) HandleGreeting1(ctx contractapi.TransactionContextInterface, name string) (string, error) {
	// Simulate a response based on the input
	response := fmt.Sprintf("Hello 	mst mst: , %s", name)
	return response, nil
}

// InitLedger initializes the ledger with some initial assets
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	return nil
}

// saveChunk saves the chunk data on the ledger using fileID and chunkID as keys, if the data is valid protobuf
func (s *SmartContract) saveChunk1(ctx contractapi.TransactionContextInterface, chunkBytes []byte, fileID string, chunkID string) error {
	// Attempt to unmarshal the chunkBytes to validate if it is valid protobuf
	// var chunk Chunk // Assuming Chunk is your protobuf message structure
	// if err := proto.Unmarshal(chunkBytes, &chunk); err != nil {
	// 	return fmt.Errorf("failed to unmarshal chunk data: %v", err)
	// }

	// // If unmarshaling is successful, save the chunkBytes to the ledger
	// key := fileID + "-" + chunkID

	// // Store the chunk on the ledger
	// err := ctx.GetStub().PutState(key, chunkBytes)
	// if err != nil {
	// 	return fmt.Errorf("failed to save chunk to the ledger: %v", err)
	// }

	return nil
}

// retrieveChunk retrieves the chunk data from the ledger using fileID and chunkID
func (s *SmartContract) retrieveChunk(ctx contractapi.TransactionContextInterface, fileID string, chunkID string) ([]byte, error) {
	// Create the key using fileID and chunkID
	key := fileID + "-" + chunkID

	// Retrieve the chunk from the ledger
	chunkBytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get chunk from the ledger: %v", err)
	}
	if chunkBytes == nil {
		return nil, fmt.Errorf("chunk not found for key: %s", key)
	}

	return chunkBytes, nil
}

// InvokeGreeting simulates an HTTP handler that receives a Protobuf message, processes it, and returns a response
func (s *SmartContract) InvokeGreeting(ctx contractapi.TransactionContextInterface, chunkBytes []byte,fileID string, chunkID string) ([]byte, error) {
	// Attempt to unmarshal the chunkBytes to validate if it is valid protobuf
	var chunk Chunk // Assuming Chunk is your protobuf message structure
	if err := proto.Unmarshal(chunkBytes, &chunk); err != nil {
		return nil,fmt.Errorf("failed to unmarshal chunk data: %v", err)
	}

	// If unmarshaling is successful, save the chunkBytes to the ledger
	key := fileID + "-" + chunkID

	// Store the chunk on the ledger
	err := ctx.GetStub().PutState(key, chunkBytes)
	if err != nil {
		return nil,fmt.Errorf("failed to save chunk to the ledger: %v", err)
	}

	// return chunkBytes, nil	
	
	
	// Unmarshal the incoming Protobuf data into the GreetingRequest message
	// var req Chunk
	// if err := proto.Unmarshal(data, &req); err != nil {
	// 	return nil, fmt.Errorf("failed to unmarshal Protobuf data: %v", err)
	// }

	// Marshal the data back to Protobuf format (without changing it)
	marshalledData, err := proto.Marshal(&chunk)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Protobuf data: %v", err)
	}

	// Return the marshalled data
	return marshalledData, nil
}

func (s *SmartContract) InvokeGreeting2(ctx contractapi.TransactionContextInterface, chunkBytes []byte,fileID string, chunkID string) ([]byte, error) {
	// Attempt to unmarshal the chunkBytes to validate if it is valid protobuf
	var chunk Chunk // Assuming Chunk is your protobuf message structure
	if err := proto.Unmarshal(chunkBytes, &chunk); err != nil {
		return nil,fmt.Errorf("failed to unmarshal chunk data: %v", err)
	}

	// If unmarshaling is successful, save the chunkBytes to the ledger
	key := fileID + "-" + chunkID

	// Store the chunk on the ledger
	err := ctx.GetStub().PutState(key, chunk.Data)
	if err != nil {
		return nil,fmt.Errorf("failed to save chunk to the ledger: %v", err)
	}

	// return chunkBytes, nil	
	
	
	// Unmarshal the incoming Protobuf data into the GreetingRequest message
	// var req Chunk
	// if err := proto.Unmarshal(data, &req); err != nil {
	// 	return nil, fmt.Errorf("failed to unmarshal Protobuf data: %v", err)
	// }

	// Marshal the data back to Protobuf format (without changing it)
	marshalledData, err := proto.Marshal(&chunk)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Protobuf data: %v", err)
	}

	// Return the marshalled data
	return marshalledData, nil
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
