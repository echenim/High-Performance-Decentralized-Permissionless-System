package consensus

import "fmt"

// DPoS represents the Delegated Proof-of-Stake consensus mechanism.
type DPoS struct {
	BaseConsensus
	Delegates []string
}

// ReachConsensus implements DPoS consensus logic.
func (d *DPoS) ReachConsensus() error {
	if len(d.Delegates) == 0 {
		return fmt.Errorf("no delegates assigned")
	}
	fmt.Println("DPoS: Consensus reached with delegates")
	return nil
}
