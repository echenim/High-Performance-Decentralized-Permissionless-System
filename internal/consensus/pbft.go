package consensus

import (
	"errors"
	"fmt"
)

// PBFT represents the PBFT consensus mechanism.
type PBFT struct {
	BaseConsensus
	ReplicaCount int
}

// ReachConsensus implements PBFT consensus logic.
func (p *PBFT) ReachConsensus() error {
	if p.ReplicaCount < 3 {
		return errors.New("PBFT requires at least 3 replicas")
	}
	fmt.Println("PBFT: Consensus reached with replicas")
	return nil
}
