package consensus

import "errors"

// Consensus is the interface for all consensus mechanisms.
type Consensus interface {
	ProposeBlock(data string) (string, error)
	ValidateBlock(block string) bool
	ReachConsensus() error
}

// BaseConsensus provides shared logic for consensus mechanisms.
type BaseConsensus struct {
	NodeID string
}

// ProposeBlock is a default implementation for proposing blocks.
func (bc *BaseConsensus) ProposeBlock(data string) (string, error) {
	if data == "" {
		return "", errors.New("data cannot be empty")
	}
	block := "Block:" + data
	return block, nil
}

// ValidateBlock is a stub for block validation logic.
func (bc *BaseConsensus) ValidateBlock(block string) bool {
	return block != ""
}
